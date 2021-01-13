package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"strings"

	"transparence/pkg/blockchains"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
)

//Tx test
const Tx = "5b958a8a200885779164562f94479d221dedeaae726c1d5744083d1ad16acd6e"//285e526be1b6392e8413e3e1594cf3f5c99c17ac20609e08223a446888e9e9c1"
//const Tx = "5574d698df9f0797dc7d558202137aef3d66ea8acea5777db11b55a97ad2f7e6" //"98b3bbd1cbe162f2b1bb1bbf71851f3162ea03b9a04838984565d148607a4ebc"
//const Tx = "2af596348e8b529a5f910d938b6b1f190a1ff5d9b6e3a77cdb7a03ef7d6f2052"
//const Tx = "fff2525b8931402dd09222c50775608f75787bd2b87e56995a7bdd30f79702c4" //2010 tx to test when reindex
//const Tx = "f5ae2ad8095c4a347adef38e7299a74df42099df2656d2f9af07f9cabb4fe24c" //2016 tx to test when reindex
//const Tx = "70364f6ef4fffec226a62f6ee7643669ac1364f203a60363e653340c2148c0f7" //coinbase example
//const Tx = "cea0dd0097e9e3afc63ddacba9496f8b19a35edd54e2bbabfa03673346cc4d30" //mint example
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

	// Start with a standard pay-to-pubkey-hash script.
	scriptHex := "76a914128004ff2fcaf13b2b91eb654b1dc2b674f7ec6188ac"
	script, err := hex.DecodeString(scriptHex)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Extract and print details from the script.
	scriptClass, addresses, reqSigs, err := txscript.ExtractPkScriptAddrs(
		script, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Script Class:", scriptClass)
	fmt.Println("Addresses:", addresses)
	fmt.Println("Required Signatures:", reqSigs)

	user := strings.Split(*userpass, ":")[0]
	pass := strings.Split(*userpass, ":")[1]
	client, _ := keeper.New(*rpcURL, user, pass)
	//fmt.Println(client.RPC.GetBestBlockHash())

	trans, er := chainhash.NewHashFromStr(Tx)
	if er != nil {
		log.Fatal(er, " bad  tx hash")
	}
	fmt.Println("Tx", trans)

	/*transaction, errwin := client.RPC.GetRawTransaction(trans)
	if errwin != nil {
		log.Fatal(errwin, "bad raw tx format")
	}
	fmt.Println("RAw Tx", transaction)
	*/
	transactionVerbose, erra := client.RPC.GetRawTransactionVerbose(trans)
	if erra != nil {
		log.Fatal(erra, "bad raw verbose tx format")
	}
	fmt.Println("vin ", transactionVerbose.Vin)
	fmt.Println("vout ", transactionVerbose.Vout)
	a:=transactionVerbose.Vin[0]
	fmt.Println("prevout:",a)
	//_ = transactionVerbose
	//fmt.Println("Raw verbose Tx", transactionVerbose)
	c,_:= client.ExtractVin2(transactionVerbose.Vin)
	fmt.Println("vin pk:",c)
	fmt.Println("vout pk", client.ExtractVout(transactionVerbose.Vout))

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
