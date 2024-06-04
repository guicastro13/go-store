-- name: CreateCategory :exec
INSERT INTO category (id, title, created_at, updated_at)
VALUES ($1, $2, $3, $4);

-- name: FindManyCategories :many
SELECT c.id, c.title
FROM category c
ORDER BY c.created_at DESC;
