-- db/queries/product.sql

-- Create a new product
-- name: CreateProduct :one
INSERT INTO product (name, category, price, image_url)
VALUES ($1, $2, $3, $4)
RETURNING product_id, name, category, price, image_url, created_at, updated_at;

-- Get all products with pagination
-- name: GetAllProductsWithPagination :many
SELECT product_id, name, category, price, image_url, created_at, updated_at
FROM product
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- Delete a product
-- name: DeleteProduct :exec
DELETE FROM product
WHERE product_id = $1;

-- Get a product by id
-- name: GetProductById :one
SELECT product_id, name, category, price, image_url, created_at, updated_at
FROM product
WHERE product_id = $1;

-- Get Products by category
-- name: GetProductsByCategory :many
SELECT product_id, name, category, price, image_url, created_at, updated_at
FROM product
WHERE category = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- Get Total Products by category
-- name: GetTotalProductsByCategory :one
SELECT COUNT(*)
FROM product
WHERE category = $1;