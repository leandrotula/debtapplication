ALTER TABLE IF EXISTS users
ADD
    COLUMN active boolean not null default false