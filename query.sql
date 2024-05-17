-- name: GetAuthor :one
SELECT * FROM author
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM author
ORDER BY name;

-- name: CreateAuthor :exec
INSERT INTO author (
  name, bio
) VALUES (
  $1, $2
);

-- name: UpdateAuthor :exec
UPDATE author
  set name = $2,
  bio = $3
WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM author
WHERE id = $1;