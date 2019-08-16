package validator

import (
	"testing"
)

func TestCheckBech32AddressValidate_Success(t *testing.T) {
	type bech32AddressTest struct {
		given string
		exp   bool
	}

	bech32AddressTests := []bech32AddressTest{
		// Mainnet P2WPKH
		{
			"bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
			true,
		},
		// Testnet P2WPKH
		{
			"tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
			true,
		},
		// Mainnet P2WSH
		{
			"bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3",
			true,
		},
		// Testnet P2WSH
		{
			"tb1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q0sl5k7",
			true,
		},
	}

	for _, tt := range bech32AddressTests {
		t.Run("", func(t *testing.T) {
			got := CheckBech32AddressValidate(tt.given)
			if got != tt.exp {
				t.Errorf("CheckBech32AddressValidate(%q) returned %t", tt.given, got)
				return
			}
		})
	}
}

func TestCheckBech32AddressValidate_Fail(t *testing.T) {
	type bech32AddressTest struct {
		given string
		exp   bool
	}

	bech32AddressTests := []bech32AddressTest{
		// nil
		{
			"",
			false,
		},
		// base58 address
		{
			"175tWpb8K1S7NmH4Zx6rewF9WQrcZv245W",
			false,
		},
		// invalid checksum
		{
			"bc16qejxtdg4y5r3zarvary0c5xw7kv8xxxxx",
			false,
		},
	}

	for _, tt := range bech32AddressTests {
		t.Run("", func(t *testing.T) {
			got := CheckBech32AddressValidate(tt.given)
			if got != tt.exp {
				t.Errorf("CheckBech32AddressValidate(%q) returned %t", tt.given, got)
				return
			}
		})
	}
}

func TestCheckBtcBase58Address_Success(t *testing.T) {
	type base58AddressTest struct {
		given string
		exp   bool
	}

	base58AddressTests := []base58AddressTest{
		// Mainnet P2PKH
		{
			"17VZNX1SN5NtKa8UQFxwQbFeFc3iqRYhem",
			true,
		},
		// Mainnet P2PKH
		{
			"3EktnHQD7RiAE6uzMj2ZifT9YgRrkSgzQX",
			true,
		},
		// Testnet P2PKH
		{
			"mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
			true,
		},
		// Testnet P2PKH
		{
			"nipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
			true,
		},

		// Testnet P2PKH
		{
			"2MzQwSSnBHWHqSAqtTVQ6v47XtaisrJa1Vc",
			true,
		},
	}

	for _, tt := range base58AddressTests {
		t.Run("", func(t *testing.T) {
			got := CheckBtcBase58Address(tt.given)
			if got != tt.exp {
				t.Errorf("CheckBtcBase58Address(%q) returned %t", tt.given, got)
				return
			}
		})
	}
}

func TestCheckBtcBase58Address_Fail(t *testing.T) {
	type base58AddressTest struct {
		given string
		exp   bool
	}

	base58AddressTests := []base58AddressTest{
		// nil
		{
			"",
			false,
		},
		// invalid prefix
		{
			"47VZNX1SN5NtKa8UQFxwQbFeFc3iqRYhem",
			false,
		},
		// invalid char
		{
			"10VZNX1SN5NtKa8UQFxwQbFeFc3iqRYhem",
			false,
		},
	}

	for _, tt := range base58AddressTests {
		t.Run("", func(t *testing.T) {
			got := CheckBtcBase58Address(tt.given)
			if got != tt.exp {
				t.Errorf("CheckBtcBase58Address(%q) returned %t", tt.given, got)
				return
			}
		})
	}
}
