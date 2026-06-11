--InsertUser
INSERT INTO users (name, email, active)
VALUES (?, ?, ?)

--UpdateUser
UPDATE users
SET name = ?, email = ?, active = ?
WHERE id = ?

--DeleteUser
DELETE FROM users WHERE id = ?