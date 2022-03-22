/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migration CLI",
	Long:  "Migraion CLI",
	Run:   migrateCommand,
}

func init() {
	rootCmd.AddCommand(migrateCmd)

}

func migrateCommand(cmd *cobra.Command, args []string) {

	m, err := migrate.New(
		"file://./service/api/db/migrations",
		"postgres://postgres:postgres@localhost:5432/api?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migrate up run succesfully")

}
