package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
)

var DB *pg.DB

func DBConnect() (err error) {
	var connString = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"))
	opt, err := pg.ParseURL(connString)
	if err != nil {
		return err
	}
	DB = pg.Connect(opt)
	ctx := context.Background()

	if err := DB.Ping(ctx); err != nil {
		return err
	}
	return
}
