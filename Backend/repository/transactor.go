package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type txCtxKey struct {
}

func txToContext(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txCtxKey{}, tx)
}

type Transactor interface {
	WithinTransaction(context.Context, func(ctx context.Context) (any, error)) (any, error)
	WithinTransactionReturnError(context.Context, func(context.Context) error) error
}

type transactor struct {
	db *sql.DB
}

func NewTransactor(db *sql.DB) *transactor {
	return &transactor{db: db}
}

func (t *transactor) WithinTransaction(ctx context.Context, txFunc func(context.Context) (any, error)) (any, error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	txCtx := txToContext(ctx, tx)
	data, err := txFunc(txCtx)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return data, nil
}

func (t *transactor) WithinTransactionReturnError(ctx context.Context, txFunc func(context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	txCtx := txToContext(ctx, tx)
	err = txFunc(txCtx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
