package main

import (
	"context"
	"flag"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/bnb-chain/bsc-mev-cases/cases"
	"github.com/bnb-chain/bsc-mev-cases/log"
	"github.com/bnb-chain/bsc-mev-cases/utils"
)

var (
	chainURL = flag.String("chain", "https://bsc-testnet-builder.bnbchain.org/", "chain rpc url")

	// setting: root bnb&abc boss
	// root 0x04d63aBCd2b9b1baa327f2Dda0f873F197ccd186
	// bob 0x88d2eb89e00ca61c225ef673fcbe4f8d1b3ee28f
	rootPrivateKey = flag.String("rootpk",
		"69c620c47a6c561c6f5e5b2496e477b626144e5b9c7257602fdeca670059f929",
		"private key of root account")
	bobPrivateKey = flag.String("bobpk",
		"69c620c47a6c561c6f5e5b2496e477b626144e5b9c7257602fdeca670059f929",
		"private key of bob account")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	rootPk := *rootPrivateKey
	bobPk := *bobPrivateKey
	url := *chainURL

	client, err := ethclient.DialOptions(ctx, url, rpc.WithHTTPClient(utils.Client))
	if err != nil {
		log.Panicw("failed to dail chain", "err", err)
	}

	arg := &cases.BidCaseArg{
		Ctx:    ctx,
		Client: client,
		RootPk: rootPk,
		BobPk:  bobPk,
	}

	for i := 0; i < 72; i++ {
		txs := cases.GenerateBNBTxsWithHighGas(arg, cases.TransferAmountPerTx, 1)

		txBytes := make([]hexutil.Bytes, 0, len(txs))
		for _, tx := range txs {
			txByte, err := tx.MarshalBinary()
			if err != nil {
				log.Panicw("failed to marshal tx", "err", err)
			}
			txBytes = append(txBytes, txByte)
		}

		bundleArgs := &types.SendBundleArgs{
			Txs: txBytes,
		}

		err = client.SendBundle(ctx, bundleArgs)
		if err != nil {
			log.Panicw("failed to send bundle", "err", err)
		}

		for _, tx := range txs {
			log.Infow("hash", "v", tx.Hash())
		}

		time.Sleep(1 * time.Hour)
	}
}
