package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"transparence/pkg/keeper"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

//Tx test
//const Tx = "5574d698df9f0797dc7d558202137aef3d66ea8acea5777db11b55a97ad2f7e6" //"98b3bbd1cbe162f2b1bb1bbf71851f3162ea03b9a04838984565d148607a4ebc"
//const Tx = "2af596348e8b529a5f910d938b6b1f190a1ff5d9b6e3a77cdb7a03ef7d6f2052"
const Tx = "70364f6ef4fffec226a62f6ee7643669ac1364f203a60363e653340c2148c0f7" //coinbase example
//Block of Tx
const Block int64 = 100000 //652710 //648000
//const timestamp int64 = 1602685676

// main
func main() {
	rpcURL := flag.String("r", "127.0.0.1:8332", "rpc url:port ")
	userpass := flag.String("pwd", "user:pass", "rpc user:password")

	flag.Parse()

	user := strings.Split(*userpass, ":")[0]
	pass := strings.Split(*userpass, ":")[1]
	client, _ := keeper.New(*rpcURL, user, pass)
	fmt.Println(client.RPC.GetBestBlockHash())

	trans, er := chainhash.NewHashFromStr(Tx)
	if er != nil {
		log.Fatal(er, "bad  tx hash")
	}
	fmt.Println("Tx", trans)

	transaction, errwin := client.RPC.GetRawTransaction(trans)
	if errwin != nil {
		log.Fatal(errwin, "bad raw tx format")
	}
	fmt.Println("RAw Tx", transaction)

	transactionVerbose, erra := client.RPC.GetRawTransactionVerbose(trans)
	if erra != nil {
		log.Fatal(erra, "bad raw verbose tx format")
	}
	fmt.Println("vin ", transactionVerbose.Vin[0].IsCoinBase())
	fmt.Println("vout ", transactionVerbose.Vout)
	//_ = transactionVerbose
	fmt.Println("RAw verbose Tx", transactionVerbose)

	//fmt.Println(client.ExtractVin(transactionVerbose.Vin))
	//fmt.Println(client.ExtractVout(transactionVerbose.Vout))

	blockHash, erra := client.RPC.GetBlockHash(Block)
	if erra != nil {
		log.Fatal(erra, "bad block hash request")

	}
	fmt.Println("block hash ", blockHash)

	blockHeader, erros := client.RPC.GetBlockHeader(blockHash)
	if erros != nil {
		log.Fatal(erros, "bad header request")

	}
	fmt.Println("Block timestamp ", blockHeader.Timestamp.Unix(), blockHeader.Timestamp)

	Tx, erros := client.GetBlockTxHash(int64(Block))
	if erros != nil {
		log.Fatal(erros, "bad header request")

	}
	_ = Tx

	//fmt.Println("Block  Verbose Txs", Tx)

	blockTime := time.Unix(blockHeader.Timestamp.Unix(), 0)
	fmt.Println(client.GetRangeFromTimesTamp(blockTime))

	fmt.Println(client.ExtractAddresses(Block))

}
