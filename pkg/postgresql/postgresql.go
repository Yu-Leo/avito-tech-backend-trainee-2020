package postgresql

import (
	"context"
	"fmt"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
}

func NewClient(ctx context.Context, maxAttempts int, sc config.StorageConfig) (conn *pgx.Conn, err error) {
	const timeDelta = 2 * time.Second

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)

	for i := 0; i < maxAttempts; i++ {
		time.Sleep(timeDelta)
		conn, err = pgx.Connect(ctx, dsn)
		if err == nil {
			return conn, err
		}
	}
	return nil, err
}
