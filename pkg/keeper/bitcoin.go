package keeper

import (
	"errors"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
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
