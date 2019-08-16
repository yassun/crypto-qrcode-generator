package validator

import (
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	cv.validator.RegisterValidation("btcdecimallength", BtcDecimalLengthValidate)
	cv.validator.RegisterValidation("btcaddress", BtcAddressValidate)
	return cv.validator.Struct(i)
}

func NewCustomValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}
