package simplebank

import (
	"context"
	"github.com/haodam/Bank-Go/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {

	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

//func TestUpdateUser(t *testing.T) {
//
//	account1 := createRandomAccount(t)
//	arg := UpdateAccountParams{
//		ID:      account1.ID,
//		Balance: util.RandomMoney(),
//	}
//	account2, err := testQueries.UpdateAccount(context.Background(), arg)
//	require.NoError(t, err)
//	require.NotEmpty(t, account2)
//
//	require.Equal(t, account1.ID, account2.ID)
//	require.Equal(t, account1.Owner, account2.Owner)
//	require.Equal(t, arg.Balance, account2.Balance)
//	require.Equal(t, account1.Currency, account2.Currency)
//	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
//}
//
//func TestDeleteUser(t *testing.T) {
//	account1 := createRandomAccount(t)
//	err := testQueries.DeleteAccount(context.Background(), account1.ID)
//
//	require.NoError(t, err)
//	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
//	require.Error(t, err)
//	require.EqualError(t, err, sql.ErrNoRows.Error())
//	require.Empty(t, account2)
//}
//
//func TestListUsers(t *testing.T) {
//	for i := 0; i < 10; i++ {
//		createRandomAccount(t)
//	}
//
//	arg := ListAccountsParams{
//		Limit:  5,
//		Offset: 0,
//	}
//
//	accounts, err := testQueries.ListAccounts(context.Background(), arg)
//	require.NoError(t, err)
//	require.Len(t, accounts, 5)
//
//	for _, account := range accounts {
//		require.NotEmpty(t, account)
//	}
//}
