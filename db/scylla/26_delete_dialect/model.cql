--InsertUser
INSERT INTO users_by_id (id, email, name, active, created_at) VALUES (?, ?, ?, ?, ?)

--UpdateUser
UPDATE users_by_id
SET name = ?, active = ?
WHERE id = ?

--DeleteUser
DELETE FROM users_by_id WHERE id = ?