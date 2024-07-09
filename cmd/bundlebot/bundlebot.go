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
	chainURL = flag.String("chain", "http://10.213.21.16:8545", "chain rpc url")

	// setting: root bnb&abc boss
	rootPrivateKey = flag.String("rootpk",
		"59ba8068eb256d520179e903f43dacf6d8d57d72bd306e1bd603fdb8c8da10e8",
		"private key of root account")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	rootPk := *rootPrivateKey
	url := *chainURL

	client, err := ethclient.DialOptions(ctx, url, rpc.WithHTTPClient(utils.Client))
	if err != nil {
		log.Panicw("failed to dail chain", "err", err)
	}

	arg := &cases.BidCaseArg{
		Ctx:    ctx,
		Client: client,
		RootPk: rootPk,
	}

	txs := cases.GenerateBNBTxs(arg, cases.TransferAmountPerTx, 3)

	txBytes := make([]hexutil.Bytes, 0, len(txs))
	for _, tx := range txs {
		txByte, err := tx.MarshalBinary()
		if err != nil {
			log.Panicw("failed to marshal tx", "err", err)
		}
		txBytes = append(txBytes, txByte)
	}

	bundleArgs := types.SendBundleArgs{
		Txs: txBytes,
	}

	_, err = client.SendBundle(ctx, bundleArgs)
	if err != nil {
		log.Panicw("failed to send bundle", "err", err)
	}

	time.Sleep(10 * time.Second)

	for _, tx := range txs {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Panicw("tx not on chain", "err", err)
		}

		if receipt.Status != 1 {
			log.Panicw("tx failed", "err", err)
		}
	}

	println("bundle success")
}
