package utils

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"
	"text/template"

	"github.com/expanse-agency/sqlboiler-erg/templates"
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
		"getTypescriptType": func(t SQLBoilerType, name string) string {
			tsType := convertGoTypeToTypescript(t)

			if strings.HasPrefix(t.FormattedName, "*") {
				return fmt.Sprintf("%v?: %v", name, tsType)
			}

			return fmt.Sprintf("%v: %v", name, tsType)
		},
		"convertSQLBoilerToErgType": func(t SQLBoilerType, modelVar string, name string) string {
			modelVarName := fmt.Sprintf("%v.%v", modelVar, name)

			if t.IsEnum {
				return fmt.Sprintf("%v(%v)", t.OriginalName, modelVarName)
			}

			if strings.HasPrefix(t.OriginalName, "null.") {
				return modelVarName + ".Ptr()"
			}

			if strings.HasPrefix(t.OriginalName, "types.") {
				s := strings.ReplaceAll(t.OriginalName, ".", "Dot")
				sn := titleize(strings.TrimPrefix(t.FormattedName, "[]"))
				if strings.HasPrefix(t.FormattedName, "[]") {
					sn = fmt.Sprintf("%vSlice", sn)
				}
				s = fmt.Sprintf("%vTo%v(%v)", s, sn, modelVarName)
				return s
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
