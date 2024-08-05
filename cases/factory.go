package cases

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bnb-chain/bsc-mev-cases/abc"
	"github.com/bnb-chain/bsc-mev-cases/log"
)

var (
	BNBGasUsed          = int64(21000)
	ABCGasUsed          = int64(21620)
	PayBidGasUsed       = int64(25000)
	BuilderFee          = big.NewInt(1e14 * 5)
	TransferAmountPerTx = big.NewInt(0)
)

type BidFactory struct {
	ctx context.Context

	root *Account
	bob  *Account

	chainID *big.Int
	client  *ethclient.Client
}

func NewBidFactory(
	ctx context.Context,
	client *ethclient.Client,
	rootPk, RootPk string,
	abcSol *abc.Abc,
) *BidFactory {
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Errorw("Client.ChainID", "err", err)
	}

	root := NewAccount(rootPk, abcSol)
	bob := NewAccount(RootPk, abcSol)

	return &BidFactory{
		ctx:     ctx,
		root:    root,
		bob:     bob,
		chainID: chainID,
		client:  client,
	}
}

func (b *BidFactory) Accounts() []*Account {
	return []*Account{b.root, b.bob}
}

func (b *BidFactory) Root() *Account {
	return b.root
}

func (b *BidFactory) Bob() *Account {
	return b.bob
}

func (b *BidFactory) BundleBNB(amount *big.Int, bundleSize int) ([]*types.Transaction, error) {
	from := b.root
	to := b.bob

	txs := make([]*types.Transaction, 0)
	for i := 0; i < bundleSize; i++ {
		tx, err := from.TransferBNB(from.Nonce, to.Address, b.chainID, amount)
		if err != nil {
			log.Errorw("failed to create BNB transfer tx", "err", err)
			return nil, err
		}

		txs = append(txs, tx)
		from.Nonce++
	}

	return txs, nil
}

func (b *BidFactory) BundleBNBWithHighGas(amount *big.Int, bundleSize int) ([]*types.Transaction, error) {
	from := b.root
	to := b.bob

	txs := make([]*types.Transaction, 0)
	for i := 0; i < bundleSize; i++ {
		tx, err := from.TransferBNBWithHighGas(from.Nonce, to.Address, b.chainID, amount)
		if err != nil {
			log.Errorw("failed to create BNB transfer tx", "err", err)
			return nil, err
		}

		txs = append(txs, tx)
		from.Nonce++
	}

	return txs, nil
}

func (b *BidFactory) BundleBNBNoSign(from, to *Account, amount *big.Int, bundleSize int) ([]*types.Transaction, error) {
	txs := make([]*types.Transaction, 0)
	for i := 0; i < bundleSize; i++ {
		tx, err := from.TransferBNBNoSign(from.Nonce, to.Address, b.chainID, amount)
		if err != nil {
			log.Errorw("failed to create BNB transfer tx", "err", err)
			return nil, err
		}

		txs = append(txs, tx)
		from.Nonce++
	}

	return txs, nil
}

func (b *BidFactory) BundleABC(amount *big.Int, bundleSize int) (types.Transactions, error) {
	from := b.root
	to := b.bob

	txs := make([]*types.Transaction, 0)
	for i := 0; i < bundleSize; i++ {
		tx, err := from.TransferABC(from.Nonce, to.Address, b.chainID, amount)
		if err != nil {
			log.Errorw("failed to create ABC transfer tx", "err", err)
			return nil, err
		}

		txs = append(txs, tx)
		from.Nonce++
	}

	return txs, nil
}

func GenerateBNBTxs(arg *BidCaseArg, amountPerTx *big.Int, txcount int) types.Transactions {
	bundleFactory := NewBidFactory(arg.Ctx, arg.Client, arg.RootPk, arg.RootPk, arg.Abc)

	txs := make([]*types.Transaction, 0)

	bundle, err := bundleFactory.BundleBNB(amountPerTx, txcount)
	if err != nil {
		log.Errorw("bundleFactory.BundleBNB", "err", err)
	}
	txs = append(txs, bundle...)

	return txs
}

func GenerateBNBTxsWithHighGas(arg *BidCaseArg, amountPerTx *big.Int, txcount int) types.Transactions {
	bundleFactory := NewBidFactory(arg.Ctx, arg.Client, arg.RootPk, arg.RootPk, arg.Abc)

	txs := make([]*types.Transaction, 0)

	bundle, err := bundleFactory.BundleBNBWithHighGas(amountPerTx, txcount)
	if err != nil {
		log.Errorw("bundleFactory.BundleBNB", "err", err)
	}
	txs = append(txs, bundle...)

	return txs
}

func generateABCTxs(arg *BidCaseArg, amountPerTx *big.Int, txcount int) types.Transactions {
	bundleFactory := NewBidFactory(arg.Ctx, arg.Client, arg.RootPk, arg.RootPk, arg.Abc)

	txs := make([]*types.Transaction, 0)

	bundle, err := bundleFactory.BundleABC(amountPerTx, txcount)
	if err != nil {
		log.Errorw("bundleFactory.BundleBNB", "err", err)
	}
	txs = append(txs, bundle...)

	return txs
}

func generateValidBid(arg *BidCaseArg, txs []*types.Transaction, gasUsed int64, gasFee *big.Int, payBuilder bool, builderFee *big.Int) *types.BidArgs {
	txBytes := make([]hexutil.Bytes, 0)
	for _, tx := range txs {
		txByte, err := tx.MarshalBinary()
		if err != nil {
			log.Panicw("tx.MarshalBinary", "err", err)
		}
		txBytes = append(txBytes, txByte)
	}

	chainID, err := fullNode.ChainID(arg.Ctx)
	for err != nil {
		chainID, err = fullNode.ChainID(arg.Ctx)
	}

	blockNumber, err := fullNode.BlockNumber(arg.Ctx)
	if err != nil {
		log.Panicw("Client.BlockNumber", "err", err)
	}

	block, err := fullNode.BlockByNumber(arg.Ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Panicw("Client.BlockByNumber", "err", err)
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
