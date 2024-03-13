package cases

import (
	"time"
)

var (
	stableCases = map[string]BidCaseFn{
		"ValidBid_NilPayBidTx_500":     ValidBid_NilPayBidTx_500,
		"InvalidBid_OldBlockNumber_20": InvalidBid_OldBlockNumber_20,
		"InvalidBid_IllegalTxs_20":     InvalidBid_IllegalTxs_20,
		"InvalidBid_GasExceed_10000":   InvalidBid_GasExceed_10000,
	}
)

// RunStableCases runs stable cases for 8h
func RunStableCases(arg *BidCaseArg) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	counter := 0

	for {
		select {
		case <-ticker.C:
			runStableCases(arg)
			counter++
			if counter > 57600 {
				println("stable test done")
				return
			}
		}
	}
}

func runStableCases(arg *BidCaseArg) {
	for n, c := range stableCases {
		waitForInTurn(arg)
		err := c(arg)
		if err != nil {
			println("stable case failed, ", "case ", n, " err ", err.Error())
		} else {
			println("stable case succeed, ", "case ", n)
		}
	}
}
