package validator

// Bech32
import (
	"github.com/btcsuite/btcutil/bech32"
	base58 "github.com/itchyny/base58-go"
	validator "gopkg.in/go-playground/validator.v9"
)

// https://en.bitcoin.it/wiki/List_of_address_prefixes
const Bech32MainHrp = "bc1"
const Bech32TestHrp = "tb1"

func CheckBech32AddressValidate(address string) bool {
	if address[:3] == Bech32MainHrp ||
		address[:3] == Bech32TestHrp {
		if _, _, err := bech32.Decode(address); err == nil {
			return true
		}
	}
	return false
}

// Base58
// https://en.bitcoin.it/wiki/List_of_address_prefixes
const P2PKHMainPreffix = "1"
const P2SHMainPreffix = "3"
const P2PKHTestPreffix1 = "m"
const P2PKHTestPreffix2 = "n"
const P2SHTestPreffix = "2"

func CheckBtcBase58Address(address string) bool {
	if address[:1] == P2PKHMainPreffix ||
		address[:1] == P2SHMainPreffix ||
		address[:1] == P2PKHTestPreffix1 ||
		address[:1] == P2PKHTestPreffix2 ||
		address[:1] == P2SHTestPreffix {
		encoding := base58.BitcoinEncoding
		if _, err := encoding.Decode([]byte(address)); err == nil {
			return true
		}
	}

	return false
}

func BtcAddressValidate(fl validator.FieldLevel) bool {
	address := fl.Field().String()
	if CheckBtcBase58Address(address) || CheckBech32AddressValidate(address) {
		return true
	}
	return false
}
