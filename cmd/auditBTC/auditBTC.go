package main

import (
	"flag"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	transparenceutils "transparence/pkg/TransparenceUtils"
	"transparence/pkg/coingecko"
	"transparence/pkg/ethereum/erc20"
	"transparence/pkg/tellor/tellorCaller"
)

const CONFIG_FILE = "config.json"
const INFURA_KEY = ""
const IP_API_INFURA = "https://mainnet.infura.io/v3/" + INFURA_KEY
const IP_API_BINANCECHAIN = "wss://bsc-ws-node.nariox.org:443"

const ETH_STRING = "Ethereum"
const BINANCE_STRING = "BinanceChain"

var TELLOR_ID_BTC_PRICE = big.NewInt(2)

func main() {
	rpcURL := flag.String("r", "127.0.0.1:8545", "rpc url ")

	flag.Parse()
	fmt.Printf("\n\n ################## Analysis of Ethereum ERC20 BTC ################## \n\n")
	totalOnEth := analyze(*rpcURL, ETH_STRING)

	fmt.Printf("\n ================ Submit to Tellor Oracle the total of ERC20 BTCs ================= \n")
	tellorRequestId := big.NewInt(50)
	totalOnEthInt := new(big.Int)
	totalOnEthInt.SetString(totalOnEth.String(), 10)
	tellorCaller.UpdateTellorPlaygroundValue(totalOnEthInt, tellorRequestId)
	tellorCaller.GetTellorPlaygroundValue(tellorRequestId)

	btcPriceOnTellor := tellorCaller.GetTellorValue(TELLOR_ID_BTC_PRICE)
	btcPriceOnTellorf := transparenceutils.ParseDecimals(btcPriceOnTellor, 6)
	btcOnEthUsdValue := new(big.Float)
	btcOnEthUsdValue = btcOnEthUsdValue.Mul(totalOnEth, btcPriceOnTellorf)
	fmt.Printf("\n According to Bitcoin's price on Tellor mainnet ($%2.f) that makes around $%2.f", btcPriceOnTellorf, btcOnEthUsdValue)

	fmt.Printf("\n\n ################## Analysis of Binance chain BEP20 BTC ################## \n\n")
	totalOnBinance := analyze(IP_API_BINANCECHAIN, BINANCE_STRING)

	fmt.Printf("\n\n ################## Analysis of Ethereum ERC20 & Binance chain BEP20 BTC ################## \n\n")
	totalOverall := totalOnEth.Add(totalOnEth, totalOnBinance)
	printBitcoinComparison(totalOverall, "Ethereum & Binance chain")
}

func analyze(ipAddress string, blockchainPlatform string) *big.Float {
	total := big.NewFloat(0)

	tokens := transparenceutils.ReadConfigFile(CONFIG_FILE)
	client := erc20.NewClientConnection(ipAddress)
	for i := 0; i < len(tokens.PeggedTokens); i++ {
		if tokens.PeggedTokens[i].BLOCKCHAIN_PLATFORM == blockchainPlatform {
			tokenAddress := common.HexToAddress(tokens.PeggedTokens[i].CONTRACT_ADDRESS)
			instance := erc20.NewToken(tokenAddress, client)
			name := erc20.GetName(instance)
			symbol := erc20.GetSymbol(instance)
			decimals := erc20.GetDecimals(instance)
			totalSupply := erc20.GetTotalSupply(instance)

			// Parsing the total supply
			ftotalSupply := transparenceutils.ParseDecimals(totalSupply, decimals)
			// Add the circulating supply of the tokens to the total to compute ETH total btc tokens
			total.Add(total, ftotalSupply)
			// Store the result of the audit/the comparison between BTC supply and token supply
			auditResult := executeAudit(tokens.PeggedTokens[i].BTC_ADDRESSES, ftotalSupply)
			displayTokenInfos(name, symbol, ftotalSupply, tokens.PeggedTokens[i].CONTRACT_ADDRESS, auditResult)
		}
	}
	printBitcoinComparison(total, blockchainPlatform)
	return total
}

func printBitcoinComparison(total *big.Float, symbolToCompare string) {
	btcInfos := coingecko.QueryBitcoinInfosFromGecko()
	ftotal, _ := total.Float64()
	fmt.Printf("Circulating supply of BTC : %2.f BTC \n", btcInfos.CirculatingSupply)
	fmt.Printf("Circulating supply of BTC on %s : %2.f BTC\n ", symbolToCompare, total)
	fmt.Printf("================== Percentage of BTC on %s : %.2f %% ================== \n", symbolToCompare, transparenceutils.PercentOf(ftotal, btcInfos.CirculatingSupply))
}

func displayTokenInfos(name string, symbol string, totalSupply *big.Float, address string, auditResult string) {
	fmt.Println("Token contract address: " + address)
	fmt.Printf("Token name: %s\n", name)
	fmt.Printf("Token symbol: %s\n", symbol)
	fmt.Printf("Token totalSupply: %.2f\n", totalSupply)
	fmt.Printf("Audit result : %s\n\n", auditResult)
}

func executeAudit(btcAddresses []string, supplyOnToken *big.Float) string {
	supplyOnReserve := big.NewFloat(0)
	fmt.Printf("Exexcuting audit on these btc addresses : %s \n", btcAddresses)
	for i := 0; i < len(btcAddresses); i++ {
		supplyOnReserve.Add(supplyOnReserve, getBalanceFromBitcoinAddress(btcAddresses[i]))
	}
	return transparenceutils.Verif(supplyOnReserve, supplyOnToken)
}

func getBalanceFromBitcoinAddress(btcAddress string) *big.Float {
	fake := big.NewFloat(1) //TODO
	return fake
}
