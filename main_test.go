package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/relax-space/go-kit/model"
	"github.com/relax-space/go-kit/test"

	"github.com/labstack/echo"
)

func Test_Paywx(t *testing.T) {
	bodyStr := `
	{
		"auth_code":"134633484682204254",
		"body":"likun test",
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

func Test_Querywx(t *testing.T) {
	bodyStr := `
	{
		"out_trade_no": "1424126122414368810709856110192"
	}`
	req, err := http.NewRequest(echo.POST, "/v1/green/query", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, Querywx(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)
}

func Test_Refoundwx(t *testing.T) {
	bodyStr := `
	{
		"out_trade_no": "1424126122414368810709856110192",
		"refund_fee": 1

	}`
	req, err := http.NewRequest(echo.POST, "/v1/green/refund", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, Refundwx(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)
}

func Test_Payali(t *testing.T) {
	bodyStr := `
	{
		"auth_code":"287175417948547958",
		"subject":"likun test",
		"total_amount":0.01
	}`
	req, err := http.NewRequest(echo.POST, "/v1/green/pay2", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, Payal(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v\n", v)
	test.Equals(t, http.StatusOK, rec.Code)
}

func Test_Qreryali(t *testing.T) {
	bodyStr := `
	{
		"out_trade_no":"111712254348458461972873388"
	}`
	req, err := http.NewRequest(echo.POST, "/v1/green/query2", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, Queryal(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)
}

func Test_Refundali(t *testing.T) {
	bodyStr := `
	{
		"out_trade_no":"111712258546243146624551637",
		"refund_amount":0.01
	}`
	req, err := http.NewRequest(echo.POST, "/v1/green/refund2", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, Refundal(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)
}
