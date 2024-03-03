package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bnb-chain/bsc-mev-cases/abc"
	"github.com/bnb-chain/bsc-mev-cases/cases"
	"github.com/bnb-chain/bsc-mev-cases/log"
)

var (
	chainURL = flag.String("chain", "http://127.0.0.1:8545", "chain rpc url")

	rootPrivateKey = flag.String("rootpk",
		"e61213d0dde8f6d7e73b27ea0304e4db5dc455eb28f6494b353f6cea19f065e7",
		"private key of root account")
	builderPrivateKey = flag.String("builderpk",
		"7b94e64fc431b0daa238d6ed8629f3747782b8bc10fb8a41619c5fb2ba55f4e3",
		"private key of builder account")

	abcAddress = flag.String("abc", "0xa45F543E97331643cAC26B075F2958d48Bd0E317", "abc contract address")

	option = flag.String("option", "deploy", "deploy or transfer")
)

func main() {
	defer log.Stop()

	flag.Parse()

	ctx := context.Background()

	client, err := ethclient.Dial(*chainURL)
	if err != nil {
		log.Panic(err)
	}

	abcSol, err := abc.NewAbc(common.HexToAddress(*abcAddress), client)
	if err != nil {
		log.Panic(err)
	}

	switch *option {
	case "deploy":
		deploy(ctx, client)
	case "balance":
		balance(ctx, client, abcSol)
	case "transfer":
		transfer(ctx, client, abcSol)
	}
}

func transfer(ctx context.Context, client *ethclient.Client, abc *abc.Abc) {
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Panic("Client.ChainID", "err", err)
	}

	root := cases.NewAccount(*rootPrivateKey, abc)
	builder := cases.NewAccount(*builderPrivateKey, abc)

	tx, err := root.TransferABC(root.Nonce, builder.Address, chainID, big.NewInt(1e18))
	if err != nil {
		log.Panic("failed to create ABC transfer tx", "err", err)
	}

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		log.Panic("failed to send ABC transfer tx", "err", err)
	}

	root.Nonce++
	tx, err = root.TransferBNB(root.Nonce, builder.Address, chainID, big.NewInt(1e18))
	if err != nil {
		log.Panic("failed to create ABC transfer tx", "err", err)
	}

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		log.Panic("failed to send ABC transfer tx", "err", err)
	}
}

func balance(ctx context.Context, client *ethclient.Client, abcSol *abc.Abc) {
	_, rootAddress := cases.PriKeyToAddress(*rootPrivateKey)
	_, builderAddress := cases.PriKeyToAddress(*builderPrivateKey)

	rootBalance, err := abcSol.BalanceOf(callOpts(), rootAddress)
	if err != nil {
		log.Panic(err)
	}

	log.Infow("query balance", "address", rootAddress, "balance", rootBalance.String())

	builderBalance, err := abcSol.BalanceOf(callOpts(), builderAddress)
	if err != nil {
		log.Panic(err)
	}

	log.Infow("query balance", "address", builderAddress, "balance", builderBalance.String())
}

func deploy(ctx context.Context, client *ethclient.Client) {
	rootKey, rootAddress := cases.PriKeyToAddress(*rootPrivateKey)
	auth := generateAccountAuth(ctx, client, rootKey, rootAddress)

	solAddress, _, _, err := abc.DeployAbc(auth, client)
	if err != nil {
		log.Panic(err)
	}

	log.Infow("deployed abc", "address", solAddress)
}

func generateAccountAuth(ctx context.Context,
	client *ethclient.Client,
	key *ecdsa.PrivateKey,
	address common.Address) *bind.TransactOpts {
	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		log.Panic(err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		log.Panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(10000000000)

	return auth
}

func callOpts() *bind.CallOpts {
	callOpts := new(bind.CallOpts)
	callOpts.Context = context.Background()
	callOpts.Pending = false
	return callOpts
}
