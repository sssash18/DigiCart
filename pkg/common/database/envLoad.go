package database

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(filepath.Join(filepath.Join(dir, ".."), ".."), ".env")
	_ = godotenv.Load(environmentPath)
}
