package bitcoin

import (
	"encoding/base64"
	"fmt"
	"slices"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

type VerifyRequest struct {
	Address   string
	Message   string
	Signature string
}

func Verify(req VerifyRequest) (bool, error) {
	pubKey, err := getPubKey(req)
	if err != nil {
		return false, fmt.Errorf("error recovering public key: %w", err)
	}

	addresses, err := deriveAddresses(pubKey)
	if err != nil {
		return false, err
	}

	if slices.Contains(addresses, req.Address) {
		return true, nil
	}

	return false, nil
}

func getPubKey(req VerifyRequest) (*btcec.PublicKey, error) {
	hash, err := MessageHash(req.Message)
	if err != nil {
		return nil, err
	}

	sigBytes, err := base64.StdEncoding.DecodeString(req.Signature)
	if err != nil {
		return nil, err
	}

	pubKey, _, err := ecdsa.RecoverCompact(sigBytes, hash)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

func deriveAddresses(pubKey *btcec.PublicKey) ([]string, error) {
	var addresses []string

	legacy, err := LegacyAddress(pubKey)
	if err == nil {
		addresses = append(addresses, legacy)
	}

	segwit, err := NativeSegwitAddress(pubKey)
	if err == nil {
		addresses = append(addresses, segwit)
	}

	return addresses, nil
}
