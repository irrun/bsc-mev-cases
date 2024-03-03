package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	jsoniter "github.com/json-iterator/go"
)

var defaultClient = &http.Client{Timeout: 5 * time.Second}

var edpoint = "http://127.0.0.1:8546"
var chainId = big.NewInt(714)

var account, _ = fromHexKey("59ba8068eb256d520179e903f43dacf6d8d57d72bd306e1bd603fdb8c8da10e8")
var toAddr = common.HexToAddress("0x04d63aBCd2b9b1baa327f2Dda0f873F197ccd186")

func main() {
	ctx := context.Background()
	c, _ := ethclient.Dial(edpoint)

	nonce, err := c.PendingNonceAt(ctx, account.addr)
	if err != nil {
		fmt.Println(err)
	}

	txs := make([]hexutil.Bytes, 0)
	for i := 0; i < 5; i++ {
		tx, err := geTx(c, account, toAddr, big.NewInt(1), nonce)
		if err != nil {
			fmt.Println(err)
			continue
		}

		bidtx, err := tx.MarshalBinary()
		if err != nil {
			fmt.Println(err)
			continue
		}

		txs = append(txs, bidtx)
		nonce++
	}

	blockNumber, err := c.BlockNumber(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	block, err := c.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		fmt.Println(err)
		return
	}

	bundleArgs := &SendBundleArgs{
		Txs:            txs,
		MaxBlockNumber: rpc.BlockNumber(block.NumberU64() + 10),
	}

	b, _ := json.Marshal(bundleArgs)

	msg := &jsonrpcMessage{
		Version: "2.0",
		ID:      json.RawMessage("1"),
		Method:  "eth_sendBundle",
		Params:  []json.RawMessage{b},
	}

	b, _ = json.Marshal(msg)

	req, err := http.NewRequest(http.MethodPost, edpoint, bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := defaultClient.Do(req)
	if err != nil {
		fmt.Println("❌ report packing result failed", "err", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("❌ report packing result failed", "method", req.Method, "code", resp.StatusCode)
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("√ report packing result success, but read resp fail")
		return
	}

	respBodyContent := jsonrpcMessage{}
	if err = jsoniter.Unmarshal(respBody, &respBodyContent); err != nil {
		fmt.Println("√ report packing result success, but unmarshal resp body fail")
		return
	}

	if respBodyContent.Error != nil {
		respError := *respBodyContent.Error
		fmt.Println("√ report packing result success, but return error message",
			"resp.Error", respError, "resp.Error.Message", respError.Message)
		return
	}

	fmt.Println("√ report packing result success", "resp", respBodyContent)

}

type ExtAcc struct {
	Key  *ecdsa.PrivateKey
	addr common.Address
}

func geTx(client *ethclient.Client, fromEO ExtAcc, toAddr common.Address, value *big.Int, nonce uint64) (
	*types.Transaction, error) {
	gasLimit := uint64(3e4)
	gasPrice := big.NewInt(params.GWei * 10)

	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, nil)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), fromEO.Key)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

func fromHexKey(hexkey string) (ExtAcc, error) {
	key, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		return ExtAcc{}, err
	}
	pubKey := key.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("publicKey is not of type *ecdsa.PublicKey")
		return ExtAcc{}, err
	}
	addr := crypto.PubkeyToAddress(*pubKeyECDSA)
	return ExtAcc{key, addr}, nil
}

// SendBundleArgs represents the arguments for a call.
type SendBundleArgs struct {
	Txs               []hexutil.Bytes `json:"txs"`
	MaxBlockNumber    rpc.BlockNumber `json:"maxBlockNumber"`
	MinTimestamp      *uint64         `json:"minTimestamp"`
	MaxTimestamp      *uint64         `json:"maxTimestamp"`
	RevertingTxHashes []common.Hash   `json:"revertingTxHashes"`
}

type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type jsonrpcMessage struct {
	Version string            `json:"jsonrpc,omitempty"`
	ID      json.RawMessage   `json:"id,omitempty"`
	Method  string            `json:"method,omitempty"`
	Params  []json.RawMessage `json:"params,omitempty"`
	Error   *jsonError        `json:"error,omitempty"`
	Result  json.RawMessage   `json:"result,omitempty"`
}
