package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomPayment(t *testing.T, b Booking) Payment {
	payment, err := testQueries.CreatePayment(context.Background(), b.BookingId)
	require.NoError(t, err)
	require.NotEmpty(t, payment)
	require.Equal(t, b.BookingId, payment.BookingId)
	require.NotZero(t, payment.PaymentDate)
	require.NotZero(t, payment.PaymentId)
	return payment
}

func TestCreatePayment(t *testing.T) {
	u := createRandomUser(t)
	b := createRandomBooking(t, u)
	createRandomPayment(t, b)
}

func TestGetBookingPayment(t *testing.T) {
	u := createRandomUser(t)
	b := createRandomBooking(t, u)
	payment1 := createRandomPayment(t, b)
	payment2, err := testQueries.GetBookingPayment(context.Background(), payment1.BookingId)
	require.NoError(t, err)
	require.NotEmpty(t, payment2)

	require.Equal(t, payment1.PaymentId, payment2.PaymentId)
	require.Equal(t, payment1.BookingId, payment2.BookingId)
	require.WithinDuration(t, payment1.PaymentDate, payment2.PaymentDate, time.Second)

}

func TestGetPayments(t *testing.T) {
	u := createRandomUser(t)
	for i := 0; i < 5; i++ {
		b := createRandomBooking(t, u)
		createRandomPayment(t, b)

	}
	arg := GetPaymentsParams{
		Limit:  5,
		Offset: 0,
	}

	payments, err := testQueries.GetPayments(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, payments, 5)

	for _, payment := range payments {
		require.NotEmpty(t, payment)
	}
}
