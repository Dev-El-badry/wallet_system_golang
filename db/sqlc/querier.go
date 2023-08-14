// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entries, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Sessions, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfers, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Accounts, error)
	GetAccountForUpdated(ctx context.Context, id int64) (Accounts, error)
	GetAccounts(ctx context.Context, arg GetAccountsParams) ([]Accounts, error)
	GetEntry(ctx context.Context, id int64) (Entries, error)
	GetSession(ctx context.Context, id uuid.UUID) (Sessions, error)
	GetTransfer(ctx context.Context, id int64) (Transfers, error)
	GetUser(ctx context.Context, username string) (Users, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entries, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfers, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Accounts, error)
	UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) (Accounts, error)
}

var _ Querier = (*Queries)(nil)
