-- name: CreateOrder :one
INSERT INTO Orders (
    order_number, 
    product_name, 
    quantity, 
    product_id, 
    additional_shelf
) VALUES (
      $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetOrders :many
SELECT * FROM Orders;

-- name: GetOrderByNumber :one
SELECT * FROM Orders WHERE order_number = $1;


-- name: GetOrderByProductID :one
SELECT * FROM Orders WHERE product_id = $1;

-- name: UpdateOrder :one
UPDATE Orders
SET quantity = $2, additional_shelf = $3
WHERE order_id = $1
RETURNING *;

-- name: DeleteOrder :one
DELETE FROM Orders WHERE order_id = $1
RETURNING order_id;


-- name: GetOrderByGrouping :many
SELECT o.order_number,
       p.title,
       o.product_id,
       o.quantity,
       s.shelf_name AS primary_shelf,
       COALESCE(STRING_AGG(ps_false_shelves.shelf_name, ', '), '') AS additional_shelves
FROM Orders o
JOIN products p ON o.product_id = p.id
JOIN product_shelves ps ON p.id = ps.id AND is_primary = true
JOIN Shelves s ON ps.shelf_id = s.shelf_id
LEFT JOIN product_shelves ps_false ON p.id = ps_false.id AND ps_false.is_primary = false
LEFT JOIN Shelves ps_false_shelves ON ps_false.shelf_id = ps_false_shelves.shelf_id
WHERE o.order_number IN ($1)
GROUP BY o.order_number, p.title, o.product_id, o.quantity, s.shelf_name ORDER BY primary_shelf, product_id LIMIT 100;


-- SELECT o.order_number,
--        p.title,
--        o.product_id,
--        o.quantity,
--        s.shelf_name AS primary_shelf,
--        COALESCE(ARRAY_AGG(ps_false_shelves.shelf_name), '{}') AS additional_shelves
-- FROM Orders o
-- JOIN products p ON o.product_id = p.id
-- JOIN product_shelves ps ON p.id = ps.id AND is_primary = true
-- JOIN Shelves s ON ps.shelf_id = s.shelf_id
-- LEFT JOIN product_shelves ps_false ON p.id = ps_false.id AND ps_false.is_primary = false
-- LEFT JOIN Shelves ps_false_shelves ON ps_false.shelf_id = ps_false_shelves.shelf_id
-- WHERE o.order_number IN ($1)
-- GROUP BY o.order_number, p.title, o.product_id, o.quantity, s.shelf_name;
