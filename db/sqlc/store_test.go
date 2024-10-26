package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createAccountTemp(t)
	account2 := createAccountTemp(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		result := <-results

		require.NoError(t, err)
		require.NotEmpty(t, result)

		transfer, errr := store.Queries.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, errr)
		require.NotEmpty(t, transfer.ID)

		fromEntry, errr := store.Queries.GetEntry(context.Background(), result.FromEntry.ID)
		require.NoError(t, errr)
		require.NotEmpty(t, fromEntry.ID)

		toEntry, errr := store.Queries.GetEntry(context.Background(), result.ToEntry.ID)
		require.NoError(t, errr)
		require.NotEmpty(t, toEntry.ID)

		require.NotEmpty(t, result.FromAccount)
		require.NotEmpty(t, result.FromAccount.ID)

		require.NotEmpty(t, result.ToAccount)
		require.NotEmpty(t, result.ToAccount.ID)

		diff1 := account1.Balance - result.FromAccount.Balance
		diff2 := result.ToAccount.Balance - account2.Balance

		require.Equal(t, diff1, diff2)
	}

	updatedAccount1, err := store.Queries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.Queries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance-(int64(n)*amount), updatedAccount1.Balance)
	require.Equal(t, account2.Balance+(int64(n)*amount), updatedAccount2.Balance)

}
