package main

import (
	"encoding/json"
	"fmt"
	"go-kit/model"
	"kit/test"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func Test_Paywx(t *testing.T) {
	bodyStr := `
	{
		"auth_code":"134625966445134769",
		"body":"xiaoxinmiao test",
		"total_fee":1
	}`
	req, err := http.NewRequest(echo.POST, "/v1/green/pay", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, Paywx(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)
}
