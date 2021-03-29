package tellor

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"transparence/pkg/ethereum/erc20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const CONTRACT_ADDRESS_TELLOR_PLAYGROUND_RINKEBY = "0x20374E579832859f180536A69093A126Db1c8aE9"
const INFURA_KEY_RINKEBY = ""
const IP_API_INFURA_RINKEBY = "https://rinkeby.infura.io/v3/" + INFURA_KEY_RINKEBY

const CONTRACT_ADDRESS_USING_TELLOR_MAINNET = "0xB3b7C09e1501FE212b58eEE9915DA625706eea95"
const INFURA_KEY_MAINNET = ""
const IP_API_INFURA_MAINNET = "https://mainnet.infura.io/v3/" + INFURA_KEY_MAINNET

const ETH_PRIVATE_KEY = ""

func UpdateTellorPlaygroundValue(value *big.Int, requestId *big.Int) error {
	client, err := ethclient.Dial(IP_API_INFURA_RINKEBY)
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(CONTRACT_ADDRESS_TELLOR_PLAYGROUND_RINKEBY)
	playgroundInstance, err := NewTellorPlayground(contractAddress, client)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(ETH_PRIVATE_KEY)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("casting public key to ECDSA : %w", err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := playgroundInstance.SubmitValue(auth, requestId, value)
	if err != nil {
		return err
	}
	fmt.Printf(" Transaction to Tellor's contract sent, txhash = %s\n", tx.Hash().Hex())
	return nil
}

func GetTellorPlaygroundValue(requestId *big.Int) error {
	client := erc20.NewClientConnection(IP_API_INFURA_RINKEBY)
	contractAddress := common.HexToAddress(CONTRACT_ADDRESS_TELLOR_PLAYGROUND_RINKEBY)
	playgroundInstance, err := NewTellorPlayground(contractAddress, client)
	if err != nil {
		return err
	}

	lastIndex, err := playgroundInstance.GetNewValueCountbyRequestId(&bind.CallOpts{}, requestId)
	if err != nil {
		return err
	}

	lastIndex.Sub(lastIndex, big.NewInt(1))
	timestamp, err := playgroundInstance.GetTimestampbyRequestIDandIndex(&bind.CallOpts{}, requestId, lastIndex)
	if err != nil {
		return err
	}

	result, err := playgroundInstance.RetrieveData(&bind.CallOpts{}, requestId, timestamp)
	if err != nil {
		return err
	}

	fmt.Println(" Data stored in Tellor's contract = ", result)
	return nil
}

func GetTellorValue(requestId *big.Int) (*big.Int, error) {
	client := erc20.NewClientConnection(IP_API_INFURA_MAINNET)
	contractAddress := common.HexToAddress(CONTRACT_ADDRESS_USING_TELLOR_MAINNET)
	usingTellorInstance, err := NewUsingTellor(contractAddress, client)
	result, err := usingTellorInstance.GetCurrentValue(&bind.CallOpts{}, requestId)
	if err != nil {
		return nil, err
	}
	return result.Value, nil
}
