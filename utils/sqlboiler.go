package utils

import (
	"fmt"
	"go/ast"
	"go/token"
	"path"
	"strings"
)

///////////////////////////////////////////////////////////////
// Cache
///////////////////////////////////////////////////////////////

var (
	enumCacheMap = map[string]SQLBoilerTableColumnEnum{}
)

///////////////////////////////////////////////////////////////
// SQLBoiler Types
///////////////////////////////////////////////////////////////

type SQLBoilerName struct {
	PascalCase string
	SnakeCase  string
	CamelCase  string
}

///////////////////////////////////////////////////////////////
// Main Function | Get SQLBoiler Models and Enums
///////////////////////////////////////////////////////////////

type SQLBoilerTable struct {
	Name      SQLBoilerName
	Relations []SQLBoilerTableRelation
	Columns   []SQLBoilerTableColumn
	Imports   []string
}

type SQLBoilerTableRelation struct {
	Name          SQLBoilerName
	MainTableName SQLBoilerName // Don't know if this is really needed
	IsMany        bool
	Type          SQLBoilerTableColumnType
}

type SQLBoilerTableColumn struct {
	Name SQLBoilerName
	Type SQLBoilerTableColumnType
}

type SQLBoilerTableColumnType struct {
	GoType         string
	GoTypeName     string
	TypescriptType string
	IsNullable     bool
	IsEnum         bool
}

func (c *Config) getSQLBoilerTablesAndEnums() ([]SQLBoilerTable, []SQLBoilerTableColumnEnum, error) {
	enums, err := c.readSQLBoilerEnumsFromFile()
	if err != nil {
		return nil, nil, err
	}

	tables, err := c.readSQLBoilerTablesFromFile()
	if err != nil {
		return nil, nil, err
	}

	for i, table := range tables {
		columns, relations, err := c.readSQLBoilerColumnsAndRelationsFromFile(table)
		if err != nil {
			return nil, nil, err
		}

		tables[i] = SQLBoilerTable{
			Name:      table.Name,
			Columns:   columns,
			Relations: relations,
			Imports:   c.getERGDefaultImports(),
		}
	}

	return tables, enums, nil
}

///////////////////////////////////////////////////////////////
// Get SQLBoiler Enums
///////////////////////////////////////////////////////////////

type SQLBoilerTableColumnEnum struct {
	Name   SQLBoilerName
	Values []SQLBoilerTableColumnEnumValue
	Type   SQLBoilerTableColumnType
}

type SQLBoilerTableColumnEnumValue struct {
	Label string
	Value string
}

func (c *Config) readSQLBoilerEnumsFromFile() ([]SQLBoilerTableColumnEnum, error) {
	f, err := parseFile(c.getSQLBoilerModelsFilePath("boil_types.go"))
	if err != nil {
		return nil, err
	}

	var currentType string
	ast.Inspect(f, func(n ast.Node) bool {
		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			currentType = typeSpec.Name.Name
			if basicLit, ok := typeSpec.Type.(*ast.Ident); ok {
				enumCacheMap[currentType] = SQLBoilerTableColumnEnum{
					Type: SQLBoilerTableColumnType{
						GoType:     basicLit.Name,
						GoTypeName: basicLit.Name,
					},
				}
			}
		}

		if currentType != "" {
			if genDecl, ok := n.(*ast.GenDecl); ok && genDecl.Tok == token.CONST {
				for _, spec := range genDecl.Specs {
					if valueSpec, ok := spec.(*ast.ValueSpec); ok {
						for i, name := range valueSpec.Names {
							if i < len(valueSpec.Values) {
								if basicLit, ok := valueSpec.Values[i].(*ast.BasicLit); ok {
									enumCacheMap[currentType] = SQLBoilerTableColumnEnum{
										Name: SQLBoilerName{
											PascalCase: currentType,
											SnakeCase:  toSnakeCase(currentType),
											CamelCase:  toCamelCase(currentType),
										},
										Values: append(enumCacheMap[currentType].Values, SQLBoilerTableColumnEnumValue{
											Label: strings.TrimPrefix(name.Name, currentType),
											Value: basicLit.Value,
										}),
										Type: enumCacheMap[currentType].Type,
									}
								}
							}
						}
					}
				}
			}
		}
		return true
	})

	var enums []SQLBoilerTableColumnEnum
	for _, enum := range enumCacheMap {
		enums = append(enums, enum)
	}

	return enums, err
}

///////////////////////////////////////////////////////////////
// Get SQLBoiler Tables
///////////////////////////////////////////////////////////////

func (c *Config) readSQLBoilerTablesFromFile() ([]SQLBoilerTable, error) {
	f, err := parseFile(c.getSQLBoilerModelsFilePath("boil_table_names.go"))
	if err != nil {
		return nil, err
	}

	var models []SQLBoilerTable
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ValueSpec:
			for i, name := range x.Names {
				if name.Name == "TableNames" {
					values := x.Values[i].(*ast.CompositeLit)
					for _, elt := range values.Elts {
						if kv, ok := elt.(*ast.KeyValueExpr); ok {
							if label, ok := kv.Key.(*ast.Ident); ok {
								if strVal, ok := kv.Value.(*ast.BasicLit); ok {
									// Extract the string value and unquote it
									value := strVal.Value
									value = value[1 : len(value)-1] // Remove surrounding quotes

									if c.isBlackListed(value, "") {
										continue
									}

									models = append(models, SQLBoilerTable{
										Name: SQLBoilerName{
											PascalCase: label.Name,
											SnakeCase:  value,
											CamelCase:  toCamelCase(label.Name),
										},
									})
								}
							}
						}
					}
				}
			}
		}

		return true
	})

	return models, nil
}

///////////////////////////////////////////////////////////////
// Get SQLBoiler Model Columns and Relations
///////////////////////////////////////////////////////////////

func (c *Config) readSQLBoilerColumnsAndRelationsFromFile(table SQLBoilerTable) ([]SQLBoilerTableColumn, []SQLBoilerTableRelation, error) {
	f, err := parseFile(c.getSQLBoilerModelsFilePath(table.Name.SnakeCase + ".go"))
	if err != nil {
		return nil, nil, err
	}

	var columns []SQLBoilerTableColumn
	var relations []SQLBoilerTableRelation

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if structType, isStruct := x.Type.(*ast.StructType); isStruct && x.Name.Name == table.Name.CamelCase+"R" {

				for _, field := range structType.Fields.List {
					pascalCase := field.Names[0].Name

					t := getFieldType(field.Type)

					mainTableNamePascalCase := strings.TrimSuffix(t.GoType, "Slice")
					if c.isBlackListed(toSnakeCase(mainTableNamePascalCase), "") {
						continue
					}

					relations = append(relations, SQLBoilerTableRelation{
						Name: SQLBoilerName{
							PascalCase: pascalCase,
							SnakeCase:  toSnakeCase(pascalCase),
							CamelCase:  toCamelCase(pascalCase),
						},
						MainTableName: SQLBoilerName{
							PascalCase: mainTableNamePascalCase,
							SnakeCase:  toSnakeCase(mainTableNamePascalCase),
							CamelCase:  toCamelCase(mainTableNamePascalCase),
						},
						IsMany: strings.HasSuffix(t.GoType, "Slice"),
						Type:   t,
					})
				}
			}
			if structType, isStruct := x.Type.(*ast.StructType); isStruct && x.Name.Name == table.Name.PascalCase {
				for _, field := range structType.Fields.List {
					if !hasTag(field, "-") {
						t := getFieldType(field.Type)
						pascalCase := field.Names[0].Name
						snakeCase := getSnakeCaseFromTag(field) // Can be named differently than the field name

						if c.isBlackListed(table.Name.SnakeCase, snakeCase) {
							continue
						}

						columns = append(columns, SQLBoilerTableColumn{
							Name: SQLBoilerName{
								PascalCase: pascalCase,
								SnakeCase:  snakeCase,
								CamelCase:  toCamelCase(pascalCase),
							},
							Type: t,
						})
					}
				}
			}
		}
		return true
	})

	return columns, relations, nil
}

///////////////////////////////////////////////////////////////
// Write SQLBoiler Models to ERG Files
///////////////////////////////////////////////////////////////

func (c *Config) writeSQLBoilerTablesToERGFiles(tables []SQLBoilerTable) error {
	for _, table := range tables {
		if err := c.writeTemplate("templates/erg_table.gotpl", path.Join(c.sqlBoilerConfig.Erg.Output, fmt.Sprintf("%v.go", table.Name.SnakeCase)), table); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) writeSQLBoilerEnumsToERGFiles(enums []SQLBoilerTableColumnEnum) error {
	if err := c.writeTemplate("templates/erg_enums.gotpl", path.Join(c.sqlBoilerConfig.Erg.Output, "erg_types.go"), enums); err != nil {
		return err
	}

	return nil
}

type ergModel struct {
	Imports []string
	Tables  []SQLBoilerTable
}

func (c *Config) writeSQLBoilerTablesToERGFile(tables []SQLBoilerTable) error {
	return c.writeTemplate("templates/erg_tables.gotpl", path.Join(c.sqlBoilerConfig.Erg.Output, "erg_tables.go"), ergModel{
		Imports: c.getERGDefaultImports(),
		Tables:  tables,
	})
}

func (c *Config) writeERGHelperFunctionsToFile() error {
	return c.writeTemplate("templates/erg_helpers.gotpl", path.Join(c.sqlBoilerConfig.Erg.Output, "erg_helpers.go"), nil)
}

type tsModel struct {
	Tables []SQLBoilerTable
	Enums  []SQLBoilerTableColumnEnum
}

func (c *Config) writeSQLBoilerTablesToTypeScriptFiles(tables []SQLBoilerTable, enums []SQLBoilerTableColumnEnum) error {
	if !c.isTSEnabled() {
		return nil
	}

	if err := c.writeTemplate("templates/ts_tables.gotpl", c.sqlBoilerConfig.Erg.OutputTS, tsModel{
		Tables: tables,
		Enums:  enums,
	}); err != nil {
		return err
	}

	return nil
}
