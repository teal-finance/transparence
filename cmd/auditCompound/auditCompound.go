package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"transparence/pkg/ethereum/erc20/compound"
	"transparence/pkg/math"
)

const ETH_STRING = "Ethereum"
const BINANCE_STRING = "BinanceChain"

var EXCLUDED_ADDRESSES = []string{"0xBdf447B39D152d6A234B4c02772B8ab5D1783F72", "0xF5DCe57282A584D2746FaF1593d3121Fcac444dC"}

// CreamFinance Ethereum requirements
const INFURA_KEY = ""
const IP_API_INFURA = "https://mainnet.infura.io/v3/" + INFURA_KEY
const COMPTROLLER_CONTRACT_CREAM_ETHEREUM = "0x3d5BC3c8d13dcB8bF317092d84783c2697AE9258"
const WETH_CONTRACT_ADDRESS = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

// CreamFinance BinanceSmartChain requirements
const IP_API_BINANCECHAIN = "wss://bsc-ws-node.nariox.org:443"
const COMPTROLLER_CONTRACT_CREAM_BINANCE = "0x589de0f0ccf905477646599bb3e5c622c84cc0ba"
const WBNB_CONTRACT_ADDRESS = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"

// Compound Ethereum requirements
const COMPTROLLER_CONTRACT_COMPOUND_ETHEREUM = "0x3d9819210A31b4961b30EF54bE2aeD79B9c9Cd3B"

func main() {
	analyze(IP_API_INFURA, COMPTROLLER_CONTRACT_CREAM_ETHEREUM, ETH_STRING)

	analyze(IP_API_BINANCECHAIN, COMPTROLLER_CONTRACT_CREAM_BINANCE, BINANCE_STRING)

	analyze(IP_API_INFURA, COMPTROLLER_CONTRACT_COMPOUND_ETHEREUM, ETH_STRING)
}

func analyze(ipAddress string, comptrollerAddress string, platform string) {
	client := compound.NewClientConnection(ipAddress)
	tokenAddress := common.HexToAddress(comptrollerAddress)
	instanceComptroller := compound.NewComptroller(tokenAddress, client)
	platformTokenAddress := compound.GetCompAddress(instanceComptroller)
	instance := compound.NewCrToken(platformTokenAddress, client)
	displayProtocolToken(compound.GetName(instance), compound.GetSymbol(instance), platform)

	markets := compound.GetAllMarkets(instanceComptroller)
	for i := 1; i < len(markets); i++ {
		_, isExcluded := findInSlice(EXCLUDED_ADDRESSES, markets[i].String())
		if !isExcluded {
			displayCompoundTokenInfos(markets[i], client, platform)
		}
	}
}

func displayCompoundTokenInfos(address common.Address, client *ethclient.Client, platform string) {
	var underlyingAddress common.Address

	instance := compound.NewCrToken(address, client)
	name := compound.GetName(instance)
	symbol := compound.GetSymbol(instance)
	decimals := compound.GetDecimals(instance)
	totalSupply := compound.GetTotalSupply(instance)
	totalBorrows := compound.GetTotalBorrows(instance)
	totalReserve := compound.GetTotalReserves(instance)
	cash := compound.GetCash(instance)

	if (symbol == "crETH" && platform == "binance") || symbol == "cETH" {
		underlyingAddress = common.HexToAddress(WETH_CONTRACT_ADDRESS)
	} else if symbol == "crBNB" {
		underlyingAddress = common.HexToAddress(WBNB_CONTRACT_ADDRESS)
	} else {
		underlyingAddress = compound.GetUnderlying(instance)
	}

	underlyingInstance := compound.NewCrToken(underlyingAddress, client)
	underlyingName := compound.GetName(underlyingInstance)
	underlyingSymbol := compound.GetSymbol(underlyingInstance)
	underlyingDecimals := compound.GetDecimals(underlyingInstance)

	fmt.Printf("\n ---------------------")
	fmt.Printf("\n - cToken Address: %s | Name: %s | Symbol: (%s)", address, name, symbol)
	fmt.Printf("\n - ERC20Token Address: %s | Name: %s | Symbol: (%s)", underlyingAddress, underlyingName, underlyingSymbol)
	fmt.Printf("\n - Defi stats : TotalBorrows: %2.f %s | Reserve: %2.f ", math.ParseDecimals(totalBorrows, decimals), symbol, math.ParseDecimals(totalReserve, decimals))
	fmt.Printf("\n - Capitalization : TotalSupply:%2.f %s | Cash:%2.f %s ", math.ParseDecimals(totalSupply, decimals), symbol, math.ParseDecimals(cash, underlyingDecimals), underlyingSymbol)
	fmt.Printf("\n ---------------------\n")
}

func displayProtocolToken(platformName string, platformSymbol string, platform string) {
	fmt.Printf("\n\n\n ==================| '%s' Protocol analysis | Symbol '%s' | Blockchain '%s' |==================\n", platformName, platformSymbol, platform)
}

func findInSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
