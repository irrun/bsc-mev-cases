package cases

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/bnb-chain/bsc-mev-cases/abc"
	"github.com/bnb-chain/bsc-mev-cases/log"
)

var (
	DefaultGasLimit    = uint64(5000000)
	GWEI               = big.NewInt(1e9)
	DefaultBNBGasPrice = big.NewInt(1e11) // 0.0000001 BNB
	HighGasPrice       = big.NewInt(1e12) // 0.001 BNB
	DefaultABCGasPrice = big.NewInt(1e11) // 0.0000001 ABC
)

type Account struct {
	Address    common.Address
	privateKey *ecdsa.PrivateKey
	Nonce      uint64
	abc        *abc.Abc
}

func NewAccount(privateKey string, abc *abc.Abc) *Account {
	privateECDSAKey, address := PriKeyToAddress(privateKey)

	nonce, err := fullNode.PendingNonceAt(context.TODO(), address)
	if err != nil {
		log.Errorw("failed to get pending Nonce", "err", err)
	}

	return &Account{
		Address:    address,
		privateKey: privateECDSAKey,
		Nonce:      nonce,
		abc:        abc,
	}
}

func (a *Account) TransferBNB(nonce uint64, toAddress common.Address, chainID *big.Int, amount *big.Int) (*types.Transaction, error) {
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    amount,
		Gas:      DefaultGasLimit,
		GasPrice: GWEI,
		Data:     nil,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), a.privateKey)
	if err != nil {
		log.Errorw("failed to sign tx", "err", err)
		return nil, err
	}

	return signedTx, nil
}

func (a *Account) TransferBNBWithHighGas(nonce uint64, toAddress common.Address, chainID *big.Int, amount *big.Int) (*types.Transaction, error) {
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    amount,
		Gas:      DefaultGasLimit,
		GasPrice: HighGasPrice,
		Data:     nil,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), a.privateKey)
	if err != nil {
		log.Errorw("failed to sign tx", "err", err)
		return nil, err
	}

	return signedTx, nil
}

func (a *Account) TransferBNBNoSign(nonce uint64, toAddress common.Address, chainID *big.Int, amount *big.Int) (*types.Transaction, error) {
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    amount,
		Gas:      DefaultGasLimit,
		GasPrice: DefaultBNBGasPrice,
		Data:     nil,
	})

	return tx, nil
}

func (a *Account) TransferABC(nonce uint64, toAddress common.Address, chainID *big.Int, amount *big.Int) (*types.Transaction, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(a.privateKey, chainID)
	if err != nil {
		log.Errorw("failed to create transactor", "err", err)
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = DefaultGasLimit
	auth.GasPrice = DefaultABCGasPrice
	auth.NoSend = true

	return a.abc.Transfer(auth, toAddress, amount)
}

func (a *Account) SignBid(rawBid *types.RawBid) *types.BidArgs {
	data, err := rlp.EncodeToBytes(rawBid)
	if err != nil {
		log.Errorw("failed to encode raw bid", "err", err)
	}

	sig, err := crypto.Sign(crypto.Keccak256(data), a.privateKey)
	if err != nil {
		log.Errorw("failed to sign raw bid", "err", err)
	}

	bidArgs := types.BidArgs{
		RawBid:    rawBid,
		Signature: sig,
	}

	return &bidArgs
}

func (a *Account) PayBidTx(receiver common.Address, chainID *big.Int, amount *big.Int) []byte {
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    a.Nonce,
		GasPrice: big.NewInt(0),
		Gas:      25000,
		To:       &receiver,
		Value:    amount,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), a.privateKey)
	for err != nil {
		log.Errorw("failed to sign tx", "err", err)
		signedTx, err = types.SignTx(tx, types.LatestSignerForChainID(chainID), a.privateKey)
	}

	txByte, _ := signedTx.MarshalBinary()
	return txByte
}

func (a *Account) BalanceBNB(client *ethclient.Client) *big.Int {
	balance, err := fullNode.BalanceAt(context.TODO(), a.Address, nil)
	if err != nil {
		log.Errorw("Client.BalanceAt", "err", err)
	}

	return balance
}

func (a *Account) BalanceABC() *big.Int {
	balance, err := a.abc.BalanceOf(callOpts(), a.Address)
	if err != nil {
		log.Errorw("Client.BalanceAt", "err", err)
	}

	return balance
}

func PriKeyToAddress(privateKey string) (*ecdsa.PrivateKey, common.Address) {
	privateECDSAKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Errorw("failed to parse private key", "err", err)
	}

	publicKey, ok := privateECDSAKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Errorw("failed to get public key", "err", err)
	}

	selfAddress := crypto.PubkeyToAddress(*publicKey)

	return privateECDSAKey, selfAddress
}

type BidCaseArg struct {
	Ctx        context.Context
	Client     *ethclient.Client
	RootPk     string
	Abc        *abc.Abc
	Builder    *Account
	Validators []common.Address
}

type BidCaseFn func(arg *BidCaseArg) error

func callOpts() *bind.CallOpts {
	callOpts := new(bind.CallOpts)
	callOpts.Context = context.Background()
	callOpts.Pending = false
	return callOpts
}
