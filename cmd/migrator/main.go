package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var postgresHost, postgresPort, postgresUser, postgresPass, postgresDB, migrationsPath, migrationsTable string
	postgresHost = os.Getenv("POSTGRES_HOST")
	postgresPort = os.Getenv("POSTGRES_PORT")
	postgresUser = os.Getenv("POSTGRES_USER")
	postgresPass = os.Getenv("POSTGRES_PASS")
	postgresDB = os.Getenv("POSTGRES_DB")

	if postgresHost == "" {
		postgresHost = "localhost"
	}

	if postgresPort == "" {
		postgresPort = "5432"
	}

	if postgresUser == "" {
		postgresUser = "postgres"
	}

	if postgresPass == "" {
		postgresPass = "postgres"
	}

	if postgresDB == "" {
		postgresDB = "postgres"
	}
	migrationsPath = os.Getenv("MIGRATIONS_PATH")
	migrationsTable = os.Getenv("MIGRATIONS_TABLE")

	storagePath := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", postgresUser, postgresPass, postgresHost, postgresPort, postgresDB)

	if migrationsPath == "" {
		panic("MIGRATIONS_PATH Env variable is required")
	}

	if migrationsTable == "" {
		panic("MIGRATIONS_TABLE Env variable is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("%s&x-migration-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migration required")

			return
		}
		panic(err)
	}

	fmt.Println("all migration applied successfuly")
}
