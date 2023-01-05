DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS chats CASCADE;
DROP TABLE IF EXISTS users_chats;
DROP TABLE IF EXISTS messages;

CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY UNIQUE,
    username   VARCHAR(30) UNIQUE NOT NULL,
    created_at TIMESTAMP          NOT NULL DEFAULT NOW()
);


CREATE TABLE chats
(
    id         BIGSERIAL PRIMARY KEY UNIQUE,
    name       VARCHAR(30) UNIQUE NOT NULL,
    created_at TIMESTAMP          NOT NULL DEFAULT NOW()
);


CREATE TABLE users_chats
(
    id      BIGSERIAL PRIMARY KEY UNIQUE,
    user_id BIGSERIAL REFERENCES users (id) NOT NULL,
    chat_id BIGSERIAL REFERENCES chats (id) NOT NULL
);


CREATE TABLE messages
(
    id         BIGSERIAL PRIMARY KEY UNIQUE,
    user_id    BIGSERIAL REFERENCES users (id) NOT NULL,
    chat_id    BIGSERIAL REFERENCES chats (id) NOT NULL,
    text       TEXT,
    created_at TIMESTAMP                       NOT NULL DEFAULT NOW()
);