package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store interface {
	Querier
	BookingTx(ctx context.Context, arg BookingTxParams) (BookingTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// executes a funcion within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type BookingTxParams struct {
	BookedBy      int64     `json:"bookedBy"`
	BookingStarts time.Time `json:"bookStarts"`
	BookingEnds   time.Time `json:"bookEnds"`
}

type BookingTxResult struct {
	Booking Booking `json:"booking"`
	User    User    `json:"user"`
}

func (store *SQLStore) BookingTx(ctx context.Context, arg BookingTxParams) (BookingTxResult, error) {
	var result BookingTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Booking, err = q.CreateBooking(ctx, CreateBookingParams{
			BookedBy:   arg.BookedBy,
			BookStarts: arg.BookingStarts,
			BookEnds:   arg.BookingEnds,
		})

		if err != nil {
			return err
		}

		result.User, err = q.GetUser(ctx, arg.BookedBy)
		if err != nil {
			return err
		}

		return nil
	})
	return result, err

}
