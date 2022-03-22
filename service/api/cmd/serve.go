/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rubiagatra/backend/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve the server",
	Long:  "server the server",
	Run: func(cmd *cobra.Command, args []string) {
		config.GetConfig()
		e := echo.New()
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", viper.GetString("port"))))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

}
