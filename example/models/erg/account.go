// Generated by sqlboiler-erg: DO NOT EDIT.
package erg

import (
	"github.com/wearepointers/sqlboiler-erg/example/models/dm"
	"time"
)

type Account struct {
	ID        string     `json:"id" toml:"id" yaml:"id"`
	Email     string     `json:"email" toml:"email" yaml:"email"`
	CreatedAt time.Time  `json:"createdAt" toml:"created_at" yaml:"created_at"`
	CreatedBy string     `json:"createdBy" toml:"created_by" yaml:"created_by"`
	UpdatedAt time.Time  `json:"updatedAt" toml:"updated_at" yaml:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" toml:"deleted_at" yaml:"deleted_at"`

	CreatedByCauser *Causer     `json:"createdByCauser,omitempty" toml:"created_by_causer" yaml:"created_by_causer"`
	Causers         CauserSlice `json:"causers,omitempty" toml:"causers" yaml:"causers"`
}

type AccountSlice []*Account

func ToAccounts(a dm.AccountSlice, exclude ...string) AccountSlice {
	s := make(AccountSlice, len(a))
	for i, d := range a {
		s[i] = ToAccount(d, exclude...)
	}
	return s
}

func ToAccount(a *dm.Account, exclude ...string) *Account {
	p := Account{
		ID:        a.ID,
		Email:     a.Email,
		CreatedAt: a.CreatedAt,
		CreatedBy: a.CreatedBy,
		UpdatedAt: a.UpdatedAt,
		DeletedAt: nullDotTimeToTimePtr(a.DeletedAt),
	}

	if a.R != nil {
		if a.R.CreatedByCauser != nil && doesNotContain(exclude, "account.causer") {
			p.CreatedByCauser = ToCauser(a.R.CreatedByCauser, append(exclude, "causer.account")...)
		}
		if a.R.Causers != nil && doesNotContain(exclude, "account.causer") {
			p.Causers = ToCausers(a.R.Causers, append(exclude, "causer.account")...)
		}
	}

	return &p
}
