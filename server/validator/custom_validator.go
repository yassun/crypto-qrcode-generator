package validator

import (
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {

	cv.validator.RegisterValidation("btcdecimallength", func(fl validator.FieldLevel) bool {
		return BtcDecimalLengthValidate(fl.Field().String())
	})

	cv.validator.RegisterValidation("btcaddress", func(fl validator.FieldLevel) bool {
		address := fl.Field().String()
		if CheckBtcBase58Address(address) || CheckBech32AddressValidate(address) {
			return true
		}
		return false
	})
	return cv.validator.Struct(i)
}

func NewCustomValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}
