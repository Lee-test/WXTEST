package main

import (
	"flag"
	"go-kit/model"
	"net/http"
	"os"

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
}

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

func Querywx(c echo.Context) error {
	return nil
}

func Reversewx(c echo.Context) error {
	return nil
}

func Refundwx(c echo.Context) error {
	return nil
}

func Account() greenAccount {

	account := greenAccount{AppId: *appId, Key: *key, MchId: *mchId,
		CertPathName: *certName, CertPathKey: *certKey, RootCa: *rootCa,
	}
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
