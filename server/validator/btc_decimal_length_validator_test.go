package validator

import (
	"testing"
)

func TestBtcDecimalLengthValidate(t *testing.T) {
	type btcDecimalLengthTest struct {
		given string
		exp   bool
	}

	btcDecimalLengthTests := []btcDecimalLengthTest{
		{
			"0",
			true,
		},
		{
			"0.0",
			true,
		},
		{
			"0.00000001",
			true,
		},
		{
			"0.000000001",
			false,
		},
		{
			"10000000.0",
			true,
		},
	}

	for _, tt := range btcDecimalLengthTests {
		t.Run("", func(t *testing.T) {
			got := BtcDecimalLengthValidate(tt.given)
			if got != tt.exp {
				t.Errorf("BtcDecimalLengthValidate(%q) returned %t", tt.given, got)
				return
			}
		})
	}

}
