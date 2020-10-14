package keeper

import (
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

func BtcClient(rpc, user, pass string) *rpcclient.Client {
	// create new client instance
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         rpc,
		User:         user,
		Pass:         pass,
	}, nil)
	if err != nil {
		log.Fatalf("error creating new btc rpc client: %v", err)
	}
	return client
}
