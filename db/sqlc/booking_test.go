package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/nga1hte/booking/util"
	"github.com/stretchr/testify/require"
)

func createRandomBooking(t *testing.T, user User) Booking {
	bookStarts, bookEnds := util.RandomBookingDates()
	arg := CreateBookingParams{
		BookedBy:   user.ID,
		BookStarts: bookStarts,
		BookEnds:   bookEnds,
	}
	booking, err := testQueries.CreateBooking(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, booking)
	require.Equal(t, arg.BookedBy, booking.BookedBy)
	require.WithinDuration(t, arg.BookStarts, booking.BookStarts, time.Second)
	require.WithinDuration(t, arg.BookEnds, booking.BookEnds, time.Second)
	require.NotZero(t, booking.BookOn)
	require.NotZero(t, booking.BookingId)
	return booking

}

func TestCreateBooking(t *testing.T) {
	u := createRandomUser(t)
	createRandomBooking(t, u)
}

func TestGetBooking(t *testing.T) {
	u := createRandomUser(t)
	booking1 := createRandomBooking(t, u)
	booking2, err := testQueries.GetBooking(context.Background(), booking1.BookingId)
	require.NoError(t, err)
	require.NotEmpty(t, booking2)

	require.Equal(t, booking1.BookingId, booking2.BookingId)
	require.Equal(t, booking1.BookedBy, booking2.BookedBy)
	require.Equal(t, booking1.BookOn, booking2.BookOn)
	require.Equal(t, booking1.BookStarts, booking2.BookStarts)
	require.Equal(t, booking1.BookEnds, booking2.BookEnds)
}

func TestDeleteBooking(t *testing.T) {
	u := createRandomUser(t)
	booking1 := createRandomBooking(t, u)
	err := testQueries.DeleteBooking(context.Background(), booking1.BookingId)
	require.NoError(t, err)
	booking2, err := testQueries.GetBooking(context.Background(), booking1.BookingId)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, booking2)
}

func TestGetBookings(t *testing.T) {
	u := createRandomUser(t)
	for i := 0; i < 10; i++ {
		createRandomBooking(t, u)
	}
	arg := GetBookingsParams{
		Limit:  5,
		Offset: 5,
	}

	bookings, err := testQueries.GetBookings(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, bookings, 5)

	for _, booking := range bookings {
		require.NotEmpty(t, booking)
	}
}

func TestGetUserBookings(t *testing.T) {
	u := createRandomUser(t)
	for i := 0; i < 5; i++ {
		createRandomBooking(t, u)
	}
	bookings, err := testQueries.GetUserBookings(context.Background(), u.ID)
	require.NoError(t, err)
	require.Len(t, bookings, 5)
	for _, booking := range bookings {
		require.NotEmpty(t, booking)
		require.Equal(t, u.ID, booking.BookedBy)
	}

}

func TestBookingsFromToday(t *testing.T) {
	u := createRandomUser(t)
	for i := 0; i < 3; i++ {
		createRandomBooking(t, u)
	}

	bookStarts := time.Now().Local().Add(time.Duration(3) * time.Hour)

	bookings, err := testQueries.GetBookingsFromToday(context.Background(), bookStarts)
	require.NoError(t, err)

	for _, booking := range bookings {
		require.GreaterOrEqual(t, booking.BookStarts, bookStarts)
	}

}
