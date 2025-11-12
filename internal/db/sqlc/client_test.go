package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/YosaZiege/banksystem/internal/util"
)

func createRandomClient(t *testing.T) Client {
	hashedPassword := util.RandomPasswordHash()

	arg := CreateClientParams{
		Username:     util.RandomUsername(),
		Balance:      util.RandomInit(100, 10000000),
		Currency:     util.RandomString(20),
		Email:        util.RandomEmail(),
		PasswordHash: hashedPassword,
		Provider:     util.RandomProvider(),
	}

	user, err := testQueries.CreateClient(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Balance, user.Balance)
	require.Equal(t, arg.Currency, user.Currency)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)
	require.Equal(t, arg.Provider, user.Provider)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomClient(t)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomClient(t)

	err := testQueries.DeleteClient(context.Background(), user1.Username)
	require.NoError(t, err)

	user2, err := testQueries.GetClient(context.Background(), user1.Username)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, user2)
}

func TestGetClient(t *testing.T) {
	user1 := createRandomClient(t)

	user2, err := testQueries.GetClient(context.Background(), user1.Username)

	require.NotEmpty(t, user2)
	require.NoError(t, err)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Balance, user2.Balance)
	require.Equal(t, user1.Currency, user2.Currency)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.Equal(t, user1.Provider, user2.Provider)
}

func TestListClients(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomClient(t)
	}

	arg := ListClientsParams{
		Limit:  5,
		Offset: 5,
	}
	clients, err := testQueries.ListClients(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, clients, 5)

	for _, client := range clients {
		require.NotEmpty(t, client)
	}
}

func TestUpdateBalance(t *testing.T) {
	user1 := createRandomClient(t)

	arg := UpdateClientBalanceParams{
		Username: user1.Username,
		Balance:  util.RandomInit(100, 10000000),
	}

	user2, err := testQueries.UpdateClientBalance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, arg.Balance, user2.Balance)
	require.Equal(t, user1.Currency, user2.Currency)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.Equal(t, user1.Provider, user2.Provider)
}
