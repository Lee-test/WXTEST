package main

import (
	"flag"
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> qry and refound
	"net/http"
	"os"

	"github.com/relax-space/go-kit/model"
<<<<<<< HEAD
=======
	alipay "github.com/relax-space/lemon-alipay-sdk"
>>>>>>> qry and refound

	wxpay "github.com/relax-space/lemon-wxpay-sdk"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	appId    = flag.String("WXPAY_APPID", os.Getenv("WXPAY_APPID"), "WXPAY_APPID")
	key      = flag.String("WXPAY_KEY", os.Getenv("WXPAY_KEY"), "WXPAY_KEY")
	mchId    = flag.String("WXPAY_MCHID", os.Getenv("WXPAY_MCHID"), "WXPAY_MCHID")
	certName = flag.String("CERT_NAME", os.Getenv("CERT_NAME"), "CERT_NAME")
	certKey  = flag.String("CERT_KEY", os.Getenv("CERT_KEY"), "CERT_KEY")
	rootCa   = flag.String("ROOT_CA", os.Getenv("ROOT_CA"), "ROOT_CA")

	aliappId = flag.String("ALIPAY_APPID", os.Getenv("ALIPAY_APPID"), "ALIPAY_APPID")
	priKey   = flag.String("ALIPAY_PRIKEY", os.Getenv("ALIPAY_PRIKEY"), "ALIPAY_PRIKEY")
	pubKey   = flag.String("ALIPAY_PUBKEY", os.Getenv("ALIPAY_PUBKEY"), "ALIPAY_PUBKEY")
)

func main() {
	flag.Parse()

	e := echo.New()
	e.Use(middleware.CORS())
	RegisterApi(e)
	//	e.Use(echomiddleware.ContextDB(db))
	e.Start(":5000")
}

func RegisterApi(e *echo.Echo) {
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "fruit-api")
	// })
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	v1 := e.Group("/v1")
	green := v1.Group("/green")
	green.POST("/pay", Paywx)
	green.POST("/query", Querywx)
	green.POST("/reverse", Reversewx)
	green.POST("/refund", Refundwx)

	green.POST("/pay2", Payal)
	green.POST("/query2", Queryal)
	green.POST("/refund2", Refundal)
}

/********wechat pay********/
func Paywx(c echo.Context) error {

	reqDto := wxpay.ReqPayDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	account := Account()
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId: account.AppId,
		MchId: account.MchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.Pay(&reqDto, &customDto)
	if err != nil {
		if err.Error() == "MESSAGE_PAYING" {
			queryDto := wxpay.ReqQueryDto{
				ReqBaseDto: reqDto.ReqBaseDto,
				OutTradeNo: result["out_trade_no"].(string),
			}
			result, err = wxpay.LoopQuery(&queryDto, &customDto, 40, 2)
			if err == nil {
				return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
			} else {
				reverseDto := wxpay.ReqReverseDto{
					ReqBaseDto: reqDto.ReqBaseDto,
					OutTradeNo: result["out_trade_no"].(string),
				}
				_, err = wxpay.Reverse(&reverseDto, &customDto, 10, 10)
				return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
			}
		} else {
			return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
		}
	}
	return c.JSON(http.StatusOK, SuccessResult(result))
}

/************alipay***********/
func Payal(c echo.Context) error {

	reqDto := alipay.ReqPayDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	account := Account2()
	reqDto.ReqBaseDto = &alipay.ReqBaseDto{
		AppId: account.AppId,
	}
	customDto := alipay.ReqCustomerDto{
		PriKey: account.PriKey,
		PubKey: account.PubKey,
	}

	result, err := alipay.Pay(&reqDto, &customDto)
	if err != nil {
		fmt.Printf("%+v\n", result)
		if err.Error() == "MESSAGE_PAYING" {
			queryDto := alipay.ReqQueryDto{
				ReqBaseDto: reqDto.ReqBaseDto,
				//	OutTradeNo: result["out_trade_no"].(string),
				OutTradeNo: result.OutTradeNo,
			}
			result, err = alipay.LoopQuery(&queryDto, &customDto, 40, 2)
			if err == nil {
				return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
			} else {
				reverseDto := alipay.ReqReverseDto{
					ReqBaseDto: reqDto.ReqBaseDto,
					//		OutTradeNo: result["out_trade_no"].(string),
					OutTradeNo: result.OutTradeNo,
				}
				_, err = alipay.Reverse(&reverseDto, &customDto, 10, 10)
				return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
			}
		} else {
			return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
		}
	}
	fmt.Printf("xxxx\n")
	return c.JSON(http.StatusOK, SuccessResult(result))
}

/****wxqury***/
func Querywx(c echo.Context) error {
	reqDto := wxpay.ReqQueryDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	account := Account()
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId: account.AppId,
		MchId: account.MchId,
	}
	cusomDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.Query(&reqDto, &cusomDto)
	//fmt.Printf("%+v,%v", result, err)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	return c.JSON(http.StatusOK, SuccessResult(result))
}

/**ali qurry**/
func Queryal(c echo.Context) error {
	reqDto := alipay.ReqQueryDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	fmt.Printf("%+v\n", reqDto)
	account := Account2()
	fmt.Printf("%+v\n", account)
	reqDto.ReqBaseDto = &alipay.ReqBaseDto{
		AppId: account.AppId,
	}
	customDto := alipay.ReqCustomerDto{
		PriKey: account.PriKey,
		PubKey: account.PubKey,
	}
	result, err := alipay.Query(&reqDto, &customDto)
	//fmt.Printf("%+v,%v", result, err)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	return c.JSON(http.StatusOK, SuccessResult(result))
}
func Reversewx(c echo.Context) error {
	return nil
}

/** wxrefund**/
func Refundwx(c echo.Context) error {
	reqDto := wxpay.ReqRefundDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	account := Account()
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId: account.AppId,
		MchId: account.MchId,
	}
	cusomDto := wxpay.ReqCustomerDto{
		Key:          account.Key,
		CertPathName: account.CertPathName,
		CertPathKey:  account.CertPathKey,
		RootCa:       account.RootCa,
	}
	result, err := wxpay.Refund(&reqDto, &cusomDto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	return c.JSON(http.StatusOK, SuccessResult(result))
}

/**alirefund**/
func Refundal(c echo.Context) error {
	reqDto := alipay.ReqRefundDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	account := Account2()
	reqDto.ReqBaseDto = &alipay.ReqBaseDto{
		AppId: account.AppId,
	}
	customDto := alipay.ReqCustomerDto{
		PriKey: account.PriKey,
		PubKey: account.PubKey,
	}
	result, err := alipay.Refund(&reqDto, &customDto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	return c.JSON(http.StatusOK, SuccessResult(result))
}
func Account() greenAccount {

	account := greenAccount{AppId: *appId, Key: *key, MchId: *mchId,
		CertPathName: *certName, CertPathKey: *certKey, RootCa: *rootCa,
	}
	return account
}
func Account2() greenAccount2 {

	account := greenAccount2{AppId: *aliappId, PriKey: *priKey, PubKey: *pubKey}
	return account
}

type greenAccount struct {
	AppId        string
	Key          string
	MchId        string
	CertPathName string
	CertPathKey  string
	RootCa       string
}

type greenAccount2 struct {
	AppId  string
	PriKey string
	PubKey string
}

func ErrorResult(errMsg string) (result model.Result) {
	result = model.Result{
		Error: model.Error{Message: errMsg},
	}
	return
}
func SuccessResult(param interface{}) (result model.Result) {
	result = model.Result{
		Success: true,
		Result:  param,
	}
	return
}
