package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	loadConfig()
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start Ingress-dns in non-daemon mode",
		Run: func(cmd *cobra.Command, args []string) {
			start(args)
		},
	}

	rootCmd := &cobra.Command{
		Use: "ingress-dns",
	}
	rootCmd.AddCommand(startCmd)
	rootCmd.Execute()
}

func start(args []string) {
	fmt.Println(GetServices())
	blockForever()
}

func blockForever() {
	select {}
}
