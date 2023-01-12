SELECT chats.id, chats.name, chats.created_at
FROM users_chats
JOIN chats on users_chats.chat_id = chats.id
WHERE users_chats.user_id = 1
ORDER BY (SELECT (MAX(created_at))
          FROM messages
          WHERE chat_id = users_chats.chat_id
          GROUP BY chat_id) DESC;


SELECT users_chats.user_id
FROM users_chats
WHERE users_chats.chat_id = 1;
