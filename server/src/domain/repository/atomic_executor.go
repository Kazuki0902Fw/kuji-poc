package repository

import (
	"context"

	"github.com/pkg/errors"
)

var globalAE AtomicExecutor

// AtomicExecutor は repository メソッドの整合性を担保した実行手段を提供する interface を表す。
type AtomicExecutor interface {
	RunInTx(context.Context, func(context.Context) error) error
}

// SetAtomicExecutor は global な AtomicExecutor をセットする
func SetAtomicExecutor(ae AtomicExecutor) {
	globalAE = ae
}

// RunInTx は global な AtomicExecutor を利用して整合性を担保した repository メソッド実行トランザクションを提供する。
func RunInTx(ctx context.Context, fc func(context.Context) error) error {
	if globalAE == nil {
		return errors.WithStack(errors.New("global consistency executor is not set"))
	}
	return globalAE.RunInTx(ctx, fc)
}
