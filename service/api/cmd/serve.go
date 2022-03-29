/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"

	"github.com/rubiagatra/backend/config"
	pb "github.com/rubiagatra/backend/service/api/pb/helloworld"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve the server",
	Long:  "server the server",
	Run: func(cmd *cobra.Command, args []string) {

		sigCh := make(chan os.Signal, 1)
		errCh := make(chan error, 1)
		signal.Notify(sigCh, os.Interrupt)
		config.GetConfig()

		go func() {
			<-sigCh
			errCh <- errors.New("received an interrupt")
		}()

		// gRPC server
		go func() {
			s := grpc.NewServer()
			pb.RegisterGreeterServer(s, &server{})

			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8001))
			if err != nil {
				errCh <- err
				return
			}
			log.Printf("grpc server listening at %v", lis.Addr())
			errCh <- s.Serve(lis)
		}()

		// HTTP server
		go func() {

			e := echo.New()
			e.GET("/", func(c echo.Context) error {
				return c.String(http.StatusOK, "Hello, World!")
			})

			// Run Echo
			errCh <- e.Start(fmt.Sprintf(":%s", viper.GetString("port")))
		}()

		// Wait forever
		log.Print(<-errCh)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

}
