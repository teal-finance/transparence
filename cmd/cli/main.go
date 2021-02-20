package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
	"transparence/cmd/auditBTC"
	"transparence/cmd/auditBsc"
	"transparence/cmd/auditCompound"
)

func main() {
	var (
		protocol   string
		symbol 	string
		blockchain string
		tellor bool
	)

	rootCmd := &cobra.Command{
		Short:         "Transparence command line interface",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			executePlatformAudit(blockchain,protocol,symbol,tellor)
			return nil
		},
	}

	//rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "configuration file")
	rootCmd.PersistentFlags().StringVarP(&blockchain, "blockchain", "b", "", "blockchain to explore ")
	rootCmd.PersistentFlags().StringVarP(&protocol, "protocol", "p", "", "protocol to explore ")
	rootCmd.PersistentFlags().StringVarP(&symbol, "symbol", "s", "", "symbol of the cryptocurrency to audit")
	rootCmd.PersistentFlags().BoolVarP(&tellor, "tellor", "t", false, "tellor feature enabled")
	//rootCmd.PersistentFlags().StringSliceVarP(&address, "address", "a", []string{}, "parameter for the function call")
	//check(rootCmd.MarkPersistentFlagRequired("config"))
	//check(rootCmd.MarkPersistentFlagRequired("protocol"))
	//check(rootCmd.MarkPersistentFlagRequired("addre-(ss"))

	check(rootCmd.Execute())
}

func check(err error) {
	if err != nil {
		fmt.Printf("%v : %v", color.RedString("error"), err)
		os.Exit(0)
	}
}

func executePlatformAudit(blockchain string, protocol string, symbol string, tellor bool){
	if blockchain == "binance" && protocol == "pegged-tokens" {
		auditBsc.RunAudit(symbol,tellor)
	}
	if (blockchain == "ethereum" && protocol == "compound" || protocol == "cream") {
		auditCompound.RunAudit(blockchain,protocol)
	}
	if blockchain == "bitcoin" && protocol == "pegged-tokens" {
		auditBTC.RunAudit()
	}
}

type configValue struct {
	Api_keys      []ApiKey       `mapstructure:"api_keys"`
}

type ApiKey struct {
	Blockchain string `mapstructure:"blockchain"`
	Api_key      string       `mapstructure:"api_key"`
}
