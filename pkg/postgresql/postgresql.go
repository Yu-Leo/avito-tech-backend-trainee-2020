package postgresql

import (
	"context"
	"fmt"
	"github.com/Yu-Leo/avito-tech-backend-trainee-2020/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Connection interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Close()
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewConnection(ctx context.Context, maxAttempts int, sc config.StorageConfig) (conn Connection, err error) {
	const timeDelta = 2 * time.Second

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)

	for i := 0; i < maxAttempts; i++ {
		conn, err = pgxpool.New(ctx, dsn)
		if err == nil {
			return conn, err
		}
		time.Sleep(timeDelta)
	}
	return nil, err
}
