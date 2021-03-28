package config

import (
	"encoding/json"
	"fmt"
	"os"

	"transparence/pkg/ethereum/erc20"
)

func ReadConfigFile(configFile string) (erc20.PeggedTokens, error) {
	file, _ := os.Open(configFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	tokens := erc20.PeggedTokens{}
	err := decoder.Decode(&tokens)
	if err != nil {
		return erc20.PeggedTokens{}, fmt.Errorf("reading config file %v : %w", configFile, err)
	}
	return tokens, nil
}
