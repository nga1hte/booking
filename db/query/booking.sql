-- name: CreateBooking :one
INSERT INTO bookings (
    "bookedBy",
    "bookStarts",
    "bookEnds"
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetBooking :one
SELECT * FROM bookings
WHERE "bookingId" = $1 LIMIT 1;

-- name: GetBookings :many
SELECT * FROM bookings
ORDER BY "bookStarts"
LIMIT $1
OFFSET $2;

-- name: GetUserBookings :many
SELECT * FROM bookings
WHERE "bookedBy" = $1
ORDER BY "bookStarts"
LIMIT $2
OFFSET $3;

-- name: DeleteBooking :exec
DELETE FROM bookings
WHERE "bookingId" = $1;

-- name: GetBookingsFromToday :many
SELECT * FROM bookings
WHERE "bookStarts" > $1;
