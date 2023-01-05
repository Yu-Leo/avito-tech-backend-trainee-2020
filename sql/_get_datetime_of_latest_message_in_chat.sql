SELECT (MAX(created_at)) FROM messages
WHERE chat_id = 1
GROUP BY chat_id;
