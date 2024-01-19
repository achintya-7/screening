-- name: InsertUser :execresult
INSERT INTO users (name, email) VALUES ($1, $2);

-- name: UpdateUser :execresult
UPDATE users 
SET 
    name = COALESCE($2, name),
    email = COALESCE($3, email)
WHERE id = $1;
