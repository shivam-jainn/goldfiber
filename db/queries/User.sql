-- name: GetAllUsers :many
SELECT * FROM users WHERE id = $1;