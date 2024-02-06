package utils

import (
	"os"
	"path"
	"strings"
)

func (c *Config) getSQLBoilerModelsImportPath() string {
	return path.Join(c.modFile.Module.Mod.Path, c.sqlBoilerConfig.Output)
}

func (c *Config) getSQLBoilerModelsFilePath(file string) string {
	return path.Join(c.sqlBoilerConfig.Output, file)
}

func (c *Config) getERGDefaultImports() []string {
	return []string{c.getSQLBoilerModelsImportPath()}
}

func (c *Config) isTSEnabled() bool {
	return c.sqlBoilerConfig.Erg.OutputTS != "" || !strings.HasSuffix(c.sqlBoilerConfig.Erg.OutputTS, ".ts")
}

func (c *Config) isBlackListed(tn, cn string) bool {
	for _, b := range c.sqlBoilerConfig.Erg.Blacklist {
		if b == tn {
			return true
		}

		if strings.ContainsAny(b, "*.") {
			if cn == strings.ReplaceAll(b, "*.", "") {
				return true
			}
		}

		if strings.ContainsAny(b, ".") {
			if b == tn+"."+cn {
				return true
			}
		}
	}

	return false
}

func (c *Config) ConvertSQLBoilerModelsToApiModels() error {
	tables, enums, err := c.getSQLBoilerTablesAndEnums()
	if err != nil {
		return err
	}

	if err := c.wipe(); err != nil {
		return err
	}

	if err := c.writeERGHelperFunctionsToFile(); err != nil {
		return err
	}

	if err := c.writeSQLBoilerEnumsToERGFiles(enums); err != nil {
		return err
	}

	if err := c.writeSQLBoilerTablesToERGFiles(tables); err != nil {
		return err
	}

	if err := c.writeSQLBoilerTablesToERGFile(tables); err != nil {
		return err
	}

	if err := c.writeSQLBoilerTablesToTypeScriptFiles(tables, enums); err != nil {
		return err
	}

	return nil
}

func (c *Config) wipe() error {
	if !c.sqlBoilerConfig.Erg.Wipe {
		return nil
	}

	if err := os.RemoveAll(c.sqlBoilerConfig.Erg.Output); err != nil {
		return err
	}

	if err := os.MkdirAll(c.sqlBoilerConfig.Erg.Output, 0755); err != nil {
		return err
	}

	return nil
}
