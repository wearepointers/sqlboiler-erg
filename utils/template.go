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
		"getPrimaryKey": func(prefix string, columns []SQLBoilerTableColumn) string {
			var pk []string
			for _, column := range columns {
				if column.IsPrimaryKey {
					pk = append(pk, fmt.Sprint(prefix, column.Name.PascalCase))
				}
			}
			return strings.Join(pk, `+"_"+`)
		},
		"getCustomFieldsName": func() SQLBoilerName {
			return SQLBoilerName{
				PascalCase: "CustomFields",
				SnakeCase:  "custom_fields",
				CamelCase:  "customFields",
			}
		},
		"getStructTag": c.getStructTagFunc,
		"getTypescriptType": func(t SQLBoilerType, name SQLBoilerName) string {
			tsType := convertGoTypeToTypescript(t)
			formattedName := c.getStructTagFunc(name)

			if strings.HasPrefix(t.FormattedName, "*") {
				return fmt.Sprintf("%v?: %v", formattedName, tsType)
			}

			return fmt.Sprintf("%v: %v", formattedName, tsType)
		},
		"convertSQLBoilerToErgType": func(t SQLBoilerType, modelVar string, name string) string {
			modelVarName := fmt.Sprintf("%v.%v", modelVar, name)

			if t.IsEnum {
				return fmt.Sprintf("%v(%v)", t.OriginalName, modelVarName)
			}

			if strings.HasPrefix(t.OriginalName, "null.") || strings.HasPrefix(t.OriginalName, "types.") {
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

func (c *Config) getStructTagFunc(name SQLBoilerName) string {
	if c.sqlBoilerConfig.StructTagCasing != nil {
		if *c.sqlBoilerConfig.StructTagCasing == "snake" {
			return name.SnakeCase
		}

		if *c.sqlBoilerConfig.StructTagCasing == "camel" {
			return name.CamelCase
		}

		if *c.sqlBoilerConfig.StructTagCasing == "title" {
			fmt.Println("title casing not supported")
		}

		if *c.sqlBoilerConfig.StructTagCasing == "alias" {
			fmt.Println("pascal casing not supported")
		}
	}

	return name.SnakeCase
}
