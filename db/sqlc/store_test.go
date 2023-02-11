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

	result1, err := store.BookingTx(context.Background(), arg)
	require.NoError(t, err)

	userRes := result1.User
	require.Equal(t, userRes.ID, user.ID)

	bookingRes := result1.Booking
	require.Equal(t, bookingRes.BookedBy, user.ID)
	require.WithinDuration(t, bookingRes.BookStarts, arg.BookingStarts, 1*time.Second)
	require.WithinDuration(t, bookingRes.BookEnds, arg.BookingEnds, 1*time.Second)

	_, err2 := store.BookingTx(context.Background(), arg)
	require.Error(t, err2)
	require.EqualError(t, err2, "slot already booked")

}
