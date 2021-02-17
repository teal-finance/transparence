package main
import (
    "context"
    "fmt"
    "log"
    "math"
    "math/big"
    "flag"

    "github.com/ethereum/go-ethereum/common"
    "transparence/pkg/blockchains/ethereum"
)
func main() {
	rpcURL := flag.String("r", "127.0.0.1:8545", "rpc url ")

	flag.Parse()
	client, err := ethereum.New(*rpcURL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected")

  account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
   balance, err := client.RPC.BalanceAt(context.Background(), account, nil)
   if err != nil {
       log.Fatal(err)
   }
   fmt.Println(balance) // 25893180161173005034

   blockNumber := big.NewInt(5532993)
   balanceAt, err := client.RPC.BalanceAt(context.Background(), account, blockNumber)
   if err != nil {
       log.Fatal(err)
   }
   fmt.Println(balanceAt) // 25729324269165216042

   fbalance := new(big.Float)
   fbalance.SetString(balanceAt.String())
   ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
   fmt.Println(ethValue) // 25.729324269165216041

   pendingBalance, _ := client.RPC.PendingBalanceAt(context.Background(), account)
   fmt.Println(pendingBalance) // 25729324269165216042
}
