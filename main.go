package main

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
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
	rootCmd.PersistentFlags().BoolVar(&debugEnabled, "debug", false, "Enable debug mode")
	rootCmd.AddCommand(startCmd)
	rootCmd.Execute()
}

func start(args []string) {
	blockForever()
}

func blockForever() {
	//select {}

	interval := viper.GetInt("CHECK_INTERVAL")
	if interval < 1 {
		interval = 5
	}

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	for _ = range ticker.C {
		runWorker()
	}
}
