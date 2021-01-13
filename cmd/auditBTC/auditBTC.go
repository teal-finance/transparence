package main

import (
    "encoding/json"
    "os"
    "flag"

    "io/ioutil"
    "net/http"

    "fmt"
    "log"
    "math"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    "transparence/pkg/token"   // for ERC20 request (package generated using smart contract ABI)
)

const CONFIG_FILE = "config.json"
const INFURA_KEY=""
const IP_API_INFURA = "https://mainnet.infura.io/v3/" + INFURA_KEY
const IP_API_COINGECKO = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=bitcoin"
const IP_API_BINANCECHAIN = "wss://bsc-ws-node.nariox.org:443"

const ETH_STRING = "Ethereum"
const BINANCE_STRING = "BinanceChain"


// Types used to represent the tokens of the JSON config file
type Token struct {
    BLOCKCHAIN_PLATFORM string `json:"BLOCKCHAIN_PLATFORM"`
    SYMBOL string `json:"SYMBOL"`
    CONTRACT_ADDRESS   string `json:"CONTRACT_ADDRESS"`
    BTC_ADDRESSES    []string `json:"BTC_ADDRESSES"`
}

type Tokens struct {
    Tokens []Token `json:"tokens"`
}

// Types used to represent Coingecko reponses (a lot of more infos are also available but not stored)
type CoinsMarketItem struct {
    Symbol                              string         `json:"symbol"`
    Name                                string         `json:"name"`
    TotalSupply                         float64        `json:"total_supply"`
    CirculatingSupply                   float64        `json:"circulating_supply"`
    CurrentPrice                        float64        `json:"current_price"`
    MarketCap                           float64        `json:"market_cap"`
    MarketCapRank                       int16          `json:"market_cap_rank"`
}

type CoinsMarket []CoinsMarketItem

func main() {
  rpcURL := flag.String("r", "127.0.0.1:8545", "rpc url ")

  flag.Parse()
  fmt.Printf("\n\n ################## Analysis of Ethereum ERC20 BTC ################## \n\n")
  totalOnEth := analyze(*rpcURL, ETH_STRING)

  fmt.Printf("\n\n ################## Analysis of Binance chain BEP20 BTC ################## \n\n")
  totalOnBinance := analyze(IP_API_BINANCECHAIN, BINANCE_STRING)

  fmt.Printf("\n\n ################## Analysis of Ethereum ERC20 & Binance chain BEP20 BTC ################## \n\n")
  totalOverall := totalOnEth.Add(totalOnEth,totalOnBinance)
  printBitcoinComparison(totalOverall,"Ethereum & Binance chain")

}

func analyze(ipAddress string, blockchainPlatform string) (*big.Float){
    total := big.NewFloat(0)

    tokens := readConfigFile(CONFIG_FILE)
    client, err := ethclient.Dial(ipAddress)
    if err != nil {
        log.Fatal(err)
    }
    for i := 0; i < len(tokens.Tokens); i++ {
        if (tokens.Tokens[i].BLOCKCHAIN_PLATFORM == blockchainPlatform) {
            tokenAddress := common.HexToAddress(tokens.Tokens[i].CONTRACT_ADDRESS)
            instance, err := token.NewToken(tokenAddress, client)
            if err != nil {
                log.Fatal(err)
            }
            name, err := instance.Name(&bind.CallOpts{})
            if err != nil {
                log.Fatal(err)
            }
            symbol, err := instance.Symbol(&bind.CallOpts{})
            if err != nil {
                log.Fatal(err)
            }
            decimals, err := instance.Decimals(&bind.CallOpts{})
            if err != nil {
                log.Fatal(err)
            }
            totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
            if err != nil {
                log.Fatal(err)
            }
            // Parsing the total supply
            ftotalSupply := parseDecimals(totalSupply,decimals)
            // Add the circulating supply of the tokens to the total to compute ETH total btc tokens
            total.Add(total,ftotalSupply)
            // Store the result of the audit/the comparison between BTC supply and token supply
            auditResult := executeAudit(tokens.Tokens[i].BTC_ADDRESSES,ftotalSupply)
            displayTokenInfos(name, symbol, ftotalSupply, tokens.Tokens[i].CONTRACT_ADDRESS, auditResult)
        }
    }
    printBitcoinComparison(total, blockchainPlatform)
    return total
}

func queryBitcoinInfosFromGecko() (CoinsMarketItem){
    body := httpRequest(IP_API_COINGECKO)
    data := CoinsMarket{}
    err := json.Unmarshal(body, &data)
    if err != nil {
        log.Fatalln(err)
    }

    return data[0]
}

func httpRequest(address string) ([]byte) {
    resp, err := http.Get(address)
    if err != nil {
      log.Fatalln(err)
    }
    //We Read the response body on the line below.
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }
    return body
}

func readConfigFile(configFile string) (Tokens){
    file, _ := os.Open(configFile)
    defer file.Close()
    decoder := json.NewDecoder(file)
    tokens := Tokens{}
    err := decoder.Decode(&tokens)
    if err != nil {
      fmt.Println("error reading conf file :", err)
    }
    return tokens
}

func printBitcoinComparison(total *big.Float, symbolToCompare string){
    btcInfos := queryBitcoinInfosFromGecko()
    ftotal, _ := total.Float64()
    fmt.Printf("Circulating supply of BTC : %2.f BTC \n", btcInfos.CirculatingSupply)
    fmt.Printf("Circulating supply of BTC on %s : %2.f BTC\n ", symbolToCompare, total)
    fmt.Printf("================== Percentage of BTC on %s : %.2f %% ================== \n", symbolToCompare ,percentOf(ftotal,btcInfos.CirculatingSupply))
}

func displayTokenInfos(name string,symbol string,totalSupply *big.Float, address string, auditResult string) {
    fmt.Println("Token contract address: " + address)
    fmt.Printf("Token name: %s\n", name)
    fmt.Printf("Token symbol: %s\n", symbol)
    fmt.Printf("Token totalSupply: %.2f\n", totalSupply)
    fmt.Printf("Audit result : %s\n\n", auditResult)
}

func parseDecimals(totalSupply *big.Int, decimals uint8) (*big.Float){
    tmp := new(big.Float)
    tmp.SetString(totalSupply.String())
    value := new(big.Float).Quo(tmp, big.NewFloat(math.Pow10(int(decimals))))
    return value
}

func percentOf(part float64, total float64) float64 {
    return (part * 100) / total
}

func executeAudit(btcAddresses []string, supplyOnToken *big.Float) (string){
    supplyOnBitcoin := big.NewFloat(0)
    fmt.Printf("Exexcuting audit on these btc addresses : %s \n", btcAddresses)
    for i := 0; i < len(btcAddresses); i++ {
        supplyOnBitcoin.Add(supplyOnBitcoin, getBalanceFromBitcoinAddress(btcAddresses[i]))
    }
    return verif(supplyOnBitcoin, supplyOnToken)
}

func getBalanceFromBitcoinAddress(btcAddress string) (*big.Float){
    fake := big.NewFloat(1) //TODO
    return fake
}

func verif(supplyOnBitcoin *big.Float, supplyOnToken *big.Float) (string){
    fsupplyOnBitcoin, _ := supplyOnBitcoin.Float64()
    fsupplyOnToken, _ := supplyOnToken.Float64()
    if(fsupplyOnBitcoin <= fsupplyOnToken){
        return "KO"
    }else{
        return "OK"
    }
}

// func queryBalanceOfTokenOnEthAddress(addressToQuery string, instance  *token.Token) (*big.Int) {
//     // To query balance of an address
//     address := common.HexToAddress(addressToQuery)
//     bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Printf("BalanceOf: %s\n", bal)
//     return bal
// }
