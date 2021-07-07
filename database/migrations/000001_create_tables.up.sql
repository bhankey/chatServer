CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL PRIMARY KEY,
    username text NOT NULL UNIQUE,
    created_at timestamp NOT NULL,
    CHECK (length(username) >= 1)
);

CREATE TABLE IF NOT EXISTS chat (
    id serial NOT NULL PRIMARY KEY,
    name text NOT NULL UNIQUE,
    created_at timestamp NOT NULL,
    CHECK (length(name) >= 1)
);

CREATE TABLE IF NOT EXISTS chat_users (
    user_id integer REFERENCES users(id),
    chat_id integer REFERENCES chat(id),
    PRIMARY KEY (user_id, chat_id)
);

CREATE TABLE IF NOT EXISTS message (
    id serial NOT NULL PRIMARY KEY,
    userId serial NOT NULL REFERENCES users (id),
    chatId serial NOT NULL REFERENCES chat (id),
    text text NOT NULL,
    created_at timestamp NOT NULL,
    CHECK (length(text) >= 1)
);