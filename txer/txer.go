package txer

import (
	"context"
	"green.env.com/auth/config"
	"green.env.com/auth/util"
	"strings"

	"gorm.io/gorm"
)

type KeyContext string

var (
	keyEnable    KeyContext = "mysql_tx_enable"
	keyTx        KeyContext = "mysql_tx"
	keyGetClient KeyContext = "db_get_client"
)

func TxBegin(ctx context.Context, getClient func(ctx context.Context) *gorm.DB) context.Context {
	db := getClient(ctx)
	tx := db.Begin()
	ctx = SetTx(ctx, tx)

	ctx = context.WithValue(ctx, keyEnable, true)

	return ctx
}

func TxEnd(ctx context.Context, txFunc func(context.Context) error) (context.Context, error) {
	var (
		err   error
		retry = config.GetConfig().MySQL.CountRetryTx
	)

	for {
		tx := GetTx(ctx)

		func(ctx context.Context) {
			defer func() {
				p := recover()

				switch {
				case p != nil:
					tx.Rollback()
					panic(p) // re-throw panic after Rollback
				case err != nil:
					tx.Rollback() // err is non-nil; don't change it
				default:
					err = tx.Commit().Error // if Commit returns error update err with commit err
					if err != nil {
						tx.Logger.Error(ctx, "fail commit transaction", err)
						err = util.ErrCommitTransaction(err)
					}
				}
			}()

			err = txFunc(ctx)
		}(ctx)

		if err == nil || !IsDeadlockError(err) {
			break
		}

		retry--

		if retry <= 0 {
			break
		}

		// update context to retry
		getClient := GetGetClientFunc(ctx)
		ctx = context.WithValue(ctx, keyEnable, false)
		ctx = TxBegin(ctx, getClient)
	}

	ctx = context.WithValue(ctx, keyEnable, false)

	return ctx, err
}

func IsEnableTx(ctx context.Context) bool {
	txEnable, ok := ctx.Value(keyEnable).(bool)

	return ok && txEnable
}

func GetTx(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(keyTx).(*gorm.DB)
	if !ok {
		return nil
	}

	return tx
}

func SetTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, keyTx, tx)
}

func GetGetClientFunc(ctx context.Context) func(ctx context.Context) *gorm.DB {
	getClient, ok := ctx.Value(keyGetClient).(func(ctx context.Context) *gorm.DB)
	if !ok {
		return nil
	}

	return getClient
}

func IsDeadlockError(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "deadlock")
}
