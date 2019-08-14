package main

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"

	"github.com/btcsuite/btcutil/bech32"
	base58 "github.com/itchyny/base58-go"
	"github.com/labstack/echo"
	bip21 "github.com/yassun/go-bip21"
	validator "gopkg.in/go-playground/validator.v9"
	"rsc.io/qr"
)

type BtcParam struct {
	Address string `json:"address" form:"address" query:"address" validate:"required,btcaddress"`
	Amount  string `json:"amount" form:"amount" query:"amount" validate:"required,btcdecimallength"`
	Label   string `json:"label" form:"label" query:"label" validate:"max=255"`
	Message string `json:"message" form:"message" query:"message" validate:"max=255"`
}

// Bech32
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

type Validator struct {
	validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	cv.validator.RegisterValidation("btcdecimallength", BtcDecimalLengthValidate)
	cv.validator.RegisterValidation("btcaddress", BtcAddressValidate)
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	initRouting(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func initRouting(e *echo.Echo) {
	e.POST("/generate-qr/btc", generateBtcQr)
}

func generateBtcQr(c echo.Context) error {

	b := new(BtcParam)

	if err := c.Bind(b); err != nil {
		return c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	if err := c.Validate(b); err != nil {
		return c.String(http.StatusBadRequest, "Validate is failed: "+err.Error())
	}

	a, err := strconv.ParseFloat(b.Amount, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Validate is failed: "+err.Error())
	}

	u := &bip21.URIResources{
		UrnScheme: "bitcoin",
		Address:   b.Address,
		Amount:    a,
		Label:     b.Label,
		Message:   b.Message,
	}

	uri, err := u.BuildURI()
	if err != nil {
		return c.String(http.StatusBadRequest, "BuildURI is failed: "+err.Error())
	}

	qrCode, err := qr.Encode(uri, qr.M)
	if err != nil {
		return c.String(http.StatusBadRequest, "QR Encode is failed: "+err.Error())
	}

	b64 := base64.StdEncoding.EncodeToString(qrCode.PNG())

	return c.JSON(
		http.StatusOK,
		struct {
			StatusCode int    `json:"status code"`
			QRCode     string `json:"QR"`
		}{
			StatusCode: http.StatusOK,
			QRCode:     b64,
		},
	)
}
