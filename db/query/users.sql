-- name: InsertUser :execresult
INSERT INTO users (name, email) VALUES (?, ?);

-- name: UpdateUser :execresult
UPDATE users 
SET 
    name = COALESCE(?, name),
    email = COALESCE(?, email)
WHERE id = ?;
