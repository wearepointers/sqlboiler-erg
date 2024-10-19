// Generated by sqlboiler-erg: DO NOT EDIT.
package erg

import (
	"github.com/wearepointers/sqlboiler-erg/example/models/dm"
	"time"
)

type SystemAccount struct {
	ID        string     `json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time  `json:"createdAt" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" toml:"updated_at" yaml:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" toml:"deleted_at" yaml:"deleted_at"`

	Causers CauserSlice `json:"causers,omitempty" toml:"causers" yaml:"causers"`
}

type SystemAccountSlice []*SystemAccount

func ToSystemAccounts(a dm.SystemAccountSlice, exclude ...string) SystemAccountSlice {
	s := make(SystemAccountSlice, len(a))
	for i, d := range a {
		s[i] = ToSystemAccount(d, exclude...)
	}
	return s
}

func ToSystemAccount(a *dm.SystemAccount, exclude ...string) *SystemAccount {
	p := SystemAccount{
		ID:        a.ID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		DeletedAt: nullDotTimeToTimePtr(a.DeletedAt),
	}

	if a.R != nil {
		if a.R.Causers != nil && doesNotContain(exclude, "system_account.causer") {
			p.Causers = ToCausers(a.R.Causers, append(exclude, "causer.system_account")...)
		}
	}

	return &p
}