DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'type_sex') THEN
CREATE TYPE type_sex AS ENUM ('male', 'female');
END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'type_wool') THEN
CREATE TYPE type_wool AS ENUM ('short', 'long');
END IF;
END$$;

CREATE TABLE IF NOT EXISTS cat(
    id          serial
        constraint cat_pk
            primary key,
    name        varchar(50)        not null,
    description varchar(1000),
    breed       varchar(50),
    birthday    date               not null,
    sex         type_sex           not null,
    tail_length int  default 25    not null,
    color       varchar(50)        not null,
    wool_type   type_wool          not null,
    is_chipped  bool default false not null,
    weight      decimal,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

comment on table cat is 'table to store cats';
comment on column cat.name is 'type of cat breed';

create unique index cat_name_uindex
            on cat (name);

create unique index cat_name_uindex_2
        on cat (name);