--
--
--
-- UUID Extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--
--
--
-- Account Table
CREATE TABLE account (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by UUID NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,
    CONSTRAINT account_pk PRIMARY KEY (id),
    CONSTRAINT account_ak_id UNIQUE (id) NOT DEFERRABLE INITIALLY IMMEDIATE,
    CONSTRAINT account_ak_email UNIQUE (email,deleted_at) NOT DEFERRABLE INITIALLY IMMEDIATE
);
CREATE INDEX account_idx_id ON account (id);
CREATE INDEX account_idx_email ON account (email);

CREATE TABLE system_account (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,
    CONSTRAINT system_user_pk PRIMARY KEY (id),
    CONSTRAINT system_user_ak_id UNIQUE (id) NOT DEFERRABLE INITIALLY IMMEDIATE
);

CREATE TABLE causer (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    account_id UUID NULL,
    system_account_id UUID NULL,
    causer_type VARCHAR(255) NOT NULL,
    CONSTRAINT causer_pk PRIMARY KEY (id),
    CONSTRAINT causer_ak_id UNIQUE (id) NOT DEFERRABLE INITIALLY IMMEDIATE,
    CONSTRAINT causer_fk_account_id FOREIGN KEY (account_id) REFERENCES account (id) NOT DEFERRABLE INITIALLY IMMEDIATE,
    CONSTRAINT causer_fk_system_account_id FOREIGN KEY (system_account_id) REFERENCES system_account (id) NOT DEFERRABLE INITIALLY IMMEDIATE
);

ALTER TABLE account ADD CONSTRAINT account_fk_created_by FOREIGN KEY (created_by) REFERENCES causer (id) NOT DEFERRABLE INITIALLY IMMEDIATE;
