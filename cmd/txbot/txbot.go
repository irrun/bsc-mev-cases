package main

import (
	"context"
	"flag"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bnb-chain/bsc-mev-cases/abc"
	"github.com/bnb-chain/bsc-mev-cases/cases"
	"github.com/bnb-chain/bsc-mev-cases/log"
)

var (
	chainURL = flag.String("chain", "http://127.0.0.1:8545", "chain rpc url")

	// setting: root bnb&abc boss
	rootPrivateKey = flag.String("rootpk",
		"59ba8068eb256d520179e903f43dacf6d8d57d72bd306e1bd603fdb8c8da10e8",
		"private key of root account")
	builderPrivateKey = flag.String("builderpk",
		"adcc2278f67e14f7578a711ce93139fc1ad4033faa993fd8524d3ecbe00a365a",
		"private key of builder account")
	bobPrivateKey = flag.String("bobpk",
		"23ca29fc7e75f2a303428ee2d5526476279cabbf15c9749d1fdb080f6287e06f",
		"private key of bob account")

	abcAddress = flag.String("abc", "0xC806e70a62eaBC56E3Ee0c2669c2FF14452A9B3d", "abc contract address")

	ctx = context.Background()

	client, err = ethclient.Dial(*chainURL)
)

func main() {
	flag.Parse()

	if err != nil {
		log.Panic(err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Panic(err)
	}

	abcSol, err := abc.NewAbc(common.HexToAddress(*abcAddress), client)
	if err != nil {
		log.Panic(err)
	}

	root := cases.NewAccount(*rootPrivateKey, abcSol)
	bob := cases.NewAccount(*bobPrivateKey, abcSol)

	for {
		transferABC(root, bob, chainID, client)

		time.Sleep(100 * time.Millisecond)
	}

}

func transferABC(root, bob *cases.Account, chainID *big.Int, client *ethclient.Client) {
	tx, err := bob.TransferABC(bob.Nonce, root.Address, chainID, big.NewInt(1e15))
	if err != nil {
		log.Panicw("failed to create ABC transfer tx", "err", err)
	}

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		log.Errorw("failed to send ABC transfer tx", "err", err)
	} else {
		bob.Nonce++
	}
}
