package bitcoin

import (
	"bytes"
	"crypto/sha256"

	"github.com/btcsuite/btcd/wire"
)

func MessageHash(message string) ([]byte, error) {
	var buf bytes.Buffer

	if err := wire.WriteVarString(&buf, 0, "Bitcoin Signed Message:\n"); err != nil {
		return nil, err
	}

	if err := wire.WriteVarString(&buf, 0, message); err != nil {
		return nil, err
	}

	firstDigest := sha256.Sum256(buf.Bytes())
	secondDigest := sha256.Sum256(firstDigest[:])

	return secondDigest[:], nil
}
