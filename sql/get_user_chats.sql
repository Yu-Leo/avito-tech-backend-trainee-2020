SELECT *
FROM users_chats
WHERE users_chats.user_id = 1
ORDER BY (SELECT (MAX(created_at))
          FROM messages
          WHERE chat_id = users_chats.chat_id
          GROUP BY chat_id) DESC;
