CREATE TABLE IF NOT EXISTS users
(
    id text,
    login text,
    public_key text,
    UNIQUE (login),
    UNIQUE (id)
);


CREATE TABLE IF NOT EXISTS challenges
(
    id text,
    payload text,
    public_key text,
    user_login text,

    FOREIGN KEY (user_login) REFERENCES users(login)
);