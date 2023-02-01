package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/nga1hte/booking/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		FullName:     util.RandomFullname(),
		Email:        util.RandomEmail(),
		MobileNumber: util.RandomMobile(),
		Password:     util.RandomPassword(),
		UserType:     "normal",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.MobileNumber, user.MobileNumber)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.UserType, user.UserType)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.ID)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.Equal(t, user1.UserType, user2.UserType)

}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateUserParams{
		ID:           user1.ID,
		FullName:     user1.FullName,
		Email:        util.RandomEmail(),
		MobileNumber: user1.MobileNumber,
		Password:     user1.Password,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.Equal(t, user1.UserType, user2.UserType)
}

func TestDeleteAccount(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestGetUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := GetUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.GetUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

}
