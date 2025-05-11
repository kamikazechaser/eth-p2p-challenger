package util

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func DumpKeyFromHex(keyHex string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
