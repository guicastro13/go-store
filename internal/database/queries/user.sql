-- name: GetUsersByID :one
SELECT * from users u where u.id = $1;
