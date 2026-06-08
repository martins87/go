package bitcoin

import (
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
)

func mustPubKey(t *testing.T, pubKeyHex string) *btcec.PublicKey {
	t.Helper()

	b, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		t.Fatalf("error decoding pubKey: %v", err)
	}

	pubKey, err := btcec.ParsePubKey(b)
	if err != nil {
		t.Fatalf("error parsing pubKey: %v", err)
	}

	return pubKey
}

func TestDerivedAddresses(t *testing.T) {
	tests := []struct {
		name      string
		pubKeyHex string
		want      string
	}{
		{
			name:      "address 1",
			pubKeyHex: "02c7fb3c334a2cdf1605b3bb1a5a1c2059d335953e6b272e619b67f9e23d633906",
			want:      "bc1qvexgkvp45utderfycn4fx8hspdred9gevu26p7",
		},
		{
			name:      "address 2",
			pubKeyHex: "032528c9eabe5479a9bfe83bed64d0abf98917fa81392e3b55c2b92216c5cc096e",
			want:      "bc1qvje8muuzajuwwpqks2jnwvkk0xe5mwz72a7jqa",
		},
		{
			name:      "address 3",
			pubKeyHex: "020acd42474b671124f034797009774427ee215a96e81d330b6dd44c64bd0c816f",
			want:      "bc1qeyejth2rfl0eez9kprzkne0937nzxmyt7yz558",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pubKey := mustPubKey(t, tt.pubKeyHex)

			got, err := NativeSegwitAddress(pubKey)
			if err != nil {
				t.Fatalf("error getting segwit address: %v", err)
			}

			if got != tt.want {
				t.Fatalf("got address %q, expected %q", got, tt.want)
			}
		})
	}
}
