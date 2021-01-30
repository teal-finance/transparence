package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	var (
		configFile string
		protocol   string
		address    []string
	)

	rootCmd := &cobra.Command{
		Short:         "Transparence command linde client",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			var config configValue

			fmt.Println("READING CONFIGURATION")
			viper.SetConfigFile(configFile)
			if err := viper.ReadInConfig(); err != nil {
				return fmt.Errorf("reading config file: %w", err)
			}
			if err := viper.Unmarshal(&config); err != nil {
				return fmt.Errorf("unable to decode configuration: %w", err)
			}

			fmt.Println("parsed flags", configFile, protocol, address)
			fmt.Println("configuration", protocol, address)
			panic("not implemented")
		},
	}

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "configuration file")
	rootCmd.PersistentFlags().StringVarP(&protocol, "protocol", "p", "", "protocol to explore ")
	rootCmd.PersistentFlags().StringSliceVarP(&address, "address", "a", []string{}, "parameter for the function call")
	check(rootCmd.MarkPersistentFlagRequired("config"))
	check(rootCmd.MarkPersistentFlagRequired("protocol"))
	check(rootCmd.MarkPersistentFlagRequired("address"))

	check(rootCmd.Execute())
}

func check(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(0)
	}
}

type configValue struct {
	RPC         []RPC         `mapstructure:"rpc"`
	Useraddress []UserAddress `mapstructure:"useraddress"`
	Tokens      []Token       `mapstructure:"tokens"`
}

type RPC struct {
	Blockchain string `mapstructure:"blockchain"`
	Userpass   string `mapstructure:"userpass"`
}

type UserAddress struct {
	Blockchain string `mapstructure:"blockchain"`
	Address    string `mapstructure:"address"`
}

type Token struct {
	BlockchainPlatform string   `mapstructure:"blockchain_platform"`
	Symbol             string   `mapstructure:"symbol"`
	ContractAddress    string   `mapstructure:"contract_address"`
	BtcAddresses       []string `mapstructure:"btc_addresses"`
}
