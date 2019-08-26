package handler

import (
	"net/http"
	"net/http/httptest"
	"server/validator"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGenerateBtcQR(t *testing.T) {
	type btcQrJSONTest struct {
		given string
		exp   int
	}

	btcQrJSONTests := []btcQrJSONTest{
		{
			`{"address":"175tWpb8K1S7NmH4Zx6rewF9WQrcZv245W"}`,
			http.StatusOK,
		},
		{
			`{"address":"175tWpb8K1S7NmH4Zx6rewF9WQrcZv245W","amount":"0.1"}`,
			http.StatusOK,
		},
		{
			`{"address":"175tWpb8K1S7NmH4Zx6rewF9WQrcZv245W","amount":"0.1", "message": "Donation+for+project+xyz"}`,
			http.StatusOK,
		},
		{
			`{"address":"175tWpb8K1S7NmH4Zx6rewF9WQrcZv245W","amount":"0.1", "message":"Donation+for+project+xyz", "label":"Luke-Jr"}`,
			http.StatusOK,
		},
		{
			`{}`,
			http.StatusBadRequest,
		},
		{
			`{"address":"175tWpb8K1S7NmH4Zx6rewF9WQrcZv245W","amount":"-0.1"}`,
			http.StatusBadRequest,
		},
	}

	e := echo.New()
	e.Validator = validator.NewCustomValidator()

	for _, tt := range btcQrJSONTests {
		t.Run("", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.given))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			if assert.NoError(t, GenerateBtcQR(c)) {
				assert.Equal(t, tt.exp, rec.Code)
			}
		})
	}
}
