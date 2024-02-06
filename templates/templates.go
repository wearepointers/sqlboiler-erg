// Package templates is an empty package strictly for embedding sqlboiler
// default templates.
package templates

import "embed"

// Builtin sqlboiler-erg templates
//go:embed main test
var Builtin embed.FS
