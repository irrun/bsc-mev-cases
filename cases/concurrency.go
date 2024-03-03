package cases

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bnb-chain/bsc-mev-cases/log"
	"github.com/bnb-chain/bsc-mev-cases/utils/syncutils"
)

func RunConcurrency(arg *BidCaseArg) {
	txCounts := []int{8, 32, 512}
	bidArgs := make([]*types.BidArgs, len(txCounts))
	txs := make([]types.Transactions, len(txCounts))

	chainID, err := arg.Client.ChainID(arg.Ctx)
	for err != nil {
		chainID, err = arg.Client.ChainID(arg.Ctx)
	}
	println("chainID ", chainID.String())

	retry := true

	for retry {
		blockNumber, err := fullNode.BlockNumber(arg.Ctx)
		if err != nil {
			log.Panicw("Client.BlockNumber", "err", err)
		}

		block, err := fullNode.BlockByNumber(arg.Ctx, big.NewInt(int64(blockNumber)))
		if err != nil {
			log.Panicw("Client.BlockByNumber", "err", err)
		}

		println("blockNumber ", blockNumber, " blockHash ", block.Hash().String())

		for i, c := range txCounts {
			bidArgs[i], txs[i] = geBidArgs(arg, c, chainID, block)
		}

		br := syncutils.NewBatchRunner().WithConcurrencyLimit(10)
		br.AddTasks(func() error {
			for _, b := range bidArgs {
				var er error
				_, er = arg.Client.SendBid(arg.Ctx, *b)
				if er != nil {
					return er
				}
			}

			return nil
		})
		err = br.Exec()

		if err != nil {
			println(err.Error())
			time.Sleep(500 * time.Millisecond)
		} else {
			retry = false
		}
	}

	time.Sleep(5 * time.Second)

	for _, tx := range txs[2] {
		receipt, err := arg.Client.TransactionReceipt(arg.Ctx, tx.Hash())
		if err != nil {
			println("concurrency failed ", err.Error())
			return
		}

		if receipt.Status != 1 {
			println("concurrency failed ", receipt.Status)
		}
	}

	fmt.Println("concurrency success")
}

func geBidArgs(arg *BidCaseArg, txCount int, chainID *big.Int, block *types.Block) (
	*types.BidArgs, types.Transactions) {
	txs := generateBNBTxs(arg, TransferAmountPerTx, txCount)
	gasUsed := BNBGasUsed * int64(txCount)
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := geValidBidWithBlock(arg, txs, gasUsed, gasFee, false, nil, chainID, block)

	return bidArgs, txs
}

func geValidBidWithBlock(
	arg *BidCaseArg, txs []*types.Transaction, gasUsed int64, gasFee *big.Int, payBuilder bool, builderFee *big.Int,
	chainID *big.Int, block *types.Block) *types.BidArgs {
	txBytes := make([]hexutil.Bytes, 0)
	for _, tx := range txs {
		txByte, err := tx.MarshalBinary()
		if err != nil {
			log.Panicw("tx.MarshalBinary", "err", err)
		}
		txBytes = append(txBytes, txByte)
	}

	rawBid := &types.RawBid{
		BlockNumber: block.NumberU64() + 1,
		ParentHash:  block.Hash(),
		Txs:         txBytes,
		GasUsed:     uint64(gasUsed),
		GasFee:      gasFee,
	}

	if payBuilder {
		rawBid.BuilderFee = builderFee
	}

	bidArgs := arg.Builder.SignBid(rawBid)
	if !payBuilder {
		return bidArgs
	}

	validator := arg.Validators[0]
	payBidTx := arg.Builder.PayBidTx(validator, chainID, builderFee)
	bidArgs.PayBidTx = payBidTx
	bidArgs.PayBidTxGasUsed = uint64(PayBidGasUsed)
	return bidArgs
}
