package utils

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"golang.org/x/mod/modfile"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "sqlboiler.toml", "Filename of config file to override default lookup")
	flag.StringVar(&configFile, "c", "sqlboiler.toml", "Filename of config file to override default lookup")
	flag.Parse()
}

type Config struct {
	sqlBoilerConfig *SQLBoilerConfig
	modFile         *modfile.File
}

func ParseConfig() (*Config, error) {
	sqlBoilerConfig, err := parseSQLBoilerConfig()
	if err != nil {
		return nil, err
	}

	fmt.Println("SQLBoilerConfig", sqlBoilerConfig.Types)

	modFile, err := parseModFile()
	if err != nil {
		return nil, err
	}

	return &Config{
		sqlBoilerConfig: sqlBoilerConfig,
		modFile:         modFile,
	}, nil
}

type ERGConfig struct {
	Output    string   `toml:"output"`
	OutputTS  string   `toml:"output-ts"`
	Pkgname   string   `toml:"pkgname"`
	Wipe      bool     `toml:"wipe"`
	Inline    bool     `toml:"inline"`
	Blacklist []string `toml:"blacklist"`
}

type TagCase string

const (
	TagCaseCamel TagCase = "camel"
	TagCaseSnake TagCase = "snake"
	TagCaseTitle TagCase = "title"
	TagCaseAlias TagCase = "alias"
)

type StructTagCases struct {
	Json TagCase `toml:"json,omitempty" json:"json,omitempty"`
	Yaml TagCase `toml:"yaml,omitempty" json:"yaml,omitempty"`
	Toml TagCase `toml:"toml,omitempty" json:"toml,omitempty"`
	Boil TagCase `toml:"boil,omitempty" json:"boil,omitempty"`
}

type SQLBoilerConfig struct {
	Output         string         `toml:"output"`
	PkgName        string         `toml:"pkgname"`
	StructTagCases StructTagCases `toml:"struct-tag-cases"`
	Types          TypesInfo      `toml:"types"`
	Erg            ERGConfig      `toml:"erg"`
}

type TypesInfo struct {
	Match   TypeInfo   `toml:"match"`
	Replace TypeInfo   `toml:"replace"`
	Imports ImportInfo `toml:"imports"`
}

type TypeInfo struct {
	Type string `toml:"type"`
}

type ImportInfo struct {
	ThirdParty []string `toml:"third_party"`
}

func parseSQLBoilerConfig() (*SQLBoilerConfig, error) {
	var sqlBoilerConfig SQLBoilerConfig
	if _, err := toml.DecodeFile(configFile, &sqlBoilerConfig); err != nil {
		return nil, err
	}

	if sqlBoilerConfig.StructTagCases.Boil == "" {
		sqlBoilerConfig.StructTagCases.Boil = TagCaseSnake
	}

	if sqlBoilerConfig.StructTagCases.Json == "" {
		sqlBoilerConfig.StructTagCases.Json = TagCaseSnake
	}

	if sqlBoilerConfig.StructTagCases.Yaml == "" {
		sqlBoilerConfig.StructTagCases.Yaml = TagCaseSnake
	}

	if sqlBoilerConfig.StructTagCases.Toml == "" {
		sqlBoilerConfig.StructTagCases.Toml = TagCaseSnake
	}

	if sqlBoilerConfig.Erg.Output == "" {
		sqlBoilerConfig.Erg.Output = "erg_models"
	}

	if sqlBoilerConfig.Erg.Pkgname == "" {
		sqlBoilerConfig.Erg.Pkgname = "erg"
	}

	return &sqlBoilerConfig, nil
}

func parseModFile() (*modfile.File, error) {
	b, err := os.ReadFile("go.mod")
	if err != nil {
		return nil, err
	}

	f, err := modfile.Parse("go.mod", b, nil)
	if err != nil {
		return nil, err
	}

	return f, nil
}
