--ListUsers
SELECT *
FROM users
ORDER BY created_at DESC
    LIMIT ?

--GetUser
SELECT *
FROM users
WHERE id = ?

--CountUsers
SELECT COUNT(*)
FROM users