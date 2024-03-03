package utils

import (
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrivateKeyFromKeystore(secretKeyFile string, password string, outPath string) {
	keyjson, e := ioutil.ReadFile(secretKeyFile)
	if e != nil {
		panic(e)
	}

	key, e := keystore.DecryptKey(keyjson, password)
	if e != nil {
		panic(e)
	}

	e = crypto.SaveECDSA(outPath, key.PrivateKey)
	if e != nil {
		panic(e)
	}
}
