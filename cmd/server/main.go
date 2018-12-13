package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/dylanlott/pangolin/pkg/database"
	"github.com/dylanlott/pangolin/pkg/server"
)

var (
	address string
	apiKey  string

	rootCmd = &cobra.Command{
		Use:   "pangolin",
		Short: "an experimental noSQL database with transactions written in Go",
		Long: `Pangolin is a NosQL database written in Go that comes with transactions,
	JSON querying, and it's named after a cute lil' guy with Go in its name.
	`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			fmt.Printf("cobra cmds: %+v\n", args)
		},
	}
)

func main() {
	rootCmd.PersistentFlags().StringVarP(&address, "port", "p", "127.0.0.1:9000", "pangolin port")
	rootCmd.PersistentFlags().StringVarP(&apiKey, "apikey", "a", "pangomango", "pangolin api key")

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln("error creating logger", err)
		os.Exit(1)
	}

	database, err := db.SetupDatabase()
	if err != nil {
		log.Fatalln("error setting up db", err)
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("error setting up listener", err)
		os.Exit(1)
	}

	s := server.NewServer(database, apiKey, logger, lis)
	err = s.GRPC.Serve(lis)
	if err != nil {
		log.Fatalln("error serving", err)
		os.Exit(1)
	}
}
