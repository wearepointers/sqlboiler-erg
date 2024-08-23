package utils

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"
	"text/template"

	"github.com/wearepointers/sqlboiler-erg/templates"
)

func (c *Config) writeTemplate(input, output string, data any) error {
	content, err := templates.Builtin.ReadFile(input)
	if err != nil {
		return err
	}

	template, err := c.parseTemplate(string(content), data, strings.HasSuffix(output, ".go"))
	if err != nil {
		return err
	}

	hF, err := os.Create(output)
	if err != nil {
		return err
	}
	defer hF.Close()

	if _, err = hF.WriteString(template); err != nil {
		return err
	}

	return nil
}

func (c *Config) parseTemplate(tmplte string, data any, shouldFormat bool) (string, error) {
	tpl, err := template.New("").Funcs(template.FuncMap{
		"sqlboilerPkgName": func() string {
			return c.sqlBoilerConfig.PkgName
		},
		"ergPkgName": func() string {
			return c.sqlBoilerConfig.Erg.Pkgname
		},
		"isInline": func() bool {
			return c.sqlBoilerConfig.Erg.Inline
		},
		"pluralize": func(s string) string {
			return pluralize(s)
		},
		"singularize": func(s string) string {
			return singularize(s)
		},
		"shouldOmit": func(tp SQLBoilerType, s string) string {
			if tp.IsNullable {
				return s
			}
			return ""
		},
		"getPrimaryKey": func(prefix string, columns []SQLBoilerTableColumn) string {
			var pk []string
			for _, column := range columns {
				if column.IsPrimaryKey {
					pk = append(pk, fmt.Sprint(prefix, column.Name.PascalCase))
				}
			}
			return strings.Join(pk, `+"_"+`)
		},
		"getStructTag": c.getStructTag,
		"getTypescriptType": func(t SQLBoilerType, name SQLBoilerName) string {
			tsType := convertGoTypeToTypescript(t)
			formattedName := c.getStructTag(name, "json")

			if t.IsNullable {
				return fmt.Sprintf("%v?: %v", formattedName, tsType)
			}

			return fmt.Sprintf("%v: %v", formattedName, tsType)
		},
		"convertSQLBoilerToErgType": func(t SQLBoilerType, modelVar string, name string) string {
			modelVarName := fmt.Sprintf("%v.%v", modelVar, name)

			if strings.HasPrefix(t.OriginalName, "null.") || strings.HasPrefix(t.OriginalName, "types.") || strings.HasPrefix(t.OriginalName, "decimal.") {
				fromName := strings.ReplaceAll(t.OriginalName, ".", "Dot")

				toNameS := strings.Split(t.OriginalName, ".")
				toNamePrefix := toNameS[0]
				toName := toNameS[1]

				if toNamePrefix == "null" {
					toName = fmt.Sprintf("%vPtr", toName)
				}

				if strings.HasSuffix(toName, "Array") {
					toName = strings.ReplaceAll(toName, "Array", "Slice")
				}

				return fmt.Sprintf("%vTo%v(%v)", fromName, toName, modelVarName)

			}

			return modelVarName
		},
	},
	).Parse(tmplte)
	if err != nil {
		return "", err
	}

	var content bytes.Buffer
	err = tpl.Execute(&content, data)
	if err != nil {
		return "", fmt.Errorf("execute: %v", err)
	}

	if shouldFormat {
		contentBytes := content.Bytes()
		formattedContent, err := format.Source(contentBytes)
		if err != nil {
			return string(contentBytes), fmt.Errorf("formatting: %v", err)
		}

		return string(formattedContent), nil
	}

	return content.String(), nil
}

func (c *Config) getStructTag(name SQLBoilerName, fieldTag string) string {
	switch fieldTag {
	case "json":
		return c.getNameCasing(name, c.sqlBoilerConfig.StructTagCases.Json)
	case "yaml":
		return c.getNameCasing(name, c.sqlBoilerConfig.StructTagCases.Yaml)
	case "toml":
		return c.getNameCasing(name, c.sqlBoilerConfig.StructTagCases.Toml)
	case "boil":
		return c.getNameCasing(name, c.sqlBoilerConfig.StructTagCases.Boil)
	default:
		fmt.Println("Unknown field tag. Using snake case.")
		return c.getNameCasing(name, TagCaseSnake)
	}
}

func (c *Config) getNameCasing(name SQLBoilerName, casing TagCase) string {
	switch casing {
	case TagCaseCamel:
		return name.CamelCase
	case TagCaseSnake:
		return name.SnakeCase
	case TagCaseTitle:
		return name.PascalCase
	case TagCaseAlias:
		fmt.Println("Alias is not supported. Using snake case.")
		return name.SnakeCase
	default:
		fmt.Println("Unknown casing. Using snake case.")
		return name.SnakeCase
	}
}
