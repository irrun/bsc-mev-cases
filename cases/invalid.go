package cases

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/bnb-chain/bsc-mev-cases/log"
)

var invalidBidCases = map[string]BidCaseFn{
	"InvalidBid_OldBlockNumber_20":               InvalidBid_OldBlockNumber_20,
	"InvalidBid_FutureNumber_20":                 InvalidBid_FutureNumber_20,
	"InvalidBid_NilNumber_20":                    InvalidBid_NilNumber_20,
	"InvalidBid_InvalidParentHash_20":            InvalidBid_InvalidParentHash_20,
	"InvalidBid_EmptyTxs_20":                     InvalidBid_EmptyTxs_20,
	"InvalidBid_IllegalTxs_3":                    InvalidBid_IllegalTxs_3,
	"InvalidBid_IllegalTxs_20":                   InvalidBid_IllegalTxs_20,
	"InvalidBid_FailedTx_20":                     InvalidBid_FailedTx_20,
	"InvalidBid_GasExceed_10000":                 InvalidBid_GasExceed_10000,
	"InvalidBid_NilGasUsed_20":                   InvalidBid_NilGasUsed_20,
	"InvalidBid_LessGasFee_20":                   InvalidBid_LessGasFee_20,
	"InvalidBid_MoreGasFee_20":                   InvalidBid_MoreGasFee_20,
	"InvalidBid_NilGasFee_20":                    InvalidBid_NilGasFee_20,
	"InvalidBid_EmptyGasFee_20":                  InvalidBid_EmptyGasFee_20,
	"InvalidBid_InvalidSignature_20":             InvalidBid_InvalidSignature_20,
	"InvalidBid_ExpensiveBuilderFee_20":          InvalidBid_ExpensiveBuilderFee_20,
	"InvalidBid_NilPayBidTx_NonNilPayGasUsed_20": InvalidBid_NilPayBidTx_NonNilPayGasUsed_20,
	"InvalidBid_NonNilPayBidTx_NilPayGasUsed_20": InvalidBid_NonNilPayBidTx_NilPayGasUsed_20,

	//"InvalidBid_LessGasUsed_20":                  InvalidBid_LessGasUsed_20,
	//"InvalidBid_MoreGasUsed_20":                  InvalidBid_MoreGasUsed_20,
}

func RunInvalidCases(arg *BidCaseArg) {
	for n, c := range invalidBidCases {
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

// InvalidBid_OldBlockNumber_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_OldBlockNumber_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
	bidArgs.RawBid.BlockNumber -= 10

	retry, err := assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
		retry, err = assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	}
	return err
}

// InvalidBid_FutureNumber_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_FutureNumber_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
	bidArgs.RawBid.BlockNumber += 100

	retry, err := assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
		retry, err = assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	}
	return err
}

// InvalidBid_NilNumber_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_NilNumber_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
	bidArgs.RawBid.BlockNumber = 0

	retry, err := assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
		retry, err = assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	}
	return err
}

// InvalidBid_InvalidParentHash_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_InvalidParentHash_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
	bidArgs.RawBid.ParentHash = common.Hash{}

	retry, err := assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
		retry, err = assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	}
	return err
}

// InvalidBid_EmptyTxs_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_EmptyTxs_20(arg *BidCaseArg) error {
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, nil, gasUsed, gasFee, false, nil)

	retry, err := assertNoError(arg.Ctx, arg.Client, bidArgs, nil)
	for retry {
		bidArgs = generateValidBid(arg, nil, gasUsed, gasFee, false, nil)
		retry, err = assertNoError(arg.Ctx, arg.Client, bidArgs, nil)
	}
	return err
	// TODO check no err log
}

// InvalidBid_IllegalTxs_3
// gasFee = 21000 * 3 * 0.0000001 BNB = 0.0063 BNB
func InvalidBid_IllegalTxs_3(arg *BidCaseArg) error {
	txs := generateBNBTxsNoSign(arg, TransferAmountPerTx, 3)
	gasUsed := BNBGasUsed * 3
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)

	retry, err := assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
		retry, err = assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	}
	return err
}

// InvalidBid_IllegalTxs_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_IllegalTxs_20(arg *BidCaseArg) error {
	txs := generateBNBTxsNoSign(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)

	retry, err := assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
		retry, err = assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	}
	return err
}

// InvalidBid_FailedTx_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_FailedTx_20(arg *BidCaseArg) error {
	txs := generateBNBFailedTxs(arg, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)

	retry, err := assertNoError(arg.Ctx, arg.Client, bidArgs, nil)
	for retry {
		bidArgs = generateValidBid(arg, nil, gasUsed, gasFee, false, nil)
		retry, err = assertNoError(arg.Ctx, arg.Client, bidArgs, nil)
	}
	return err
	// TODO check has err log: insufficient funds
}

// InvalidBid_GasExceed_10000
// gasFee = 21000 * 10000 * 0.0000001 BNB = 0.042*500 BNB
func InvalidBid_GasExceed_10000(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 10000)
	gasUsed := BNBGasUsed * 10000
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)

	retry, err := assertNoError(arg.Ctx, arg.Client, bidArgs, nil)
	for retry {
		bidArgs = generateValidBid(arg, nil, gasUsed, gasFee, false, nil)
		retry, err = assertNoError(arg.Ctx, arg.Client, bidArgs, nil)
	}
	return err
	// TODO check has err log: gas limit reached
}

//// InvalidBid_LessGasUsed_20
//// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
//func InvalidBid_LessGasUsed_20(arg *BidCaseArg) error {
//	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
//	gasUsed := BNBGasUsed * 10
//	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
//	_ = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
//
//	// expect: Builder receive issue
//
//	return nil
//}
//
//// InvalidBid_MoreGasUsed_20
//// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
//func InvalidBid_MoreGasUsed_20(arg *BidCaseArg) error {
//	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
//	gasUsed := BNBGasUsed * 30
//	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
//	_ = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
//
//	// expect: Builder receive issue
//
//	return nil
//}

// InvalidBid_NilGasUsed_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_NilGasUsed_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := int64(0)
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)

	retry, err := assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
		retry, err = assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	}
	return err
}

// InvalidBid_LessGasFee_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_LessGasFee_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * big.NewInt(1e9).Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)

	retry, err := assertNoError(arg.Ctx, arg.Client, bidArgs, nil)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
		retry, err = assertNoError(arg.Ctx, arg.Client, bidArgs, nil)
	}
	return err
	// TODO check succeed when idle
}

// InvalidBid_MoreGasFee_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_MoreGasFee_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * big.NewInt(1e12).Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)

	retry, err := assertNoError(arg.Ctx, arg.Client, bidArgs, nil)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
		retry, err = assertNoError(arg.Ctx, arg.Client, bidArgs, nil)
	}
	return err
	// TODO check has err log: invalid reward
}

// InvalidBid_NilGasFee_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_NilGasFee_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	bidArgs := generateValidBid(arg, txs, gasUsed, nil, false, nil)

	retry, err := assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, nil, false, nil)
		retry, err = assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	}
	return err
}

// InvalidBid_EmptyGasFee_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_EmptyGasFee_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(0)
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)

	retry, err := assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
		retry, err = assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	}
	return err
}

// InvalidBid_InvalidSignature_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_InvalidSignature_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
	bidArgs.Signature = []byte("invalid signature")

	retry, err := assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
		bidArgs.Signature = []byte("invalid signature")
		retry, err = assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	}
	return err
}

// InvalidBid_ExpensiveBuilderFee_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_ExpensiveBuilderFee_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	builderFee := big.NewInt(gasUsed*DefaultBNBGasPrice.Int64() + 1)
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, true, builderFee)

	retry, err := assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, true, builderFee)
		retry, err = assertInvalidBidParam(arg.Ctx, arg.Client, bidArgs)
	}

	return err
}

// InvalidBid_NilPayBidTx_NonNilPayGasUsed_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_NilPayBidTx_NonNilPayGasUsed_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	builderFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64() / 5)
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, true, builderFee)
	bidArgs.PayBidTx = nil

	retry, err := assertInvalidPayBidTx(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, true, builderFee)
		bidArgs.PayBidTx = nil
		retry, err = assertInvalidPayBidTx(arg.Ctx, arg.Client, bidArgs)
	}

	return err
}

// InvalidBid_NonNilPayBidTx_NilPayGasUsed_20
// gasFee = 21000 * 20 * 0.0000001 BNB = 0.042 BNB
func InvalidBid_NonNilPayBidTx_NilPayGasUsed_20(arg *BidCaseArg) error {
	txs := generateBNBTxs(arg, TransferAmountPerTx, 20)
	gasUsed := BNBGasUsed * 20
	gasFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64())
	builderFee := big.NewInt(gasUsed * DefaultBNBGasPrice.Int64() / 5)
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, true, builderFee)
	bidArgs.PayBidTxGasUsed = 0

	retry, err := assertInvalidPayBidTx(arg.Ctx, arg.Client, bidArgs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, true, builderFee)
		bidArgs.PayBidTx = nil
		retry, err = assertInvalidPayBidTx(arg.Ctx, arg.Client, bidArgs)
	}

	return err
}

func generateBNBTxsNoSign(arg *BidCaseArg, amountPerTx *big.Int, txcount int) types.Transactions {
	bundleFactory := NewBidFactory(arg.Ctx, arg.Client, arg.RootPk, arg.BobPk, arg.Abc)
	root := bundleFactory.Root()
	bob := bundleFactory.Bob()

	txs := make([]*types.Transaction, 0)

	bundle, err := bundleFactory.BundleBNBNoSign(root, bob, amountPerTx, txcount)
	if err != nil {
		log.Errorw("bundleFactory.BundleBNB", "err", err)
	}
	txs = append(txs, bundle...)

	return txs
}

func assertInvalidBidParam(ctx context.Context, client *ethclient.Client, bidArgs *types.BidArgs) (
	bool, error) {
	_, err := client.SendBid(ctx, *bidArgs)
	if err == nil {
		return false, fmt.Errorf("expect error but return nil")
	}

	bidErr, ok := err.(rpc.Error)
	if !ok {
		return false, fmt.Errorf("expect jsonrpc error but not")
	}

	if bidErr.ErrorCode() == types.InvalidBidParamError {
		return false, nil
	}

	return true, bidErr
}

func assertInvalidPayBidTx(ctx context.Context, client *ethclient.Client, bidArgs *types.BidArgs) (
	bool, error) {
	_, err := client.SendBid(ctx, *bidArgs)
	if err == nil {
		return false, fmt.Errorf("expect error but return nil")
	}

	bidErr, ok := err.(rpc.Error)
	if !ok {
		return false, fmt.Errorf("expect jsonrpc error but not")
	}

	if bidErr.ErrorCode() == types.InvalidPayBidTxError {
		return false, nil
	}

	return true, bidErr
}

func assertNoError(ctx context.Context, client *ethclient.Client, bidArgs *types.BidArgs, txs types.Transactions) (
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

	return false, nil
}
