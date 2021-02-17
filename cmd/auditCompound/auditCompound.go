package auditCompound

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	transparenceutils "transparence/pkg/TransparenceUtils"
	"transparence/pkg/cTokenAdapter"
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

func RunAudit(blockchain string, platform string) {
	if blockchain == "ethereum" && platform == "cream" {
		analyze(IP_API_INFURA, COMPTROLLER_CONTRACT_CREAM_ETHEREUM, ETH_STRING)
	}
	if blockchain == "binance" && platform == "cream"{
		analyze(IP_API_BINANCECHAIN, COMPTROLLER_CONTRACT_CREAM_BINANCE, BINANCE_STRING)
	}
	if platform == "compound" {
		analyze(IP_API_INFURA, COMPTROLLER_CONTRACT_COMPOUND_ETHEREUM, ETH_STRING)
	}
}

func analyze(ipAddress string, comptrollerAddress string, platform string) {
	client := cTokenAdapter.NewClientConnection(ipAddress)
	tokenAddress := common.HexToAddress(comptrollerAddress)
	instanceComptroller := cTokenAdapter.NewComptroller(tokenAddress, client)
	platformTokenAddress := cTokenAdapter.GetCompAddress(instanceComptroller)
	instance := cTokenAdapter.NewCrToken(platformTokenAddress, client)
	displayProtocolToken(cTokenAdapter.GetName(instance), cTokenAdapter.GetSymbol(instance), platform)

	markets := cTokenAdapter.GetAllMarkets(instanceComptroller)
	for i := 1; i < len(markets); i++ {
		_, isExcluded := transparenceutils.Find(EXCLUDED_ADDRESSES, markets[i].String())
		if !isExcluded {
			displayCompoundTokenInfos(markets[i], client, platform)
		}
	}
}

func displayCompoundTokenInfos(address common.Address, client *ethclient.Client, platform string) {
	var underlyingAddress common.Address

	instance := cTokenAdapter.NewCrToken(address, client)
	name := cTokenAdapter.GetName(instance)
	symbol := cTokenAdapter.GetSymbol(instance)
	decimals := cTokenAdapter.GetDecimals(instance)
	totalSupply := cTokenAdapter.GetTotalSupply(instance)
	totalBorrows := cTokenAdapter.GetTotalBorrows(instance)
	totalReserve := cTokenAdapter.GetTotalReserves(instance)
	cash := cTokenAdapter.GetCash(instance)

	if (symbol == "crETH" && platform == "binance") || symbol == "cETH" {
		underlyingAddress = common.HexToAddress(WETH_CONTRACT_ADDRESS)
	} else if symbol == "crBNB" {
		underlyingAddress = common.HexToAddress(WBNB_CONTRACT_ADDRESS)
	} else {
		underlyingAddress = cTokenAdapter.GetUnderlying(instance)
	}

	underlyingInstance := cTokenAdapter.NewCrToken(underlyingAddress, client)
	underlyingName := cTokenAdapter.GetName(underlyingInstance)
	underlyingSymbol := cTokenAdapter.GetSymbol(underlyingInstance)
	underlyingDecimals := cTokenAdapter.GetDecimals(underlyingInstance)

	fmt.Printf("\n ---------------------")
	fmt.Printf("\n - cToken Address: %s | Name: %s | Symbol: (%s)", address, name, symbol)
	fmt.Printf("\n - ERC20Token Address: %s | Name: %s | Symbol: (%s)", underlyingAddress, underlyingName, underlyingSymbol)
	fmt.Printf("\n - Defi stats : TotalBorrows: %2.f %s | Reserve: %2.f ", transparenceutils.ParseDecimals(totalBorrows, decimals), symbol, transparenceutils.ParseDecimals(totalReserve, decimals))
	fmt.Printf("\n - Capitalization : TotalSupply:%2.f %s | Cash:%2.f %s ", transparenceutils.ParseDecimals(totalSupply, decimals), symbol, transparenceutils.ParseDecimals(cash, underlyingDecimals), underlyingSymbol)
	fmt.Printf("\n ---------------------\n")
}

func displayProtocolToken(platformName string, platformSymbol string, platform string) {
	fmt.Printf("\n\n\n ==================| '%s' Protocol analysis | Symbol '%s' | Blockchain '%s' |==================\n", platformName, platformSymbol, platform)
}
