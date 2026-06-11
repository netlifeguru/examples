--ListUsers
SELECT *
FROM users
ORDER BY created_at DESC
    LIMIT ?

--GetUser
SELECT *
FROM users
WHERE id = ?
    LIMIT ?

--CountUsers
SELECT COUNT(*)
FROM users