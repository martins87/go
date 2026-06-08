package bitcoin

import (
	"encoding/hex"
	"testing"
)

const want = "3294bdc711cc7d884495553656041c16244012576be4b8541af075d328f7ba1b"

func TestMessageHash(t *testing.T) {
	got, err := MessageHash("Bitcoin is better than gold")
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != 32 {
		t.Fatalf("expected hash length 32, got %d", len(got))
	}

	if gotHex := hex.EncodeToString(got); gotHex != want {
		t.Fatalf("expected hash %s, got %s", want, gotHex)
	}
}
