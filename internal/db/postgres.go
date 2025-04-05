package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"realtime-chat/internal/common"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectPostgres() {
	host := common.GetEnv("DB_HOST", "localhost")
	port := common.GetEnv("DB_PORT", "5432")
	user := common.GetEnv("DB_USER", "postgres")
	password := common.GetEnv("DB_PASSWORD", "postgres")
	dbname := common.GetEnv("DB_NAME", "chatdb")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbname,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	DB, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("❌ Erro ao conectar no PostgreSQL: %v", err)
	}

	if err = DB.Ping(ctx); err != nil {
		log.Fatalf("❌ PostgreSQL não respondeu ao ping: %v", err)
	}

	log.Println("✅ Conectado ao PostgreSQL!")
}
