package ethereum

import(


"github.com/ethereum/go-ethereum/ethclient"
)

// EthClient ...
type EthClient struct {
	RPC *ethclient.Client
}

//New client function
// Geth does not use user:pass
func New(rpcURL string) (*EthClient, error) {
	// create new client instance
  client, err := ethclient.Dial(rpcURL)

	return &EthClient{client}, err
}
