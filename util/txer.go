package util

import (
	"context"
	"gorm.io/gorm"
)

type KeyContext string

var (
	keyEnable KeyContext = "mysql_tx_enable"
	keyTx     KeyContext = "mysql_tx"
)

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
