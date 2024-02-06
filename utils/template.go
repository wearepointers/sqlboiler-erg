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
		"pluralize": func(s string) string {
			return pluralize(s)
		},
		"singularize": func(s string) string {
			return singularize(s)
		},

		"convertToNullFuncType": func(s string) string {
			return titleize(strings.TrimPrefix(strings.TrimPrefix(s, "*"), "[]"))
		},
		"convertToFuncType": func(t SQLBoilerTableColumnType, pfx, n string) string {
			n = fmt.Sprint(pfx, n)
			if t.GoType == "time.Time" {
				return fmt.Sprintf("ConvertTime(%v)", n)
			}

			if t.GoType == "json" {
				return fmt.Sprintf("json(%v)", n)
			}

			if t.IsEnum {
				return fmt.Sprintf("%v(%v)", t.GoType, n)
			}

			if t.GoTypeName == "bool" {
				return fmt.Sprintf("&%v", n)
			}

			return n
		}},
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
