// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: booking.sql

package db

import (
	"context"
	"time"
)

const createBooking = `-- name: CreateBooking :one
INSERT INTO bookings (
    "bookedBy",
    "bookStarts",
    "bookEnds"
) VALUES (
    $1, $2, $3
) RETURNING "bookingId", "bookedBy", "bookOn", "bookStarts", "bookEnds"
`

type CreateBookingParams struct {
	BookedBy   int64     `json:"bookedBy"`
	BookStarts time.Time `json:"bookStarts"`
	BookEnds   time.Time `json:"bookEnds"`
}

func (q *Queries) CreateBooking(ctx context.Context, arg CreateBookingParams) (Booking, error) {
	row := q.db.QueryRowContext(ctx, createBooking, arg.BookedBy, arg.BookStarts, arg.BookEnds)
	var i Booking
	err := row.Scan(
		&i.BookingId,
		&i.BookedBy,
		&i.BookOn,
		&i.BookStarts,
		&i.BookEnds,
	)
	return i, err
}

const deleteBooking = `-- name: DeleteBooking :exec
DELETE FROM bookings
WHERE "bookingId" = $1
`

func (q *Queries) DeleteBooking(ctx context.Context, bookingid int64) error {
	_, err := q.db.ExecContext(ctx, deleteBooking, bookingid)
	return err
}

const getBooking = `-- name: GetBooking :one
SELECT "bookingId", "bookedBy", "bookOn", "bookStarts", "bookEnds" FROM bookings
WHERE "bookingId" = $1 LIMIT 1
`

func (q *Queries) GetBooking(ctx context.Context, bookingid int64) (Booking, error) {
	row := q.db.QueryRowContext(ctx, getBooking, bookingid)
	var i Booking
	err := row.Scan(
		&i.BookingId,
		&i.BookedBy,
		&i.BookOn,
		&i.BookStarts,
		&i.BookEnds,
	)
	return i, err
}

const getBookings = `-- name: GetBookings :many
SELECT "bookingId", "bookedBy", "bookOn", "bookStarts", "bookEnds" FROM bookings
ORDER BY "bookStarts"
LIMIT $1
OFFSET $2
`

type GetBookingsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetBookings(ctx context.Context, arg GetBookingsParams) ([]Booking, error) {
	rows, err := q.db.QueryContext(ctx, getBookings, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Booking{}
	for rows.Next() {
		var i Booking
		if err := rows.Scan(
			&i.BookingId,
			&i.BookedBy,
			&i.BookOn,
			&i.BookStarts,
			&i.BookEnds,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBookingsFromToday = `-- name: GetBookingsFromToday :many
SELECT "bookingId", "bookedBy", "bookOn", "bookStarts", "bookEnds" FROM bookings
WHERE "bookStarts" > $1
`

func (q *Queries) GetBookingsFromToday(ctx context.Context, bookstarts time.Time) ([]Booking, error) {
	rows, err := q.db.QueryContext(ctx, getBookingsFromToday, bookstarts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Booking{}
	for rows.Next() {
		var i Booking
		if err := rows.Scan(
			&i.BookingId,
			&i.BookedBy,
			&i.BookOn,
			&i.BookStarts,
			&i.BookEnds,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserBookings = `-- name: GetUserBookings :many
SELECT "bookingId", "bookedBy", "bookOn", "bookStarts", "bookEnds" FROM bookings
WHERE "bookedBy" = $1
ORDER BY "bookStarts"
LIMIT $2
OFFSET $3
`

type GetUserBookingsParams struct {
	BookedBy int64 `json:"bookedBy"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) GetUserBookings(ctx context.Context, arg GetUserBookingsParams) ([]Booking, error) {
	rows, err := q.db.QueryContext(ctx, getUserBookings, arg.BookedBy, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Booking{}
	for rows.Next() {
		var i Booking
		if err := rows.Scan(
			&i.BookingId,
			&i.BookedBy,
			&i.BookOn,
			&i.BookStarts,
			&i.BookEnds,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
