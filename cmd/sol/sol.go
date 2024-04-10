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
		"59ba8068eb256d520179e903f43dacf6d8d57d72bd306e1bd603fdb8c8da10e8",
		"private key of root account")
	builderPrivateKey = flag.String("builderpk",
		"adcc2278f67e14f7578a711ce93139fc1ad4033faa993fd8524d3ecbe00a365a",
		"private key of builder account")
	bobPrivateKey = flag.String("bobpk",
		"23ca29fc7e75f2a303428ee2d5526476279cabbf15c9749d1fdb080f6287e06f",
		"private key of bob account")

	abcAddress = flag.String("abc", "0xC806e70a62eaBC56E3Ee0c2669c2FF14452A9B3d", "abc contract address")

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
	bob := cases.NewAccount(*bobPrivateKey, abc)

	tx, err := root.TransferABC(root.Nonce, builder.Address, chainID, big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18)))
	if err != nil {
		log.Panic("failed to create ABC transfer tx", "err", err)
	}

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		log.Panic("failed to send ABC transfer tx", "err", err)
	}

	root.Nonce++
	tx, err = root.TransferBNB(root.Nonce, builder.Address, chainID, big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18)))
	if err != nil {
		log.Panicw("failed to create BNB transfer tx", "err", err)
	}

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		log.Panicw("failed to send BNB transfer tx", "err", err)
	}

	root.Nonce++
	tx, err = root.TransferBNB(root.Nonce, bob.Address, chainID, big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18)))
	if err != nil {
		log.Panicw("failed to create BNB transfer tx", "err", err)
	}

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		log.Panicw("failed to send BNB transfer tx", "err", err)
	}

	root.Nonce++
	tx, err = root.TransferABC(root.Nonce, bob.Address, chainID, big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18)))
	if err != nil {
		log.Panicw("failed to create BNB transfer tx", "err", err)
	}

	err = client.SendTransaction(ctx, tx)
	if err != nil {
		log.Panicw("failed to send BNB transfer tx", "err", err)
	}

	balance(ctx, client, abc)
}

func balance(ctx context.Context, client *ethclient.Client, abcSol *abc.Abc) {
	_, rootAddress := cases.PriKeyToAddress(*rootPrivateKey)
	_, builderAddress := cases.PriKeyToAddress(*builderPrivateKey)
	_, bobAddress := cases.PriKeyToAddress(*bobPrivateKey)

	rootBalance, err := abcSol.BalanceOf(callOpts(), rootAddress)
	if err != nil {
		log.Panic(err)
	}

	log.Infow("abc balance", "address", rootAddress, "balance", rootBalance.String())

	builderBalance, err := abcSol.BalanceOf(callOpts(), builderAddress)
	if err != nil {
		log.Panic(err)
	}

	log.Infow("abc balance", "address", builderAddress, "balance", builderBalance.String())

	bobBalance, err := abcSol.BalanceOf(callOpts(), bobAddress)
	if err != nil {
		log.Panic(err)
	}

	log.Infow("abc balance", "address", bobAddress, "balance", bobBalance.String())

	rootBNBBalance, err := client.BalanceAt(ctx, rootAddress, nil)
	if err != nil {
		log.Panic(err)
	}

	log.Infow("bnb balance", "address", rootAddress, "balance", rootBNBBalance.String())

	builderBNBBalance, err := client.BalanceAt(ctx, builderAddress, nil)
	if err != nil {
		log.Panic(err)
	}

	log.Infow("bnb balance", "address", builderAddress, "balance", builderBNBBalance.String())

	bobBNBBalance, err := client.BalanceAt(ctx, bobAddress, nil)
	if err != nil {
		log.Panic(err)
	}

	log.Infow("bnb balance", "address", bobAddress, "balance", bobBNBBalance.String())
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
