package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct{
		name string
		a, b int
		want int
	}{
		{ "positive", 1, 2, 3 },
		{ "negative", -1, -20, -21 },
		{ "zero", 1, -1, 0 },
		{ "large numbers", 1000, 20000, 21000 },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add() = %v, want = %v", got, tt.want)
			}
		})
	}
}