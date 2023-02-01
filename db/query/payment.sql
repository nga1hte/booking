-- name: CreatePayment :one
INSERT INTO payments (
    "bookingId"
) VALUES (
    $1
)RETURNING *;

-- name: GetBookingPayment :one
SELECT * FROM payments
WHERE "bookingId" = $1;

-- name: GetPayments :many
SELECT * FROM payments
ORDER BY "paymentDate"
LIMIT $1
OFFSET $2;