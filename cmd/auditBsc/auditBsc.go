package auditBsc

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	transparenceutils "transparence/pkg/TransparenceUtils"
	"transparence/pkg/erc20adapter"
	"github.com/ethereum/go-ethereum/ethclient"

	"transparence/pkg/tellor/tellorCaller"
)

const CONFIG_FILE = "./cmd/auditBsc/binancePeggedTokens.json"
const INFURA_KEY = ""
const IP_API_INFURA = "Https://mainnet.infura.io/v3/" + INFURA_KEY
const IP_API_BINANCESMARTCHAIN = "wss://bsc-ws-node.nariox.org:443"
const IP_API_BINANCECHAIN = "Https://dex.binance.org/api/v1/tokens"

func RunAudit(symbol string,tellor bool) {
	if symbol == "" {
		auditAllBinanceTokens(IP_API_INFURA)
	} else if symbol == "BTCB" {
		auditBTC()
	} else {
		auditSpecificBinanceToken(IP_API_INFURA,symbol,tellor)	
	}	
}

func auditAllBinanceTokens(ipAddress string) {
	tokens := transparenceutils.ReadConfigFile(CONFIG_FILE)
	client := erc20adapter.NewClientConnection(ipAddress)

	for i := 0; i < len(tokens.PeggedTokens); i++ {
		auditToken(tokens.PeggedTokens[i],client,false)
	}
}

func auditSpecificBinanceToken(ipAddress string, symbol string, tellor bool) {
	tokens := transparenceutils.ReadConfigFile(CONFIG_FILE)
	client := erc20adapter.NewClientConnection(ipAddress)

	for i := 0; i < len(tokens.PeggedTokens); i++ {
		if symbol == tokens.PeggedTokens[i].SYMBOL {
			auditToken(tokens.PeggedTokens[i],client, tellor)
		}

	}
}

func executeAuditOnEth(tokenAudited erc20adapter.PeggedToken, supplyOnEth *big.Float, tellor bool) {
	supplyOnBitcoin := big.NewFloat(0)
	tokenSupply := new(big.Float)

	binanceTokens, _ := getBinanceChainTokens()
	for i := 0; i < len(tokenAudited.BNC_PEGGED_ADDRESSES); i++ {
		for _, item := range binanceTokens {
			if item.ContractAddress == tokenAudited.BNC_PEGGED_ADDRESSES[i] {
				tokenSupply.SetString(item.TotalSupply)
				supplyOnBitcoin.Add(supplyOnBitcoin, tokenSupply)
			}
		}
	}

	client := erc20adapter.NewClientConnection(IP_API_BINANCESMARTCHAIN)

	for i := 0; i < len(tokenAudited.BSC_PEGGED_ADDRESSES); i++ {
		tokenAddress := common.HexToAddress(tokenAudited.BSC_PEGGED_ADDRESSES[i])
		instance := erc20adapter.NewToken(tokenAddress, client)
		decimals := erc20adapter.GetDecimals(instance)
		totalSupply := erc20adapter.GetTotalSupply(instance)
		ftotalSupply := transparenceutils.ParseDecimals(totalSupply, decimals)
		supplyOnBitcoin.Add(supplyOnBitcoin, ftotalSupply)
	}
	auditResult := transparenceutils.Verif(supplyOnEth, supplyOnBitcoin)
	fmt.Printf("- Reserve on Ethereum : %f \n- TotalSupply on BNC+BSC: %f\n---- Audit result : %t \n", supplyOnEth, supplyOnBitcoin, auditResult)
	if tellor == true {
		j := 0.0
		if auditResult{
			j=1.0
		}
		executeTellor(big.NewFloat(j))
	}
}

func auditBTC(){
	totalSupply := big.NewFloat(0)
	totalSupplyBnc := new(big.Float)

	tokens := transparenceutils.ReadConfigFile(CONFIG_FILE)
	clientEth := erc20adapter.NewClientConnection(IP_API_INFURA)
	//clientBnc := erc20adapter.NewClientConnection(IP_API_BINANCECHAIN)
	clientBsc := erc20adapter.NewClientConnection(IP_API_BINANCESMARTCHAIN)

	binanceTokens, _ := getBinanceChainTokens()

	for i := 0; i < len(tokens.PeggedTokens); i++ {
		if tokens.PeggedTokens[i].SYMBOL == "BTCB" {
			tokenAddressEth := common.HexToAddress(tokens.PeggedTokens[i].CONTRACT_ADDRESS)
			instanceEth := erc20adapter.NewToken(tokenAddressEth, clientEth)
			totalSupplyEth := erc20adapter.GetTotalSupply(instanceEth)
			decimals := erc20adapter.GetDecimals(instanceEth)
			ftotalSupplyEth := transparenceutils.ParseDecimals(totalSupplyEth, decimals)
			totalSupply.Add(totalSupply, ftotalSupplyEth)

			// tokenAddressBnc := common.HexToAddress(tokens.PeggedTokens[i].BNC_PEGGED_ADDRESSES[0])
			// instanceBnc := erc20adapter.NewToken(tokenAddressBnc, clientBnc)
			// totalSupplyBnc := erc20adapter.GetTotalSupply(instanceBnc)
			// decimals = erc20adapter.GetDecimals(instanceBnc)
			// ftotalSupply = transparenceutils.ParseDecimals(totalSupplyBnc, decimals)
			// totalSupply.Add(totalSupply, ftotalSupply)
			// fmt.Print(totalSupply)
			for _, item := range binanceTokens {
				if item.Symbol == "BTCB-1DE" {
					totalSupplyBnc.SetString(item.TotalSupply)
					totalSupply.Add(totalSupply, totalSupplyBnc)
				}
			}

			tokenAddressBsc := common.HexToAddress(tokens.PeggedTokens[i].BSC_PEGGED_ADDRESSES[0])
			instanceBsc := erc20adapter.NewToken(tokenAddressBsc, clientBsc)
			totalSupplyBsc := erc20adapter.GetTotalSupply(instanceBsc)
			decimals = erc20adapter.GetDecimals(instanceBsc)
			ftotalSupplyBsc := transparenceutils.ParseDecimals(totalSupplyBsc, decimals)
			totalSupply.Add(totalSupply, ftotalSupplyBsc)
			fmt.Printf("- TotalSupply on Ethereum : %f \n- TotalSupply on BNC: %f\n- TotalSupply on BSC: %f\n--- Overall TotalSupply: %f\n", ftotalSupplyEth, totalSupplyBnc, ftotalSupplyBsc,totalSupply)
		}

	}
}

func getBinanceChainTokens() ([]erc20adapter.BEP2Token, error) {
	var tokens []erc20adapter.BEP2Token
	body := transparenceutils.HttpRequest(IP_API_BINANCECHAIN)
	if err := json.Unmarshal(body, &tokens); err != nil {
		return nil, err
	}
	return tokens, nil
}

func auditToken(token erc20adapter.PeggedToken, client *ethclient.Client, tellor bool) {
	total := big.NewFloat(0)
	var name, symbol string
	tokenAddress := common.HexToAddress(token.CONTRACT_ADDRESS)
	instance := erc20adapter.NewToken(tokenAddress, client)

	if token.SYMBOL == "MKR" {
		name = "Maker"
		symbol = token.SYMBOL
	} else {
		name = erc20adapter.GetName(instance)
		symbol = erc20adapter.GetSymbol(instance)
	}

	decimals := erc20adapter.GetDecimals(instance)

	fmt.Printf("\n---- Auditing Binance Pegged Token : '%s'  (%s)\n", name, symbol)
	fmt.Printf("- Ethereum reserve addresses : %s \n", token.CONTRACT_ADDRESS)
	fmt.Printf("- Binance Chain reserve addresses : %s \n", token.BNC_PEGGED_ADDRESSES)
	fmt.Printf("- Binance Smart Chain reserve addresses : %s \n", token.BSC_PEGGED_ADDRESSES)

	totalReserve := erc20adapter.GetBalanceOfToken(token.RESERVE_ADDRESS, instance)
	ftotalReserve := transparenceutils.ParseDecimals(totalReserve, decimals)
	total.Add(total, ftotalReserve)
	executeAuditOnEth(token, ftotalReserve, tellor)
}



func executeTellor(value *big.Float){
	fmt.Printf("\n ================ Submitting to Tellor Oracle the Proof of Audit ================= \n")
	tellorRequestId := big.NewInt(50)
	totalOnEthInt := new(big.Int)
	totalOnEthInt.SetString(value.String(), 10)
	tellorCaller.UpdateTellorPlaygroundValue(totalOnEthInt, tellorRequestId)
	tellorCaller.GetTellorPlaygroundValue(tellorRequestId)

	// btcPriceOnTellor := tellorCaller.GetTellorValue(TELLOR_ID_BTC_PRICE)
	// btcPriceOnTellorf := transparenceutils.ParseDecimals(btcPriceOnTellor, 6)
	// btcOnEthUsdValue := new(big.Float)
	// btcOnEthUsdValue = btcOnEthUsdValue.Mul(totalOnEth, btcPriceOnTellorf)
	// fmt.Printf("\n According to Bitcoin's price on Tellor mainnet ($%2.f) that makes around $%2.f", btcPriceOnTellorf, btcOnEthUsdValue)

}