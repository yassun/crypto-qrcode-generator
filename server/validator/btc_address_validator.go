package validator

// Bech32
import (
	"github.com/btcsuite/btcutil/bech32"
	base58 "github.com/itchyny/base58-go"
)

// Bech32MainHrp mainnnet Leading symbol
const Bech32MainHrp = "bc1"

// Bech32TestHrp testnet Leading symbol
const Bech32TestHrp = "tb1"

// CheckBech32AddressValidate validate beh32 address
func CheckBech32AddressValidate(address string) bool {
	if len(address) < 3 {
		return false
	}
	if address[:3] == Bech32MainHrp ||
		address[:3] == Bech32TestHrp {
		if _, _, err := bech32.Decode(address); err == nil {
			return true
		}
	}
	return false
}

// P2PKHMainPreffix mainnet Leading symbol
const P2PKHMainPreffix = "1"

// P2SHMainPreffix mainnet Leading symbol
const P2SHMainPreffix = "3"

// P2PKHTestPreffix1 testnet Leading symbol
const P2PKHTestPreffix1 = "m"

// P2PKHTestPreffix2 testnet Leading symbol
const P2PKHTestPreffix2 = "n"

// P2SHTestPreffix testnet Leading symbol
const P2SHTestPreffix = "2"

// CheckBtcBase58Address validate beh32 address
func CheckBtcBase58Address(address string) bool {
	if len(address) < 1 {
		return false
	}

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
