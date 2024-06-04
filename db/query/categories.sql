-- name: GetCategoryByName :one
SELECT * FROM categories WHERE name = $1 LIMIT 1;

-- name: GetCategoryIDByName :one
SELECT id FROM categories WHERE name = $1 LIMIT 1;

-- name: GetCategoryByID :one
SELECT * FROM categories WHERE id = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM categories;

-- name: ListCategoryNamesToIDs :many
SELECT name, id FROM categories 
WHERE name = ANY($1)
;
