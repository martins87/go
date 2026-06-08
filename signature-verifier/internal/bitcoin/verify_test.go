package bitcoin

import "testing"

func TestVerify(t *testing.T) {
	req := VerifyRequest{
		Address:   "bc1qzzh67ch0ucgu5cg6qhvjfwfjly99x9k8ex5slp",
		Message:   "Bitcoin is better than gold",
		Signature: "IBxc7vFmwYGGFmJjp5Hr1W6G49ednnYYkklAy4SuaFpGYE9gO7rjnrgCO/r+i2HFSKRl3TVpu87UPZPSUJp2ZjM=",
	}

	valid, err := Verify(req)
	if err != nil {
		t.Fatal(err)
	}

	if !valid {
		t.Fatalf("expected true, got %t", valid)
	}
}
