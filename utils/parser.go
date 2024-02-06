package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func parseFile(fileName string) (*ast.File, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func hasTag(field *ast.Field, tag string) bool {
	if field.Tag == nil {
		return false
	}
	return strings.Contains(field.Tag.Value, tag)
}

func getSnakeCaseFromTag(field *ast.Field) string {
	if field.Tag != nil {
		tag := field.Tag.Value
		tag = strings.Trim(tag, "`")
		tagParts := strings.Split(tag, " ")
		for _, part := range tagParts {
			if strings.HasPrefix(part, "boil:") {
				return strings.Trim(strings.TrimPrefix(part, "boil:"), `"`)
			}
		}
	}
	return ""
}

func getFieldType(fieldType ast.Expr) SQLBoilerTableColumnType {
	var tp SQLBoilerTableColumnType

	switch t := fieldType.(type) {
	case *ast.Ident:
		tp = SQLBoilerTableColumnType{
			GoType:     t.Name,
			GoTypeName: t.Name,
			IsNullable: false,
		}
		break
	case *ast.StarExpr:
		ft := getFieldType(t.X)

		tp = SQLBoilerTableColumnType{
			GoType:     ft.GoType,
			GoTypeName: ft.GoTypeName,
			IsNullable: true,
		}
		break
	case *ast.SelectorExpr:
		ft := getFieldType(t.X)

		tp = SQLBoilerTableColumnType{
			GoType:     ft.GoType + "." + t.Sel.Name,
			GoTypeName: ft.GoTypeName + "." + t.Sel.Name,
			IsNullable: ft.IsNullable,
		}
		break
	default:
		fmt.Println("unknown", t)
	}

	tp.IsNullable = strings.Contains(tp.GoType, "null.") || tp.IsNullable

	if strings.Contains(tp.GoTypeName, ".") {
		s := strings.Split(tp.GoTypeName, ".")
		tp.GoTypeName = toSnakeCase(s[len(s)-1])
	}

	if tp.GoTypeName == "time" {
		tp.GoTypeName = "string"
	}

	if tp.GoTypeName == "json" {
		tp.GoTypeName = "json"
		tp.GoType = "json"
	}

	if tp.IsNullable {
		tp.GoTypeName = "*" + tp.GoTypeName
	}

	if strings.Contains(tp.GoTypeName, "_array") {
		s := strings.Split(tp.GoTypeName, "_")
		tp.GoTypeName = "[]" + s[0]
		tp.GoType = tp.GoTypeName
	}

	tp.TypescriptType = convertGoTypeToTypescript(tp)

	if _, ok := enumCacheMap[tp.GoType]; ok {
		tp.IsEnum = true
	}

	return tp
}

func convertGoTypeToTypescript(t SQLBoilerTableColumnType) string {
	goType := t.GoType

	if strings.HasPrefix(goType, "null.") {
		goType = strings.TrimPrefix(goType, "null.")
		goType = strings.ToLower(goType)
	}

	if goType == "string" {
		return "string"
	}

	if goType == "types.JSON" {
		return "any"
	}
	if goType == "time.Time" || goType == "time" {
		return "Date"
	}

	goType = strings.ReplaceAll(goType, "bool", "boolean")
	goType = strings.ReplaceAll(goType, "int", "number")
	goType = strings.ReplaceAll(goType, "int64", "number")
	goType = strings.ReplaceAll(goType, "int32", "number")
	goType = strings.ReplaceAll(goType, "float", "number")

	if strings.HasSuffix(goType, "Slice") {
		return fmt.Sprintf("%v[]", strings.TrimSuffix(goType, "Slice"))
	}

	if strings.HasPrefix(goType, "[]") {
		return fmt.Sprintf("%v[]", strings.TrimPrefix(goType, "[]"))
	}

	return goType
}
