DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS chats CASCADE;
DROP TABLE IF EXISTS users_chats;
DROP TABLE IF EXISTS messages;

CREATE TABLE users
(
    id         bigserial PRIMARY KEY UNIQUE,
    username   varchar(30) UNIQUE NOT NULL,
    created_at timestamp          NOT NULL DEFAULT now()
);


CREATE TABLE chats
(
    id         bigserial PRIMARY KEY UNIQUE,
    name       varchar(30) UNIQUE NOT NULL,
    created_at timestamp          NOT NULL DEFAULT now()
);


CREATE TABLE users_chats
(
    id      bigserial PRIMARY KEY UNIQUE,
    user_id bigserial REFERENCES users (id) NOT NULL,
    chat_id bigserial REFERENCES chats (id) NOT NULL
);


CREATE TABLE messages
(
    id           bigserial PRIMARY KEY UNIQUE,
    user_id      bigserial REFERENCES users (id) NOT NULL,
    chat_id      bigserial REFERENCES chats (id) NOT NULL,
    message_text text,
    created_at   timestamp                       NOT NULL DEFAULT now()
);