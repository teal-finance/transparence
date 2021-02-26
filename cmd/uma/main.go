package main
import (
    "context"
    "fmt"
    "sort"
    "log"
    "math/big"
    "flag"

    "github.com/ethereum/go-ethereum/common"
    "github.com/superhawk610/bar"
	  "github.com/ttacon/chalk"
    "transparence/pkg/blockchains/ethereum"
)

const start = 11766939
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

  blockStart := big.NewInt(start)
  end,err:=client.RPC.BlockNumber(context.Background())
  if err != nil {
      log.Fatal(err)
  }
  blockEnd:=big.NewInt(int64(end))
  b := bar.NewWithOpts(
		bar.WithDimensions(int(end-start+1), 50),
		bar.WithFormat(
			fmt.Sprintf(
				"   %sloading...%s :percent :bar %s:rate ops/s%s ",
				chalk.Blue,
				chalk.Reset,
				chalk.Green,
				chalk.Reset,
			),
		),
	)
  var TXes transactions
  var totalGasUsed uint64
  for i:=blockStart.Int64();i<=blockEnd.Int64();i++{
    block,err:=client.RPC.BlockByNumber(context.Background(),big.NewInt(i))
    if err != nil {
        log.Fatal(err)
    }
    //fmt.Println(block)
    //fmt.Println(block.Header())
    b.Tick()
    totalGasUsed+=block.GasUsed()
    body:=block.Body()
    for _,tx:=range body.Transactions{
      TXes=append(TXes,transaction{tx.GasPrice(),tx.Gas()})
    }
  }
  b.Done()
  fmt.Println("Total Gas used ",totalGasUsed)
  //fmt.Println(TXes)
  sort.Slice(TXes, func(i, j int) bool {
    return TXes[i].gasPrice.Cmp(TXes[j].gasPrice)<0
  })
  //fmt.Println(TXes)
  med:=median(TXes,totalGasUsed)
  //medianGwei:=med.Div()
  fmt.Println("Median Gas Price: ",med)

}

type transaction struct{
  gasPrice *big.Int
  gasUsed  uint64
}
type transactions []transaction

func median(sortedTxs transactions, totalGasUsed uint64) *big.Int{
  var rollingSum uint64
  for _,tx:= range sortedTxs{
    rollingSum+=tx.gasUsed
    if rollingSum>totalGasUsed/2{
      return tx.gasPrice
      }
  }
  return big.NewInt(0)
}
