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

func getTypeFromFieldType(fieldType ast.Expr) SQLBoilerType {
	var tp SQLBoilerType
	switch t := fieldType.(type) {
	case *ast.Ident:
		tp = SQLBoilerType{
			OriginalName:  t.Name,
			FormattedName: t.Name,
		}
	case *ast.StarExpr:
		ft := getTypeFromFieldType(t.X)

		tp = SQLBoilerType{
			OriginalName:  ft.OriginalName,
			FormattedName: "*" + ft.OriginalName,
		}
	case *ast.SelectorExpr:
		ft := getTypeFromFieldType(t.X)

		tp = SQLBoilerType{
			OriginalName:  ft.OriginalName + "." + t.Sel.Name,
			FormattedName: ft.OriginalName + "." + t.Sel.Name,
		}
	default:
		fmt.Println("unknown", t)
	}

	tp.FormattedName, tp.IsNullable = sqlboilerTypeToType(tp.FormattedName)

	if _, ok := enumCacheMap[tp.OriginalName]; ok {
		tp.IsEnum = true
	}

	if _, ok := enumCacheMap[strings.TrimPrefix(tp.FormattedName, "*")]; ok {
		tp.IsEnum = true

		if tp.IsNullable {
			tp.FormattedName = "dm.Null" + strings.TrimPrefix(tp.FormattedName, "*")
		}

		if !tp.IsNullable {
			tp.FormattedName = "dm." + tp.FormattedName
		}
	}

	return tp
}

var sqlboilerTypes = map[string]string{
	"time":  "time.Time",
	"bytes": "[]byte",
}

func sqlboilerTypeToType(s string) (string, bool) {
	var formattedString = s
	var isNullable = false
	var isSlice = false

	if strings.Contains(formattedString, "time") {
		modelImports = append(modelImports, "time")
	}

	if strings.HasPrefix(formattedString, "null.") {
		isNullable = true
		formattedString = strings.ToLower(strings.TrimPrefix(formattedString, "null."))
	}

	if strings.HasPrefix(formattedString, "types.") {
		formattedString = strings.ToLower(strings.TrimPrefix(formattedString, "types."))

		if strings.HasPrefix(formattedString, "null") {
			isNullable = true
			formattedString = strings.TrimPrefix(formattedString, "null")
		}
	}

	if strings.HasPrefix(formattedString, "decimal.") {
		formattedString = strings.ToLower(strings.TrimPrefix(formattedString, "decimal."))
	}

	if strings.Contains(formattedString, "array") {
		isSlice = true
		formattedString = strings.TrimSuffix(formattedString, "array")
	}

	if strings.HasPrefix(formattedString, "Null") {
		isNullable = true
		formattedString = strings.TrimPrefix(formattedString, "Null")
	}

	if val, ok := sqlboilerTypes[formattedString]; ok {
		formattedString = val
	}

	if isNullable {
		formattedString = "*" + formattedString
	}

	if isSlice {
		formattedString = "[]" + formattedString
	}

	return formattedString, isNullable
}

func convertGoTypeToTypescript(t SQLBoilerType) string {
	var formattedString = t.FormattedName

	formattedString = strings.TrimPrefix(formattedString, "*")

	if strings.Contains(formattedString, "int") || strings.Contains(formattedString, "float") || strings.Contains(formattedString, "decimal") {
		formattedString = "number"
	}

	formattedString = strings.ReplaceAll(formattedString, "bool", "boolean")
	formattedString = strings.ReplaceAll(formattedString, "time.Time", "Date")
	formattedString = strings.ReplaceAll(formattedString, "byte", "any")
	formattedString = strings.ReplaceAll(formattedString, "json", "any")

	if strings.HasSuffix(formattedString, "Slice") {
		return fmt.Sprintf("%v[]", strings.TrimSuffix(formattedString, "Slice"))
	}

	if strings.HasPrefix(formattedString, "[]") {
		for strings.HasPrefix(formattedString, "[]") {
			formattedString = strings.TrimPrefix(formattedString, "[]")
			formattedString += "[]"
		}
	}

	if t.IsEnum {
		formattedString = strings.TrimPrefix(formattedString, "dm.")

		if t.IsNullable {
			formattedString = strings.TrimPrefix(formattedString, "Null")
		}
	}

	return formattedString

}
