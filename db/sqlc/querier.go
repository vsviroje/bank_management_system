// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteEntry(ctx context.Context, id int64) error
	DeleteTransfer(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	ListAccounts(ctx context.Context) ([]Account, error)
	ListAccountsWithPagination(ctx context.Context, arg ListAccountsWithPaginationParams) ([]Account, error)
	ListEntries(ctx context.Context) ([]Entry, error)
	ListTransfersByFromAccId(ctx context.Context) ([]Transfer, error)
	ListTransfersByToAccId(ctx context.Context) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
