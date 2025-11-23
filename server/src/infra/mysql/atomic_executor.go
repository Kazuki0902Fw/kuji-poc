package mysql

import (
	"context"

	"kujicole/domain/repository"
	"github.com/pkg/errors"
	"github.com/stephenafamo/bob"
)

var _ repository.AtomicExecutor = new(AtomicExecutor)

type AtomicExecutor struct {
	db *bob.DB
}

func NewAtomicExecutor(db *bob.DB) *AtomicExecutor {
	return &AtomicExecutor{db}
}

type ctxMysqlTxKeyType string

const ctxMysqlTxKey ctxMysqlTxKeyType = "ctxMysqlTxKey"

func (e *AtomicExecutor) RunInTx(ctx context.Context, fc func(context.Context) error) error {
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	newCtx := context.WithValue(ctx, ctxMysqlTxKey, tx)
	err = fc(newCtx)
	if err != nil {
		tx.Rollback(ctx)
		return errors.WithStack(err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func execInTx(ctx context.Context, db *bob.DB, fc func(context.Context, bob.Transaction) error) error {
	embedTx := ctx.Value(ctxMysqlTxKey)
	if embedTx == nil {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return errors.WithStack(err)
		}

		err = fc(ctx, tx)
		if err != nil {
			tx.Rollback(ctx)
			return errors.WithStack(err)
		}

		err = tx.Commit(ctx)
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	} else {
		err := fc(ctx, embedTx.(bob.Transaction))
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
}
