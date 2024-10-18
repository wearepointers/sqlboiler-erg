// Generated by sqlboiler-erg: DO NOT EDIT.
export interface AccountRelations {
    createdByCauser?: Causer;
    causers?: Causer[];
}

export interface Account extends AccountRelations {
    id: string;
    email: string;
    createdAt: Date;
    createdBy: string;
    updatedAt: Date;
    deletedAt?: Date;
}

export interface CauserRelations {
    account?: Account;
    systemAccount?: SystemAccount;
    createdByAccounts?: Account[];
}

export interface Causer extends CauserRelations {
    id: string;
    accountId?: string;
    systemAccountId?: string;
    causerType: string;
}

export interface SystemAccountRelations {
    causers?: Causer[];
}

export interface SystemAccount extends SystemAccountRelations {
    id: string;
    createdAt: Date;
    updatedAt: Date;
    deletedAt?: Date;
}
