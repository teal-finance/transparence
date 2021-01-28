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

const INFURA_KEY=""
const IP_API_INFURA = "https://mainnet.infura.io/v3/" + INFURA_KEY
const CONTRACT_TO_EXCLUDE = "0xBdf447B39D152d6A234B4c02772B8ab5D1783F72"
const COMPTROLLER_CONTRACT_ADDRESS = "0x3d5BC3c8d13dcB8bF317092d84783c2697AE9258"
// WETH CONTRACT ADDRESS AND CONTRACT_TO_EXCLUDE are specific because smart contract does not contain same function than others crTokens
const WETH_CONTRACT_ADDRESS = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"



func main() {
    var underlyingAddress common.Address


	client, err := ethclient.Dial(IP_API_INFURA)
    if err != nil {
        log.Fatal(err)
    }

    markets := getAllCrTokensAddress(COMPTROLLER_CONTRACT_ADDRESS, client)

    for i := 0; i < len(markets); i++ {
        if markets[i] != common.HexToAddress(CONTRACT_TO_EXCLUDE) {
            instance, err := crToken.NewCrToken(markets[i], client)
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

            if symbol == "crETH" {
                underlyingAddress = common.HexToAddress(WETH_CONTRACT_ADDRESS)
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

            // Querying the uderlying token informations (ex: crUSDT underly USDT)
            underlyingSymbol, underlyingName, underlyingDecimals := getERC20InfosFromAddress(underlyingAddress, underlyingInstance)
            
            fmt.Printf("\n ---------------------")
            fmt.Printf("\n - CreamToken Address: %s | Name: %s | Symbol: (%s)", markets[i], name, symbol)
            fmt.Printf("\n - ERC20Token Address: %s | Name: %s | Symbol: (%s)", underlyingAddress, underlyingName, underlyingSymbol)
            fmt.Printf("\n - CreamFinance stats : TotalBorrows: %2.f %s | Reserve: %2.f ",parseDecimals(totalBorrows,decimals), symbol,parseDecimals(totalReserve,decimals))
            fmt.Printf("\n - Capitalization : TotalSupply:%2.f %s | Cash:%2.f %s ", parseDecimals(totalSupply,decimals), symbol, parseDecimals(cash,underlyingDecimals), underlyingSymbol)
            fmt.Printf("\n ---------------------\n")
            }
        }

    
}


func parseDecimals(totalSupply *big.Int, decimals uint8) (*big.Float){
    tmp := new(big.Float)
    tmp.SetString(totalSupply.String())
    value := new(big.Float).Quo(tmp, big.NewFloat(math.Pow10(int(decimals))))
    return value
}

func getAllCrTokensAddress(comptrollerAddress string, client *ethclient.Client) ([]common.Address){
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
