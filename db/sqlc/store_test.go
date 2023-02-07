package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createBookingDayAfter(t *testing.T, user User) CreateBookingParams {
	bookStarts := time.Now().UTC().Add(time.Duration(1 * time.Minute))
	bookEnds := time.Now().UTC().Add(time.Duration(2 * time.Minute))
	arg := CreateBookingParams{
		BookedBy:   user.ID,
		BookStarts: bookStarts,
		BookEnds:   bookEnds,
	}
	return arg
}

func TestBookingTx(t *testing.T) {
	store := NewStore(testDB)
	user := createRandomUser(t)
	booking := createBookingDayAfter(t, user)

	arg := BookingTxParams{
		BookedBy:      user.ID,
		BookingStarts: booking.BookStarts,
		BookingEnds:   booking.BookEnds,
	}

	result, err := store.BookingTx(context.Background(), arg)
	require.NoError(t, err)

	userRes := result.User
	require.Equal(t, userRes.ID, user.ID)

	bookingRes := result.Booking
	require.Equal(t, bookingRes.BookedBy, user.ID)
	require.WithinDuration(t, bookingRes.BookStarts, arg.BookingStarts, 5*time.Second)
	require.WithinDuration(t, bookingRes.BookEnds, arg.BookingEnds, 5*time.Second)

	result, err = store.BookingTx(context.Background(), arg)
	require.ErrorContains(t, err, "slot already booked")

}
