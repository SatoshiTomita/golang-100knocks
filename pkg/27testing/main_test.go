package main

import (
	"testing"
)

// テストの実行
// go test ./pkg/27testing/

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 1, 0},
		{0, 0, 0},
		{2, 1, 1},
	}

	for _, test := range tests {
		result := Sub(test.a, test.b)
		if result != test.expected {
			t.Errorf("Sub(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
		}
	}
}
