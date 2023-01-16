DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS chats CASCADE;
DROP TABLE IF EXISTS users_chats;
DROP TABLE IF EXISTS messages;

CREATE TABLE users
(
    id         serial UNIQUE,
    username   varchar(30) UNIQUE NOT NULL,
    created_at timestamp          NOT NULL DEFAULT now(),
    PRIMARY KEY (id)
);


CREATE TABLE chats
(
    id         serial UNIQUE,
    name       varchar(30) UNIQUE NOT NULL,
    created_at timestamp          NOT NULL DEFAULT now(),
    PRIMARY KEY (id)
);


CREATE TABLE users_chats
(
    id      serial UNIQUE,
    user_id serial NOT NULL,
    chat_id serial NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE
);


CREATE TABLE messages
(
    id           serial UNIQUE,
    user_id      serial    NOT NULL,
    chat_id      serial    NOT NULL,
    message_text text,
    created_at   timestamp NOT NULL DEFAULT now(),
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE
);