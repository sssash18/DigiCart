package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	dir, err := os.Getwd()
	fmt.Println(dir)
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(filepath.Join(filepath.Join(dir, ".."), ".."), ".env")
	_ = godotenv.Load(environmentPath)
}
