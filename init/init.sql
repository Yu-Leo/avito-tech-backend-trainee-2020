DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS chats CASCADE;
DROP TABLE IF EXISTS users_chats;
DROP TABLE IF EXISTS messages;

CREATE TABLE users
(
    id         integer PRIMARY KEY UNIQUE,
    username   varchar(30) UNIQUE NOT NULL,
    created_at timestamp          NOT NULL DEFAULT NOW()
);


CREATE TABLE chats
(
    id         integer PRIMARY KEY UNIQUE,
    name       varchar(30) UNIQUE NOT NULL,
    created_at timestamp          NOT NULL DEFAULT NOW()
);


CREATE TABLE users_chats
(
    id      integer PRIMARY KEY UNIQUE,
    user_id integer REFERENCES users (id) NOT NULL,
    chat_id integer REFERENCES chats (id) NOT NULL
);


CREATE TABLE messages
(
    id         integer PRIMARY KEY UNIQUE,
    user_id    integer REFERENCES users (id) NOT NULL,
    chat_id    integer REFERENCES chats (id) NOT NULL,
    text       text,
    created_at timestamp                     NOT NULL DEFAULT NOW()
);