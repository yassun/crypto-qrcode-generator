package validator

import "strings"

// 0.00000001btc
const BtcSignificantDigit = 8

func BtcDecimalLengthValidate(decimal string) bool {
	di := strings.Index(decimal, ".")
	if di != -1 && len(decimal[di+1:]) > BtcSignificantDigit {
		return false
	}
	return true
}
