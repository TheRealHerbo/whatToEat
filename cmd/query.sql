-- name: GetRecipie :one
SELECT * FROM recipie
WHERE id = ? LIMIT 1;

-- name: ListRecipies :many
SELECT * FROM recipie
ORDER BY name;

-- name: CreateRecipie :one
INSERT INTO recipie (
  name, ingredients, directions
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateRecipie :exec
UPDATE recipie
set name = ?,
ingredients = ?,
directions = ?
WHERE id = ?;

-- name: DeleteRecipie :exec
DELETE FROM recipie
WHERE id = ?;