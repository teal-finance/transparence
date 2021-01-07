package keeper

import (
	"errors"
	"fmt"
	"time"
	"encoding/hex"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/chaincfg"

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
// No call to rpc so o need to return an error
func (c *BtcClient) ExtractVout(vouts []btcjson.Vout) []string {
	addresses := make([]string, 0)
	if vouts == nil {
		return addresses
	}
	for _, vout := range vouts {
		addresses = append(addresses, vout.ScriptPubKey.Addresses...)
	}
	return addresses
}

// ExtractVin2 extract the list of addresses involved.
// The Tx hash should be taken care of by the caller of this function
// we also need to think of the coinbase tx that is the first tx, vin & vout, in each block
func (c *BtcClient) ExtractVin2(vins []btcjson.Vin) ([]string, error) {
	addresses := make([]string, 0)
	if vins == nil {
		return addresses, nil
	}

	for _, vin := range vins {
		script, err := hex.DecodeString(vin.ScriptSig.Hex)
		if err != nil {
			fmt.Println(err)
			return addresses,err
		}
		_, addrs, _, err := txscript.ExtractPkScriptAddrs(
			script, &chaincfg.MainNetParams)
		if err != nil {
			fmt.Println(err)
			return addresses,err
		}
		for _, addr := range addrs {
			addresses = append(addresses,addr.EncodeAddress())
		}
	}
	return addresses,nil

}

// ExtractVin extract the list of addresses involved.
// The Tx hash should be taken care of by the caller of this function
// we also need to think of the coinbase tx that is the first tx, vin & vout, in each block
func (c *BtcClient) ExtractVin(vins []btcjson.Vin, TxHash string) ([]string, error) {
	addresses := make([]string, 0)
	if vins == nil {
		return addresses, nil
	}
	for _, vin := range vins {
		// the easy way to get the address in Vin is to look at Vout
		//TODO just do the conversion to PubKey instead because it messed coinbase tx

		//watchout for coinbase transactions
		//This is not working for coinbase tx and transfers from coinbase tx
		if vin.IsCoinBase() {
			/*
			hash, err := chainhash.NewHashFromStr(TxHash)
			if err != nil {
				return []string{}, err
			}
			transaction, err := c.RPC.GetRawTransactionVerbose(hash)
			if err != nil {
				return []string{}, err
			}
			addresses = append(addresses, transaction.Vout[0].ScriptPubKey.Addresses...)
			*/
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
		addresses = append(addresses, transaction.Vout[vin.Vout].ScriptPubKey.Addresses...)
	}
	return addresses, nil
}

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

		addressesVinVout, err := c.ExtractVin(transactionVerbose.Vin, transactionVerbose.Hash)
		if err != nil {
			return [][]string{}, err
		}
		addressesVinVout = append(addressesVinVout, c.ExtractVout(transactionVerbose.Vout)...)
		addresses = append(addresses, [][]string{addressesVinVout}...)
	}

	return addresses, nil
}

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
}
