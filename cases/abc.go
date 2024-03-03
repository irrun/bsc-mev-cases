package cases

import (
	"errors"
	"math/big"
	"time"

	"github.com/bnb-chain/bsc-mev-cases/log"
)

var abcCases = map[string]BidCaseFn{
	"ValidBid_NilPayBidTx_ABC1":   ValidBid_NilPayBidTx_ABC1,
	"ValidBid_NilPayBidTx_ABC200": ValidBid_NilPayBidTx_ABC200,
}

// ValidBid_NilPayBidTx_ABC1
// gasFee = 21000 * 200 * 0.0000001 BNB = 0.42 BNB
func ValidBid_NilPayBidTx_ABC1(arg *BidCaseArg) error {
	txs := generateABCTxs(arg, big.NewInt(1e16), 1)
	gasUsed := ABCGasUsed * 1
	gasFee := big.NewInt(gasUsed * DefaultABCGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)

	_, err := arg.Client.SendBid(arg.Ctx, *bidArgs)
	if err != nil {
		log.Errorw("not expect error", "err", err)
		return errors.New("not expect error")
	}

	time.Sleep(5 * time.Second)

	for _, tx := range txs {
		receipt, err := arg.Client.TransactionReceipt(arg.Ctx, tx.Hash())
		if err != nil {
			log.Errorw("Client.TransactionReceipt", "err", err)
			return err
		}

		if receipt.Status != 1 {
			err = errors.New("tx failed")
			return err
		}
	}

	return nil

}

// ValidBid_NilPayBidTx_ABC200
// gasFee = 21000 * 200 * 0.0000001 BNB = 0.42 BNB
func ValidBid_NilPayBidTx_ABC200(arg *BidCaseArg) error {
	txs := generateABCTxs(arg, big.NewInt(1e16), 200)
	gasUsed := ABCGasUsed * 200
	gasFee := big.NewInt(gasUsed * DefaultABCGasPrice.Int64())
	bidArgs := generateValidBid(arg, txs, gasUsed, gasFee, false, nil)

	retry, err := assertTxSucceed(arg.Ctx, arg.Client, bidArgs, txs)
	for retry {
		bidArgs = generateValidBid(arg, txs, gasUsed, gasFee, false, nil)
		retry, err = assertTxSucceed(arg.Ctx, arg.Client, bidArgs, txs)
	}
	return err

}
