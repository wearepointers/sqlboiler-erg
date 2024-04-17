DROP SCHEMA IF EXISTS public CASCADE; 
CREATE SCHEMA public; 
CREATE TABLE schema_migrations (
    version VARCHAR(255) NOT NULL,
    dirty boolean NOT NULL DEFAULT false,
    CONSTRAINT schema_migrations_pk PRIMARY KEY (version)
);