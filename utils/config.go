package utils

import (
	"flag"
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
	Blacklist []string `toml:"blacklist"`
}

type SQLBoilerConfig struct {
	Output  string    `toml:"output"`
	PkgName string    `toml:"pkgname"`
	Erg     ERGConfig `toml:"erg"`
}

func parseSQLBoilerConfig() (*SQLBoilerConfig, error) {
	var sqlBoilerConfig SQLBoilerConfig
	if _, err := toml.DecodeFile(configFile, &sqlBoilerConfig); err != nil {
		return nil, err
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
