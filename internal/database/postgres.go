package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	var ctx context.Context = context.Background()
	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = pool.Ping(ctx)

	if err != nil {
		log.Println(err)
		pool.Close()
		return nil, err
	}

	log.Println("successfully connected ot database")

	return pool, nil

}
