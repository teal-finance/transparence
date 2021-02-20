package bitcoin

import (
	"encoding/hex"
	"encoding/base64"
	"errors"
	"fmt"
	"time"
	"strings"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/ybbus/jsonrpc/v2"
)

// BtcClient ...
type BtcClient struct {
	RPC *rpcclient.Client
}

//New create a new client
func New(rpcURL, user, pass string) (*BtcClient, error) {
	// create new client instance
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         rpcURL,
		User:         user,
		Pass:         pass,
	}, nil)

	return &BtcClient{client}, err
}

//GetBlockTxHash  ...
func (c *BtcClient) GetBlockTxHash(blockNumber int64) ([]string, error) {
	blockHash, err := c.RPC.GetBlockHash(blockNumber)
	if err != nil {
		return nil, err
	}

	blockVerbose, erros := c.RPC.GetBlockVerbose(blockHash)
	if erros != nil {
		return nil, err

	}

	return blockVerbose.Tx, nil
}

//GetBlockTimestamp ...
func (c *BtcClient) GetBlockTimestamp(blockNumber int64) (time.Time, error) {
	hash, err := c.RPC.GetBlockHash(blockNumber)
	if err != nil {
		return time.Time{}, err
	}

	header, err := c.RPC.GetBlockHeader(hash)
	if err != nil {
		return time.Time{}, err
	}

	return header.Timestamp, nil
}

// KeepBlockGenesis is block on Bitcoin blockchain
// a lilttle before the contracts creation on Ethereum blokhain
// The address which deployed the contracts was funded a bit before that block
//source: https://etherscan.io/txs?a=0x123694886dbf5ac94dda07135349534536d14caf
const KeepBlockGenesis int64 = 648000

// KeepBlockGenesisTime is the timestamp of KeepBlockGenesis
const KeepBlockGenesisTime int64 = 1600186924

// GetRangeFromTimesTamp ...
func (c *BtcClient) GetRangeFromTimesTamp(t time.Time) (int64, int64, error) {
	lastBlock, err := c.RPC.GetBlockCount()
	if err != nil {
		return 0, 0, err
	}
	//fmt.Println("last block ", lastBlock)
	return c.getBlockTimestamp(t, KeepBlockGenesis, lastBlock)

}

//period= in seconds.. 12h * 6blocks/h * 10min/block * 60s/min
const period int64 = 12 * 6 * 10 * 60

// getBlockTimestamp is a quick a very dirty function that tries to find a
// range of block around a timestamp
func (c *BtcClient) getBlockTimestamp(t time.Time, start int64, end int64) (int64, int64, error) {
	startTime, err := c.GetBlockTimestamp(start)
	if err != nil {
		return 0, 0, err
	}
	endTime, err := c.GetBlockTimestamp(end)
	if err != nil {
		return 0, 0, err
	}

	//fmt.Println("start block ", start, "timestamp", startTime)
	//fmt.Println("end block ", end, " timestamp", endTime)

	if startTime.Unix() > t.Unix() || t.Unix() > endTime.Unix() {
		return start, end, errors.New("Timestamp out of Range")
	}

	if startTime.Unix() < t.Unix()-period && t.Unix()+period < endTime.Unix() {

		return c.getBlockTimestamp(t, moveGoalpostsInf(startTime, t, start), moveGoalpostsSup(t, endTime, end))
	}

	if startTime.Unix() < t.Unix()-period {
		return c.getBlockTimestamp(t, moveGoalpostsInf(startTime, t, start), end)
	}
	if t.Unix()+period < endTime.Unix() {
		return c.getBlockTimestamp(t, start, moveGoalpostsSup(t, endTime, end))
	}

	return start, end, nil

}

// safeStep=4 because we never know with hashrate and block estimation in a certain period of time
const safeStep int64 = 4

// moveGoalpostsInf assumes t1 < t2 and we want to get t1(and b1) closer to t2
// Assuming there is no drop of 1/safeStep the hashrate, we can advance in slow steps of safety
func moveGoalpostsInf(t1, t2 time.Time, b1 int64) int64 {
	numberOfBlocks := (t2.Unix() - t1.Unix()) / (10 * 60 * safeStep)
	return b1 + numberOfBlocks
}

// moveGoalpostsSup assumes t1 < t2 and we want to get t2(and b2) closer to t1
// Assuming there is no drop of 1/safeStep the hashrate, we can advance in slow steps of safety
func moveGoalpostsSup(t1, t2 time.Time, b2 int64) int64 {
	numberOfBlocks := (t2.Unix() - t1.Unix()) / (10 * 60 * safeStep)
	return b2 - numberOfBlocks
}

// ExtractVout extract the list of addresses involved.
// The Tx hash should be taken care of by the caller of this function
// No call to rpc so  error just to have the same signature as ExtractVin
func (c *BtcClient) ExtractVout(vouts []btcjson.Vout) ([]string, error) {
	addresses := make([]string, 0)
	if vouts == nil {
		return addresses, nil
	}
	for _, vout := range vouts {
		addresses = append(addresses, vout.ScriptPubKey.Addresses...)
	}
	return addresses, nil
}

// ExtractVin extract the list of addresses involved.
// The Tx hash should be taken care of by the caller of this function
// we also need to think of the coinbase tx that is the first tx, vin & vout, in each block
// inspiration mainly comes from: https://github.com/btcsuite/btcd/blob/master/rpcserver.go
// and I am too tired to explain how it works.
func (c *BtcClient) ExtractVin(vins []btcjson.Vin) ([]string, error) {
	addresses := make([]string, 0)
	//this should never happen actually
	if vins == nil {
		return addresses, nil
	}
	for _, vin := range vins {
		// the easy way to get the address in Vin is to look at Vout

		//watchout for coinbase transactions
		// because they don't have TxId and only have coinbase field :-(
		if vin.IsCoinBase() {
			//pretty sure we can return here but might be some weird blocks
			//with weird coinbase tx
			continue
		}
		hash, err := chainhash.NewHashFromStr(vin.Txid)
		if err != nil {
			return []string{}, err
		}
		transaction, err := c.RPC.GetRawTransactionVerbose(hash)
		if err != nil {
			return []string{}, err
		}
		tx := transaction.Vout[vin.Vout]
		scriptHex := tx.ScriptPubKey.Hex
		script, err := hex.DecodeString(scriptHex)
		if err != nil {
			fmt.Println(err)
			return []string{}, err
		}

		_, addrs, _, err := txscript.ExtractPkScriptAddrs(
			script, &chaincfg.MainNetParams)
		if err != nil {
			fmt.Println(err)
			return []string{}, err
		}
		for _, addr := range addrs {
			addresses = append(addresses, addr.EncodeAddress())
		}
	}
	return addresses, nil
}

// ExtractAddressesFromTx an array of the addresses present in the Tx
// vin (inputs) and vout (outputs)
func (c *BtcClient) ExtractAddressesFromTx(txHash string) ([]string, error) {
	addresses := make([]string, 0)
	tx, er := chainhash.NewHashFromStr(txHash)
	if er != nil {
		return addresses, er
	}

	transactionVerbose, erra := c.RPC.GetRawTransactionVerbose(tx)
	if erra != nil {
		return addresses, erra
	}

	vins, erro := c.ExtractVin(transactionVerbose.Vin)
	if erro != nil {
		return addresses, erro
	}
	vouts, err := c.ExtractVout(transactionVerbose.Vout)
	if err != nil {
		return addresses, err
	}
	addresses = append(addresses, vins...)
	addresses = append(addresses, vouts...)

	return addresses, nil
}

// ImportAddressRescan adds address to be watched in the client
func (c *BtcClient) ImportAddressRescan(address string, account string, rescan bool) rpcclient.FutureImportAddressResult {
	return c.RPC.ImportAddressRescanAsync(address, account, rescan)
}

//GetBalance of address , for multiple addresses, use  GetBalances
func (c *BtcClient) GetBalance(addr string)(float64,error){

	return c.GetBalances([]string{addr})
}

//GetBalances of all the  addresses in the array
func (c *BtcClient) GetBalances(addresses []string)(float64,error){
	defaultNet := &chaincfg.MainNetParams
	arrayAddresses := []btcutil.Address{}

	for _,addr:= range addresses{
		addressDecoded, err := btcutil.DecodeAddress(addr, defaultNet)
		if err != nil {
			return 0.0,err
		}
		arrayAddresses = append(arrayAddresses,addressDecoded)
	}

	listUnspentOutputs,err:=c.RPC.ListUnspentMinMaxAddresses(1,999999,arrayAddresses)
	if err != nil {
		return 0.0,err
	}
	var  balance float64
	for _,v:= range listUnspentOutputs{
		balance+=v.Amount
	}
	return balance,nil
}


//AltRPC because btcd does not support all function from Bitcoin Core
//and also I was not able to use btcd's RawRequest.
type AltRPC struct  {
	RPC jsonrpc.RPCClient

}

// NewAlt create an alternative rpc client to btcd RPC client
func NewAlt(rpcURL, user, pass string) (*AltRPC, error) {
	if !(strings.HasPrefix(rpcURL,"http://")){
		rpcURL="http://"+rpcURL
	}
	client := jsonrpc.NewClientWithOpts(rpcURL, &jsonrpc.RPCClientOpts{
   		CustomHeaders: map[string]string{
   			"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+pass)),
   		},
   	})
		//dummy call to see an error early, if not we have to wait when the call is made
	_, err := client.Call("getrpcinfo")
	return &AltRPC{client},err
}

//GetAddressesByLabel return an array of the addresses of a label saved on the Bitcoin node
// use ImportAddressRescan first with the node RPC
func (a AltRPC) GetAddressesByLabel(label string) ([]string,error){
	var listAddr []string
	response, err := a.RPC.Call("getaddressesbylabel", label)
	if err!=nil{
		return listAddr,err
	}
	//type assertion because we know what to expect
	var respmap map[string]interface{} = response.Result.(map[string]interface{})
	for k:=range respmap{
		listAddr=append(listAddr,k)
	}
	return listAddr,err

}
/*func (a AltRPC) GetBalance(address string) (float32,error){
	return 0.0,nil
}

func (a AltRPC) GetBalanceByLabel(label string) (float32,error){
	var balance float32

	listAddr,err:=a.GetAddressesByLabel(label)
	if  err!=nil{
		return balance,err
	}

	for _,addr:= range listAddr{
		b,err:=a.GetBalance(addr)
		if  err!=nil{
			return 0.0,err
		}
		balance+=b
	}
	return balance,err
}

*/

//TODO getaddressinfo (note that we can only know more details, like required sig,  once it is used )
/*
// ExtractAddresses an array of the addresses present in the block. the index follow the tx indexes in the block
func (c *BtcClient) ExtractAddresses(block int64) ([][]string, error) {
	transactions, erros := c.GetBlockTxHash(block)
	if erros != nil {
		return [][]string{}, erros
	}
	addresses := make([][]string, 0)
	if transactions == nil {
		return addresses, nil
	}
	for _, tx := range transactions {
		hash, er := chainhash.NewHashFromStr(tx)
		if er != nil {
			return [][]string{}, er
		}
		transactionVerbose, erra := c.RPC.GetRawTransactionVerbose(hash)
		if erra != nil {
			return [][]string{}, erra
		}

		addressesVinVout, err := c.ExtractVin(transactionVerbose.Vin)
		if err != nil {
			return [][]string{}, err
		}
		//addressesVinVout = append(addressesVinVout, c.ExtractVout(transactionVerbose.Vout)...)
		addresses = append(addresses, [][]string{addressesVinVout}...)
	}

	return addresses, nil
}*/

/*
//Transparence shows the tx involving around a timestamp
func (c *BtcClient) Transparence(t time.Time, userAddress string) error {
	start, end, err := c.GetRangeFromTimesTamp(t)
	if err != nil {
		return err
	}
	for i := start; i <= end; i++ {
		blockTx, err := c.GetBlockTxHash(i)
		if err != nil {
			return err
		}
		addresses, err := c.ExtractAddresses(i)
		if err != nil {
			return err
		}
		for index, txAddresses := range addresses {
			if txAddresses != nil {
				for _, address := range txAddresses {
					if userAddress == address {
						fmt.Println(blockTx[index])
					}
				}

			}
		}

	}

	return nil
}*/
