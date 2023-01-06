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

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
}

func NewClient(ctx context.Context, maxAttempts int, sc config.StorageConfig) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)

	for i := 0; i < maxAttempts; i++ {
		pool, err = tryToConnectPostgres(ctx, dsn)
		if err == nil {
			//fmt.Println("No error")
			return pool, nil
		}
		time.Sleep(1 * time.Second)
	}
	return nil, err
}

func tryToConnectPostgres(ctx context.Context, dsn string) (pool *pgxpool.Pool, err error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		//fmt.Println("Error 2", err)
		return nil, err
	}
	//fmt.Println("No error 2", pool)
	return pool, nil
}
