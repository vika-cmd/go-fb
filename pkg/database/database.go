package database

import (
	"context"
	"fmt"
	"app/go-fb/config"	

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDbPool(config *config.DatabaseConfig) *pgxpool.Pool {
	dbpool, err :=pgxpool.New(context.Background(), config.Url)
	if err != nil {
		fmt.Println("error conection db!")
	}
	fmt.Println("OK! conection db!")
	return dbpool
}