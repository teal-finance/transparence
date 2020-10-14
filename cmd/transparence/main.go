package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"transparence/pkg/keeper"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

//Tx test
const Tx = "5574d698df9f0797dc7d558202137aef3d66ea8acea5777db11b55a97ad2f7e6" //"98b3bbd1cbe162f2b1bb1bbf71851f3162ea03b9a04838984565d148607a4ebc"

// main
func main() {
	rpc := flag.String("r", "127.0.0.1:8332", "rpc url:port ")
	userpass := flag.String("pwd", "user:pass", "rpc user:password")

	flag.Parse()

	user := strings.Split(*userpass, ":")[0]
	pass := strings.Split(*userpass, ":")[1]
	client := keeper.BtcClient(*rpc, user, pass)

	fmt.Println(client.GetBestBlockHash())
	trans, er := chainhash.NewHashFromStr(Tx)
	if er != nil {
		log.Fatal(er, "bad  tx hash")
	}
	fmt.Println(trans)

	transaction, errwin := client.GetRawTransaction(trans)
	if errwin != nil {
		log.Fatal(errwin, "bad raw tx format")
	}
	fmt.Println(transaction)

	transactionVerbose, erra := client.GetRawTransactionVerbose(trans)
	if errwin != nil {
		log.Fatal(erra, "bad raw verbose tx format")
	}
	fmt.Println("vin ", transactionVerbose.Vin)
	fmt.Println("vout ", transactionVerbose.Vout)

}
