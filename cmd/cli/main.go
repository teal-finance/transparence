package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"strings"

	"transparence/pkg/configuration"
	"transparence/pkg/blockchains/bitcoin"


	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	var (
		configFile string
		protocol   string
		address    []string
		mode       string
	)

	rootCmd := &cobra.Command{
		Short:         "Transparence command line client",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			var config configuration.ConfigValue

			fmt.Println("READING CONFIGURATION")
			viper.SetConfigFile(configFile)
			if err := viper.ReadInConfig(); err != nil {
				return fmt.Errorf("reading config file: %w", err)
			}
			if err := viper.Unmarshal(&config); err != nil {
				return fmt.Errorf("unable to decode configuration: %w", err)
			}

			fmt.Println("parsed flags", configFile, protocol, address, mode)
			fmt.Println("configuration", config)

			err := runCmd(config, protocol, address, mode)
			if err != nil {
				fmt.Println(color.RedString(err.Error()))
			} else {
				fmt.Println(color.GreenString("ok"))
			}
			return nil
		},
	}

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "configuration file")
	rootCmd.PersistentFlags().StringVarP(&protocol, "protocol", "p", "", "protocol to explore ")
	rootCmd.PersistentFlags().StringSliceVarP(&address, "address", "a", []string{}, "parameter for the function call")
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "cli", "server, cli")

	check(rootCmd.MarkPersistentFlagRequired("config"))
	check(rootCmd.MarkPersistentFlagRequired("protocol"))
	//check(rootCmd.MarkPersistentFlagRequired("address"))
	check(rootCmd.Execute())
}

func runCmd(config configuration.ConfigValue,protocol string, address []string, mode string) error {
	if protocol=="bitcoin"{
		rpc := config.RPC[0]
		user := strings.Split(rpc.Userpass, ":")[0]
		pass := strings.Split(rpc.Userpass, ":")[1]
		client, err := bitcoin.New(rpc.Link, user, pass)
		if err!=nil{
			return err
		}
		balance,err := client.GetBalance(address[0])
	 	if err != nil {
			return err
		}
		fmt.Println(color.GreenString("Address: "),address[0])
		fmt.Println(color.GreenString("Balance: "),balance)
		return nil
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Intn(3) == 0 {
		return fmt.Errorf("error")
	}
	return nil
}

func check(err error) {
	if err != nil {
		fmt.Printf("%v : %v\n", color.RedString("error"), err)
		os.Exit(0)
	}
}
