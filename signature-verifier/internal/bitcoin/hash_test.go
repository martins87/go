package bitcoin

import (
	"encoding/hex"
	"testing"
)

func TestMessageHash(t *testing.T) {
	tests := []struct {
		name    string
		message string
		wantHex string
	}{
		{
			name:    "1st vector",
			message: "Bitcoin is better than gold",
			wantHex: "3294bdc711cc7d884495553656041c16244012576be4b8541af075d328f7ba1b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MessageHash(tt.message)
			if err != nil {
				t.Fatal(err)
			}

			if gotHex := hex.EncodeToString(got); gotHex != tt.wantHex {
				t.Fatalf("expected hash %s, got %s", tt.wantHex, gotHex)
			}
		})
	}

}
