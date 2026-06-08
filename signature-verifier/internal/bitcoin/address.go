package bitcoin

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func NativeSegwitAddress(pubKey *btcec.PublicKey) (string, error) {
	pubKeyBytes := pubKey.SerializeCompressed()
	fmt.Printf("[NativeSegwitAddress] public key: %s\n", hex.EncodeToString(pubKeyBytes))

	hash160 := btcutil.Hash160(pubKeyBytes)

	addr, err := btcutil.NewAddressWitnessPubKeyHash(hash160, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	return addr.EncodeAddress(), nil
}

func LegacyAddress(pubKey *btcec.PublicKey) (string, error) {
	pubKeyBytes := pubKey.SerializeCompressed()
	fmt.Printf("[LegacyAddress] public key: %s\n", hex.EncodeToString(pubKeyBytes))

	hash160 := btcutil.Hash160(pubKeyBytes)

	addr, err := btcutil.NewAddressPubKeyHash(hash160, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	return addr.EncodeAddress(), nil
}
