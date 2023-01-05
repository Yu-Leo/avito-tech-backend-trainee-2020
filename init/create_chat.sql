INSERT INTO chats (name)
VALUES ('chat_name') RETURNING chats.id;

INSERT INTO users_chats (user_id, chat_id)
VALUES (1, 1);