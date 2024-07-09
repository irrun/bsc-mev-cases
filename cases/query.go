package cases

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

var queryCases = map[string]BidCaseFn{
	"MevRunning":    MevRunning,
	"BestBidGasFee": BestBidGasFee,
	"MevParams":     MevParams,
}

var fullNode *ethclient.Client

func init() {
	var err error
	fullNode, err = ethclient.Dial("http://10.213.21.16:8545")
	if err != nil {
		panic(err)
	}
}

func RunQueryCases(arg *BidCaseArg) {
	for n, c := range queryCases {
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

func MevRunning(arg *BidCaseArg) error {
	_, err := arg.Client.MevRunning(arg.Ctx)
	if err != nil {
		return err
	}

	return nil
}

func BestBidGasFee(arg *BidCaseArg) error {
	number, err := fullNode.BlockNumber(arg.Ctx)
	if err != nil {
		return err
	}

	block, err := fullNode.BlockByNumber(arg.Ctx, big.NewInt(int64(number)))
	if err != nil {
		return err
	}

	_, err = arg.Client.BestBidGasFee(arg.Ctx, block.Hash())
	if err != nil {
		return err
	}

	return nil
}

func MevParams(arg *BidCaseArg) error {
	_, err := arg.Client.MevParams(arg.Ctx)
	if err != nil {
		return err
	}

	return nil
}
