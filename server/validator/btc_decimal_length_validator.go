package validator

import "strings"

// BtcSignificantDigit is minimum unit of BTC (0.00000001btc)
const BtcSignificantDigit = 8

// BtcDecimalLengthValidate checks amount's number of digits
func BtcDecimalLengthValidate(decimal string) bool {
	di := strings.Index(decimal, ".")
	if di != -1 && len(decimal[di+1:]) > BtcSignificantDigit {
		return false
	}
	return true
}
