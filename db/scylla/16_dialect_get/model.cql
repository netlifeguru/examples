--ListPostsByUser
SELECT *
FROM posts_by_user
WHERE user_id = ?
LIMIT ?

--GetUserByEmail
SELECT *
FROM users_by_email
WHERE email = ?
LIMIT 1

--CountPostsByUser
SELECT COUNT(*)
FROM posts_by_user
WHERE user_id = ?