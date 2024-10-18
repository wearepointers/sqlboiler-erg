// Generated by sqlboiler-erg: DO NOT EDIT.
package erg

import (
	"github.com/wearepointers/sqlboiler-erg/example/models/dm"
)

type TablesWithColumns map[string]map[string]bool

var TablesWithColumnsMap = TablesWithColumns{
	dm.TableNames.Account: map[string]bool{
		dm.AccountColumns.ID:        true,
		dm.AccountColumns.Email:     true,
		dm.AccountColumns.CreatedAt: true,
		dm.AccountColumns.CreatedBy: true,
		dm.AccountColumns.UpdatedAt: true,
		dm.AccountColumns.DeletedAt: true,
	},
	dm.TableNames.Causer: map[string]bool{
		dm.CauserColumns.ID:              true,
		dm.CauserColumns.AccountID:       true,
		dm.CauserColumns.SystemAccountID: true,
		dm.CauserColumns.CauserType:      true,
	},
	dm.TableNames.SystemAccount: map[string]bool{
		dm.SystemAccountColumns.ID:        true,
		dm.SystemAccountColumns.CreatedAt: true,
		dm.SystemAccountColumns.UpdatedAt: true,
		dm.SystemAccountColumns.DeletedAt: true,
	},
}

func TableHasColumn(tableName, columnName string) bool {
	if _, ok := TablesWithColumnsMap[tableName]; !ok {
		return false
	}

	if _, ok := TablesWithColumnsMap[tableName][columnName]; !ok {
		return false
	}

	return true
}
