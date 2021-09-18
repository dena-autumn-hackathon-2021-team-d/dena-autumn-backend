package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	src = "file://./sql/"
	dsn = "sqlite3://./data.db"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Printf("the number of arguments must be 1, got %d\n", flag.NArg())
		os.Exit(1)
	}

	cmd := flag.Arg(0)

	m, err := migrate.New(src, dsn)
	if err != nil {
		fmt.Printf("error occurred: %v\n", err)
		os.Exit(1)
	}

	applyQuery(m, cmd)
}

func applyQuery(m *migrate.Migrate, cmd string) error {
	switch cmd {
	case "up":
		if err := m.Up(); err != nil {
			return fmt.Errorf("failed to migrate up: %w", err)
		}
		return nil
	case "down":
		if err := m.Down(); err != nil {
			return fmt.Errorf("failed to migrate down: %w", err)
		}
		return nil
	default:
	}

	return fmt.Errorf("invalid command: %s", cmd)
}
