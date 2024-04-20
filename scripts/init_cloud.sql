CREATE TABLE IF NOT EXISTS users
(
    id text,
    login text,
    password text,
    payload text,
    UNIQUE (login),
    UNIQUE (id)
);
