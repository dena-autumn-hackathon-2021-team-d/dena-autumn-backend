package config

import "os"

func DSN() string {
	return os.Getenv("DB_PATH")
}
