package bitcoin

import "testing"

func TestVerify(t *testing.T) {
	tests := []struct {
		name    string
		req     VerifyRequest
		want    bool
		wantErr error
	}{
		{
			name: "",
			req: VerifyRequest{
				Address:   "bc1qzzh67ch0ucgu5cg6qhvjfwfjly99x9k8ex5slp",
				Message:   "Bitcoin is better than gold",
				Signature: "IBxc7vFmwYGGFmJjp5Hr1W6G49ednnYYkklAy4SuaFpGYE9gO7rjnrgCO/r+i2HFSKRl3TVpu87UPZPSUJp2ZjM=",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Verify(tt.req)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
			}

			if err != nil {
				t.Fatalf("function Verify returned error: %v", err)
			}

			if got != tt.want {
				t.Fatalf("expected %t, got %t", tt.want, got)
			}
		})
	}
}
