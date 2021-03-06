package handler

import (
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	bip21 "github.com/yassun/go-bip21"
	qr "rsc.io/qr"
)

type generateBtcQRParam struct {
	Address string `json:"address" form:"address" query:"address" validate:"required,btcaddress"`
	Amount  string `json:"amount" form:"amount" query:"amount" validate:"btcdecimallength"`
	Label   string `json:"label" form:"label" query:"label" validate:"max=255"`
	Message string `json:"message" form:"message" query:"message" validate:"max=255"`
}

// GenerateBtcQR is Generate QR code based on generateBtcQRParam
func GenerateBtcQR(c echo.Context) error {

	p := new(generateBtcQRParam)

	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	if err := c.Validate(p); err != nil {
		return c.String(http.StatusBadRequest, "Validate is failed: "+err.Error())
	}

	a, err := strconv.ParseFloat(p.Amount, 64)
	if err != nil {
		a = 0
	}

	u := &bip21.URIResources{
		UrnScheme: "bitcoin",
		Address:   p.Address,
		Amount:    a,
		Label:     p.Label,
		Message:   p.Message,
	}

	uri, err := u.BuildURI()
	if err != nil {
		return c.String(http.StatusBadRequest, "BuildURI is failed: "+err.Error())
	}

	qrCode, err := qr.Encode(uri, qr.M)
	if err != nil {
		return c.String(http.StatusBadRequest, "QR Encoding is failed: "+err.Error())
	}

	b64 := base64.StdEncoding.EncodeToString(qrCode.PNG())

	return c.JSON(
		http.StatusOK,
		struct {
			StatusCode int    `json:"status code"`
			QRCode     string `json:"qr"`
		}{
			StatusCode: http.StatusOK,
			QRCode:     b64,
		},
	)
}
