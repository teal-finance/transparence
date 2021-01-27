package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	configFile string
	protocol   string
	address    []string
)

func main() {
	var rootCmd = &cobra.Command{
		Short: "Transparence command linde client",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(configFile, protocol, address)
			panic("not implemented yet")
		},
	}

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "configuration file")
	rootCmd.PersistentFlags().StringVarP(&protocol, "protocol", "p", "", "protocol to explore ")
	rootCmd.PersistentFlags().StringSliceVarP(&address, "address", "a", []string{}, "parameter for the function call")
	rootCmd.MarkPersistentFlagRequired("config")
	rootCmd.MarkPersistentFlagRequired("protocol")
	rootCmd.MarkPersistentFlagRequired("address")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
