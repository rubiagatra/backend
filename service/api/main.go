package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	args := os.Args

	if len(args) >= 2 {
		option := args[1]

		m, err := migrate.New(
			"file://./db/migrations",
			"postgres://postgres:postgres@localhost:5432/api?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		switch option {
		case "up":
			if err := m.Up(); err != nil {
				log.Fatal(err)
			}
			fmt.Println("Migrate up run succesfully")
		case "down":
			if err := m.Down(); err != nil {
				log.Fatal(err)
			}
			fmt.Println("Migrate down run succesfully")
		default:
			fmt.Println("The option not yet provided")
		}

	} else {
		fmt.Println("Do nothing")
	}
}
