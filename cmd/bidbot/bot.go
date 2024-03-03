package main

import (
	"context"
	"flag"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/bnb-chain/bsc-mev-cases/abc"
	"github.com/bnb-chain/bsc-mev-cases/cases"
	"github.com/bnb-chain/bsc-mev-cases/log"
	"github.com/bnb-chain/bsc-mev-cases/utils"
)

var (
	chainURL = flag.String("chain", "http://127.0.0.1:8545", "chain rpc url")

	// setting: root bnb boss, bob abc boss
	rootPrivateKey = flag.String("rootpk",
		"59ba8068eb256d520179e903f43dacf6d8d57d72bd306e1bd603fdb8c8da10e8",
		"private key of root account")
	bobPrivateKey = flag.String("bobpk",
		"23ca29fc7e75f2a303428ee2d5526476279cabbf15c9749d1fdb080f6287e06f",
		"private key of bob account")
	builderPrivateKey = flag.String("builderpk",
		"7b94e64fc431b0daa238d6ed8629f3747782b8bc10fb8a41619c5fb2ba55f4e3",
		"private key of builder account")

	abcAddress = flag.String("abc", "0xC806e70a62eaBC56E3Ee0c2669c2FF14452A9B3d", "abc contract address")

	validator = flag.String("validator", "0xe0239549edd90eb0e4abf5cbc9edad1a4af20d3e", "validator address")

	casetype = flag.String("casetype", "valid", "case type")
	sentry   = flag.String("sentry", "http://127.0.0.1:8080", "sentry url")
	casename = flag.String("casename", "", "case name")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	rootPk := *rootPrivateKey
	bobPk := *bobPrivateKey
	builderPk := *builderPrivateKey
	url := *chainURL
	whatcase := *casetype

	client, err := ethclient.DialOptions(ctx, url, rpc.WithHTTPClient(utils.Client))
	if err != nil {
		log.Errorw("ethclient.DialOptions", "err", err)
	}

	abcSol, err := abc.NewAbc(common.HexToAddress(*abcAddress), client)
	if err != nil {
		log.Errorw("abc.NewAbc", "err", err)
	}

	arg := &cases.BidCaseArg{
		Ctx:        ctx,
		Client:     client,
		RootPk:     rootPk,
		BobPk:      bobPk,
		Abc:        abcSol,
		Builder:    cases.NewAccount(builderPk, abcSol),
		Validators: []common.Address{common.HexToAddress(*validator)},
	}

	switch whatcase {
	case "valid":
		cases.RunValidCases(arg)
	case "invalid":
		cases.RunInvalidCases(arg)
	case "stable":
		cases.RunStableCases(arg)
	case "concurrency":
		cases.RunConcurrency(arg)
	case "single":
		cases.RunCase(arg, *casename)
	case "query":
		cases.RunQueryCases(arg)
	}
}
