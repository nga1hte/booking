package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBookingTx(t *testing.T) {
	store := NewStore(testDB)

	user := createRandomUser(t)
	booking := createRandomBooking(t, user)
	n := 5
	errs := make(chan error)
	results := make(chan BookingTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.BookingTx(context.Background(), BookingTxParams{
				BookedBy:      booking.BookedBy,
				BookingStarts: booking.BookStarts,
				BookingEnds:   booking.BookEnds,
			})
			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)

		userT := result.User
		require.NotEmpty(t, userT)
		require.Equal(t, userT.ID, user.ID)
		require.Equal(t, userT.FullName, user.FullName)
		require.Equal(t, userT.Email, user.Email)
		require.Equal(t, userT.MobileNumber, user.MobileNumber)
		require.Equal(t, userT.Password, user.Password)
		require.Equal(t, userT.UserType, user.UserType)
		require.NotZero(t, userT.CreatedAt)
		require.NotZero(t, userT.ID)

		bookingT := result.Booking
		require.NoError(t, err)
		require.NotEmpty(t, bookingT)
		require.Equal(t, bookingT.BookedBy, booking.BookedBy)
		require.WithinDuration(t, bookingT.BookStarts, booking.BookStarts, time.Second)
		require.WithinDuration(t, bookingT.BookEnds, booking.BookEnds, time.Second)
		require.NotZero(t, bookingT.BookOn)
		require.NotZero(t, bookingT.BookingId)

		_, err = store.GetBooking(context.Background(), booking.BookingId)
		require.NoError(t, err)

	}

}
