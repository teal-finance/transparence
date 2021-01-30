package main

import (
    "encoding/json"
    "os"

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

const CONFIG_FILE = "binancePeggedTokens.json"
const INFURA_KEY=""
const IP_API_INFURA = "https://mainnet.infura.io/v3/" + INFURA_KEY
const IP_API_BINANCESMARTCHAIN = "wss://bsc-ws-node.nariox.org:443"
const IP_API_BINANCECHAIN = "https://dex.binance.org/api/v1/tokens"

// Types used to represent the tokens of the JSON config file
type Token struct {
    BLOCKCHAIN_PLATFORM string `json:"BLOCKCHAIN_PLATFORM"`
    SYMBOL string `json:"SYMBOL"`
    CONTRACT_ADDRESS   string `json:"CONTRACT_ADDRESS"`
    RESERVE_ADDRESS   string `json:"RESERVE_ADDRESS"`
    BNC_PEGGED_ADDRESSES []string `json:"BNC_PEGGED_ADDRESSES"`
    BSC_PEGGED_ADDRESSES []string `json:"BSC_PEGGED_ADDRESSES"`
}

type BEP2Token struct {
    Name             string     `json:"name"`
    Symbol           string     `json:"symbol"`
    OrigSymbol       string     `json:"original_symbol"`
    TotalSupply      string     `json:"total_supply"`
    Owner            string `json:"owner"`
    Mintable         bool       `json:"mintable"`
    ContractAddress  string     `json:"contract_address,omitempty"`
    ContractDecimals int8       `json:"contract_decimals,omitempty"`
}

type Tokens struct {
    Tokens []Token `json:"tokens"`
}

func main() {
  auditBinanceTokens(IP_API_INFURA)
}

func auditBinanceTokens(ipAddress string){
    total := big.NewFloat(0)
    var name, symbol string

    tokens := readConfigFile(CONFIG_FILE)
    client, err := ethclient.Dial(ipAddress)
    if err != nil {
        log.Fatal(err)
    }
    for i := 0; i < len(tokens.Tokens); i++ {
        tokenAddress := common.HexToAddress(tokens.Tokens[i].CONTRACT_ADDRESS)
        instance, err := token.NewToken(tokenAddress, client)
        if err != nil {
            log.Fatal(err)
        }

        if tokens.Tokens[i].SYMBOL == "MKR" {
            name = "Maker"
            symbol = tokens.Tokens[i].SYMBOL
        } else {
            name, err = instance.Name(&bind.CallOpts{})
            if err != nil {
                log.Fatal(err)
            }
            symbol, err = instance.Symbol(&bind.CallOpts{})
            if err != nil {
                log.Fatal(err)
            }
        }

        decimals, err := instance.Decimals(&bind.CallOpts{})
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("\n---- Auditing Binance Pegged Token : '%s'  (%s)\n",name,symbol)
        fmt.Printf("- Ethereum reserve addresses : %s \n", tokens.Tokens[i].CONTRACT_ADDRESS)
        fmt.Printf("- Binance Chain reserve addresses : %s \n", tokens.Tokens[i].BNC_PEGGED_ADDRESSES)
        fmt.Printf("- Binance Smart Chain reserve addresses : %s \n", tokens.Tokens[i].BSC_PEGGED_ADDRESSES)

        totalReserve := queryBalanceOfTokenOnEthAddress(tokens.Tokens[i].RESERVE_ADDRESS,instance)
        ftotalReserve := parseDecimals(totalReserve,decimals)
        total.Add(total,ftotalReserve)
        executeAuditOnEth(tokens.Tokens[i],ftotalReserve)
    }
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

func parseDecimals(totalSupply *big.Int, decimals uint8) (*big.Float){
    tmp := new(big.Float)
    tmp.SetString(totalSupply.String())
    value := new(big.Float).Quo(tmp, big.NewFloat(math.Pow10(int(decimals))))
    return value
}

func executeAuditOnEth(tokenAudited Token, supplyOnEth *big.Float) (string){
    supplyOnBitcoin := big.NewFloat(0)
    tokenSupply := new(big.Float)

    binanceTokens,_ := getBinanceChainTokens()
    for i := 0; i < len(tokenAudited.BNC_PEGGED_ADDRESSES); i++ {
        for _, item := range binanceTokens {
            if item.ContractAddress == tokenAudited.BNC_PEGGED_ADDRESSES[i] {
                tokenSupply.SetString(item.TotalSupply)
                supplyOnBitcoin.Add(supplyOnBitcoin, tokenSupply)
            }
        } 
    }

    client, err := ethclient.Dial(IP_API_BINANCESMARTCHAIN)
    if err != nil {
        log.Fatal(err)
    }
    for i := 0; i < len(tokenAudited.BSC_PEGGED_ADDRESSES); i++ {
        tokenAddress := common.HexToAddress(tokenAudited.BSC_PEGGED_ADDRESSES[i])
        instance, err := token.NewToken(tokenAddress, client)
        if err != nil {
            log.Fatal(err)
        }
        totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
        if err != nil {
            log.Fatal(err)
        }

        decimals, err := instance.Decimals(&bind.CallOpts{})
        if err != nil {
            log.Fatal(err)
        }
        ftotalSupply := parseDecimals(totalSupply,decimals)
        supplyOnBitcoin.Add(supplyOnBitcoin, ftotalSupply)
    }
    return verif(supplyOnEth,supplyOnBitcoin)
}

func verif(supplyOnEth *big.Float, supplyOnToken *big.Float) (string){
    fsupplyOnEth, _ := supplyOnEth.Float64()
    fsupplyOnToken, _ := supplyOnToken.Float64()
    if(fsupplyOnToken <= fsupplyOnEth){
        fmt.Printf("- Reserve on Ethereum : %f \n- TotalSupply on BNC+BSC: %f\n---- Audit result : 'OK' \n",supplyOnEth,supplyOnToken)
        return "OK"
    }else{
        fmt.Printf("- Reserve on Ethereum : %f \n- TotalSupply on BNC+BSC: %f \n---- Audit result : 'KO' \n",supplyOnEth,supplyOnToken)
        return "KO"
    }
}

func queryBalanceOfTokenOnEthAddress(addressToQuery string, instance  *token.Token) (*big.Int) {
    address := common.HexToAddress(addressToQuery)
    bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
    if err != nil {
        log.Fatal(err)
    }
    return bal
}

func getBinanceChainTokens() ([]BEP2Token, error){
    var tokens []BEP2Token
    body := httpRequest(IP_API_BINANCECHAIN)
    if err := json.Unmarshal(body, &tokens); err != nil {
        return nil, err
    }
    return tokens, nil
}

func find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}