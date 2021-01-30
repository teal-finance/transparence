package main


import (


    "fmt"
    "log"
    "math"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"  
    "github.com/ethereum/go-ethereum/accounts/abi/bind" 

     "transparence/pkg/crToken"
     "transparence/pkg/comptroller"
)

const ETH_STRING = "Ethereum"
const BINANCE_STRING = "BinanceChain"

var EXCLUDED_ADDRESSES = []string{"0xBdf447B39D152d6A234B4c02772B8ab5D1783F72","0xF5DCe57282A584D2746FaF1593d3121Fcac444dC"}

// CreamFinance Ethereum requirements
const INFURA_KEY="aefe8d98c6fd40adb30edab0e0954557"
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

func analyze(ipAddress string, comptrollerAddress string, platform string){
	client, err := ethclient.Dial(ipAddress)
    if err != nil {
        log.Fatal(err)
    }
    displayProtocolToken(comptrollerAddress,client,platform)
    markets := getAllCTokensAddress(comptrollerAddress, client)
    for i := 1; i < len(markets); i++ {
        if !find(EXCLUDED_ADDRESSES, markets[i].String()) {
            displayCompoundTokenInfos(markets[i],client,platform)
        }
    }
}

func displayCompoundTokenInfos(address common.Address, client *ethclient.Client, platform string) {
    var underlyingAddress common.Address
    
    instance, err := crToken.NewCrToken(address, client)
    if err != nil {
        log.Fatal(err)
    }

    symbol, name, decimals := getERC20InfosFromAddress(underlyingAddress, instance)

    totalBorrows, err := instance.TotalBorrows(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }

    totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }

    totalReserve, err := instance.TotalReserves(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }

    cash, err := instance.GetCash(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }

    if (symbol == "crETH" && platform == "binance") || symbol == "cETH" {
        underlyingAddress = common.HexToAddress(WETH_CONTRACT_ADDRESS)
    } else if symbol == "crBNB" {
        underlyingAddress = common.HexToAddress(WBNB_CONTRACT_ADDRESS)
    } else {
        underlyingAddress, err = instance.Underlying(&bind.CallOpts{})
        if err != nil {
            log.Fatal(err)
        }
    }

    underlyingInstance, err := crToken.NewCrToken(underlyingAddress, client)
    if err != nil {
        log.Fatal(err)
    }

    // Querying the uderlying token using erc20 abi (ex: query USDT contract for crUSDT)
    underlyingSymbol, underlyingName, underlyingDecimals := getERC20InfosFromAddress(underlyingAddress, underlyingInstance)
    
    fmt.Printf("\n ---------------------")
    fmt.Printf("\n - CreamToken Address: %s | Name: %s | Symbol: (%s)", address, name, symbol)
    fmt.Printf("\n - ERC20Token Address: %s | Name: %s | Symbol: (%s)", underlyingAddress, underlyingName, underlyingSymbol)
    fmt.Printf("\n - CreamFinance stats : TotalBorrows: %2.f %s | Reserve: %2.f ",parseDecimals(totalBorrows,decimals), symbol,parseDecimals(totalReserve,decimals))
    fmt.Printf("\n - Capitalization : TotalSupply:%2.f %s | Cash:%2.f %s ", parseDecimals(totalSupply,decimals), symbol, parseDecimals(cash,underlyingDecimals), underlyingSymbol)
    fmt.Printf("\n ---------------------\n")
}

func displayProtocolToken (comptrollerAddress string,client *ethclient.Client, platform string) {
    platformTokenAddress := getPlatformTokenAddress(comptrollerAddress, client)
    instance, err := crToken.NewCrToken(platformTokenAddress, client)
    if err != nil {
        log.Fatal(err)
    }
    platformSymbol, platformName, _ := getERC20InfosFromAddress(platformTokenAddress, instance)
    fmt.Printf("\n\n\n ==================| '%s' Protocol analysis | Symbol '%s' | Blockchain '%s' |==================\n", platformName, platformSymbol, platform)
}


func parseDecimals(totalSupply *big.Int, decimals uint8) (*big.Float){
    tmp := new(big.Float)
    tmp.SetString(totalSupply.String())
    value := new(big.Float).Quo(tmp, big.NewFloat(math.Pow10(int(decimals))))
    return value
}

func getAllCTokensAddress(comptrollerAddress string, client *ethclient.Client) ([]common.Address){
    tokenAddress := common.HexToAddress(comptrollerAddress)
    instance, err := comptroller.NewComptroller(tokenAddress, client)
    if err != nil {
        log.Fatal(err)
    }
    markets, err := instance.GetAllMarkets(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }
    return markets
}

func getPlatformTokenAddress(comptrollerAddress string, client *ethclient.Client) (common.Address){
    tokenAddress := common.HexToAddress(comptrollerAddress)
    instance, err := comptroller.NewComptroller(tokenAddress, client)
    if err != nil {
        log.Fatal(err)
    }
    platformTokenAddress, err := instance.GetCompAddress(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }
    return platformTokenAddress
}


func getERC20InfosFromAddress(address common.Address, instance *crToken.CrToken) (string, string, uint8) {
    symbol, err := instance.Symbol(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }
    decimals, err := instance.Decimals(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }

    name, err := instance.Name(&bind.CallOpts{})
    if err != nil {
        log.Fatal(err)
    }
    return symbol, name, decimals
}

func find(slice []string, val string) (bool) {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}