package db

import (
	"context"
	"testing"

	"github.com/Golang/bank_management_system/util"
	"github.com/stretchr/testify/require"
)

func createAccountTemp(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Balance, arg.Balance)
	return account
}

func TestCreateAccount(t *testing.T) {
	createAccountTemp(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createAccountTemp(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.Balance, account2.Balance)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createAccountTemp(t)
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}
	account, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createAccountTemp(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	_, err = testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createAccountTemp(t)
	}

	accounts, err := testQueries.ListAccounts(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, accounts)
}
