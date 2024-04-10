package cases

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/bnb-chain/bsc-mev-cases/log"
)

var (
	validBidCases = map[string]BidCaseFn{
		"ValidBid_PayBidTx_200": ValidBid_PayBidTx_200,
	}
)

func RunValidCases(arg *BidCaseArg) {
	for n, c := range validBidCases {
		waitForInTurn(arg)
		print("run case ", n)
		err := c(arg)
		if err != nil {
			print(" failed: ", err.Error())
		} else {
			print(" succeed")
		}
		println()
	}
}

func RunCase(arg *BidCaseArg, name string) {
	caseFn, err := getCaseFn(name)
	if err != nil {
		println(err.Error())
		return
	}

	waitForInTurn(arg)
	print("run case ", name)
	err = caseFn(arg)
	if err != nil {
		print(" failed: ", err.Error())
	} else {
		print(" succeed")
	}

	println()
}

func getCaseFn(name string) (BidCaseFn, error) {
	c, ok := validBidCases[name]
	if ok {
		return c, nil
	}

	c, ok = invalidBidCases[name]
	if ok {
		return c, nil
	}

	c, ok = abcCases[name]
	if ok {
		return c, nil
	}

	return nil, errors.New("case fn not found")
}

// ValidBid_PayBidTx_1
// gasFee = 21000 * 1 * 0.0000001 BNB = 0.0021 BNB
// builderFee = 0.05 BNB
func ValidBid_PayBidTx_1(arg *BidCaseArg) error {
	txs := GenerateBNBTxs(arg, TransferAmountPerTx, 1)
	gasUsed := BNBGasUsed * 1
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, true, BuilderFee)

	retry, err := assertTxSucceed(arg.Ctx, arg.Client, bidArgs, txs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, true, BuilderFee)
		retry, err = assertTxSucceed(arg.Ctx, arg.Client, bidArgs, txs)
	}

	return err
}

// ValidBid_PayBidTx_200
// gasFee = 21000 * 200 * 0.0000001 BNB = 0.42 BNB
// builderFee = 0.05 BNB
func ValidBid_PayBidTx_200(arg *BidCaseArg) error {
	txs := GenerateBNBTxs(arg, TransferAmountPerTx, 200)
	gasUsed := BNBGasUsed * 200
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, true, BuilderFee)

	retry, err := assertTxSucceed(arg.Ctx, arg.Client, bidArgs, txs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, true, BuilderFee)
		retry, err = assertTxSucceed(arg.Ctx, arg.Client, bidArgs, txs)
	}

	return err
}

func assertTxSucceed(ctx context.Context, client *ethclient.Client, bidArgs *types.BidArgs, txs types.Transactions) (
	bool, error) {
	_, err := client.SendBid(ctx, *bidArgs)
	if err != nil {
		bidErr, ok := err.(rpc.Error)
		if !ok {
			log.Infow("retry", "reason", err)
			return true, err
		}

		if bidErr.ErrorCode() == types.InvalidBidParamError ||
			bidErr.ErrorCode() == types.InvalidPayBidTxError {
			log.Infow("retry", "reason", "InvalidBidParamError or InvalidPayBidTxError")
			return true, err
		}
	}

	time.Sleep(5 * time.Second)

	for i, tx := range txs {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Errorw("receipt err", "err", err)
			return false, fmt.Errorf("receipt err, %v", err)
		}

		if receipt.Status != 1 {
			log.Errorw("tx failed", "tx", tx.Hash().Hex())
			return false, fmt.Errorf("tx at index %v failed, %v", i, err)
		}
	}

	return false, nil
}

// waitForInTurn wait for the current validator in turn
func waitForInTurn(arg *BidCaseArg) {
	bidArgs := generateValidBid(arg, nil, 0, big.NewInt(0), false, nil)

	inTurn := false

	ping := func() {
		_, err := arg.Client.SendBid(arg.Ctx, *bidArgs)
		if err != nil {
			bidErr, ok := err.(rpc.Error)
			if ok && bidErr.ErrorCode() == types.InvalidBidParamError {
				inTurn = true
			}
		}
	}

	ping()

	for inTurn != true {
		println("wait for in turn")
		time.Sleep(500 * time.Millisecond)
		ping()
	}
}
