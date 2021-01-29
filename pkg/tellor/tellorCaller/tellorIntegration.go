package tellorCaller

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "transparence/pkg/tellor/tellorPlayground"
)


const CONTRACT_ADDRESS_USING_TELLOR_RINKEBY = "0xFC1Cc83775cf42433C3E06410Bc5367E6676e56b"
const CONTRACT_ADDRESS_TELLOR_PLAYGROUND_RINKEBY = "0x20374E579832859f180536A69093A126Db1c8aE9"
const INFURA_KEY_RINKEBY=""
const IP_API_INFURA_RINKEBY = "https://rinkeby.infura.io/v3/" + INFURA_KEY_RINKEBY

const ETH_PRIVATE_KEY=""


func UpdateTellorValue(value *big.Int,requestId *big.Int) {
    client, err := ethclient.Dial(IP_API_INFURA_RINKEBY)
    if err != nil {
        log.Fatal(err)
    }

    contractAddress := common.HexToAddress(CONTRACT_ADDRESS_TELLOR_PLAYGROUND_RINKEBY)
    playgroundInstance, err := tellorPlayground.NewTellorPlayground(contractAddress, client)
    if err != nil {
        log.Fatal(err)
    }

    privateKey, err := crypto.HexToECDSA(ETH_PRIVATE_KEY)
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice    

    tx, err := playgroundInstance.SubmitValue(auth, requestId, value)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf(" Transaction to Tellor's contract sent, txhash = %s", tx.Hash().Hex())
    fmt.Print("\n")
}

func GetTellorValue(requestId *big.Int) {
    client, err := ethclient.Dial(IP_API_INFURA_RINKEBY)
    if err != nil {
        log.Fatal(err)
    }

    contractAddress := common.HexToAddress(CONTRACT_ADDRESS_TELLOR_PLAYGROUND_RINKEBY)
    playgroundInstance, err := tellorPlayground.NewTellorPlayground(contractAddress, client)
    if err != nil {
        log.Fatal(err)
    }

    lastIndex, err := playgroundInstance.GetNewValueCountbyRequestId(&bind.CallOpts{}, requestId)
    if err != nil {
        log.Fatal(err)
    }

    lastIndex.Sub(lastIndex,big.NewInt(1))
    timestamp, err := playgroundInstance.GetTimestampbyRequestIDandIndex(&bind.CallOpts{}, requestId, lastIndex)
    if err != nil {
        log.Fatal(err)
    }

    result, err := playgroundInstance.RetrieveData(&bind.CallOpts{}, requestId, timestamp)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(" Value stored in Tellor's contract = ",result)
}