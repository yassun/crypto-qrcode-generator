package validator

import (
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

// 0.00000001btc
const BtcSignificantDigit = 8

func BtcDecimalLengthValidate(fl validator.FieldLevel) bool {
	decimal := fl.Field().String()
	di := strings.Index(decimal, ".")
	if di != -1 && len(decimal[di+1:]) > BtcSignificantDigit {
		return false
	}
	return true
}
