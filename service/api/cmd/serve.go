/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/rubiagatra/backend/config"
	log "github.com/sirupsen/logrus"
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

		log.WithFields(log.Fields{
			"animal": "walrus",
		}).Info("A walrus appears")
		fmt.Println(viper.GetString("port"))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

}
