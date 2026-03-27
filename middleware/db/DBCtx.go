package db

import (
	"context"

	"gorm.io/gorm"
)

type contextTxKey struct{}

// Transaction 包装器
func Transaction(ctx context.Context, db *gorm.DB, fn func(ctx context.Context) error) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 将事务对象 tx 注入新的 context 中
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// GetDB 从 context 中获取事务或返回默认 DB
func GetDB(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB); ok {
		return tx
	}
	return defaultDB.WithContext(ctx)
}
