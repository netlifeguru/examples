--ListUsers
SELECT *
FROM users
ORDER BY created_at DESC
    LIMIT ?

--GetUser
SELECT *
FROM users
WHERE id = ?
    LIMIT 1

--CountUsers
SELECT COUNT(*)
FROM users