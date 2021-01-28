package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"transparence/pkg/blockchains"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

//Tx test
//const Tx = "5b958a8a200885779164562f94479d221dedeaae726c1d5744083d1ad16acd6e"//285e526be1b6392e8413e3e1594cf3f5c99c17ac20609e08223a446888e9e9c1"
//const Tx = "5574d698df9f0797dc7d558202137aef3d66ea8acea5777db11b55a97ad2f7e6" //"98b3bbd1cbe162f2b1bb1bbf71851f3162ea03b9a04838984565d148607a4ebc"
//const Tx = "2af596348e8b529a5f910d938b6b1f190a1ff5d9b6e3a77cdb7a03ef7d6f2052"
//const Tx = "fff2525b8931402dd09222c50775608f75787bd2b87e56995a7bdd30f79702c4" //2010 tx to test when reindex
//const Tx = "f5ae2ad8095c4a347adef38e7299a74df42099df2656d2f9af07f9cabb4fe24c" //2016 tx to test when reindex
//const Tx = "70364f6ef4fffec226a62f6ee7643669ac1364f203a60363e653340c2148c0f7" //coinbase example
//const Tx ="d7d2d42bef367f3b648faa60ffe688ae9084fab1ae781316a4e0f9062b6c8135" //previous coinbase spent
const Tx = "cea0dd0097e9e3afc63ddacba9496f8b19a35edd54e2bbabfa03673346cc4d30" //mint example
//https://keepscan.com/deposits/0x2e0758f781c8d125bdfeefdbd5d89428da4be357
//https://blockchair.com/bitcoin/address/bc1qvtv808v582jtnw0ejdy7mzjjsej78x6qzrpgc9

//Block of Tx
//const Block int64 = 100000 //652710 //648000
//const timestamp int64 = 1602685676
const Block int64 = 653897 //demo block

// main
func main() {
	rpcURL := flag.String("r", "127.0.0.1:8332", "rpc url:port ")
	userpass := flag.String("pwd", "user:pass", "rpc user:password")

	flag.Parse()

	user := strings.Split(*userpass, ":")[0]
	pass := strings.Split(*userpass, ":")[1]
	client, _ := bitcoin.New(*rpcURL, user, pass)

	trans, er := chainhash.NewHashFromStr(Tx)
	if er != nil {
		log.Fatal(er, " bad  tx hash")
	}
	fmt.Println("Tx", trans)

	transactionVerbose, erra := client.RPC.GetRawTransactionVerbose(trans)
	if erra != nil {
		log.Fatal(erra, "bad raw verbose tx format")
	}

	c,_:= client.ExtractVin(transactionVerbose.Vin)
	fmt.Println("vins pk:",c)
	fmt.Println("vouts pk", client.ExtractVout(transactionVerbose.Vout))

	blockHash, erra := client.RPC.GetBlockHash(Block)
	if erra != nil {
		log.Fatal(erra, "bad block hash request")

	}
	fmt.Println("block hash ", blockHash)

	blockHeader, erros := client.RPC.GetBlockHeader(blockHash)
	if erros != nil {
		log.Fatal(erros, "bad header request")

	}
	fmt.Println("Block timestamp ", blockHeader.Timestamp.Unix())
	fmt.Println("Block date ", blockHeader.Timestamp)

	Tx, erros := client.GetBlockTxHash(int64(Block))
	if erros != nil {
		log.Fatal(erros, "bad header request")

	}
	_ = Tx

	//fmt.Println("Block  Verbose Txs", Tx)
	/*
		blockTime := time.Unix(blockHeader.Timestamp.Unix(), 0)
		fmt.Println("block number ", Block)
		fmt.Print("range computed ")
		fmt.Println(client.GetRangeFromTimesTamp(blockTime))
	*/
	//fmt.Println(client.ExtractAddresses(Block))

}
