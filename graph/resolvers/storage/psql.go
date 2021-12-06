package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
)

type Psql struct {
	DB *pg.DB
}

func (db *Psql) DBConnect() error {
	var connString = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"))
	opt, err := pg.ParseURL(connString)
	if err != nil {
		return err
	}
	db.DB = pg.Connect(opt)
	ctx := context.Background()
	if err = db.DB.Ping(ctx); err != nil {
		return err
	}
	return nil
}
