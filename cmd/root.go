package main

import (
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
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

// Execute runs the rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
