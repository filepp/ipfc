package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"ipfc/api/status"
	"net/http"
	"runtime"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	More string      `json:"more,omitempty"`
	Data interface{} `json:"data"`
}

func respondSuccess(c *gin.Context, data interface{}) {
	respondData(c, 0, data)
}

func respondData(c *gin.Context, code status.Code, data interface{}) {
	if code != status.StatusOK {
		_, file, line, _ := runtime.Caller(1)
		log.Errorf("%v:%v, response %v", file, line, code)
	}
	ret := Response{
		Code: int(code),
		Msg:  code.Text(status.Cn),
		Data: data,
	}
	c.JSON(http.StatusOK, ret)
}

func respondDataEx(c *gin.Context, code status.Code, msg string, data interface{}) {
	if msg == "" {
		msg = code.Text(status.Cn)
	}
	if code != status.StatusOK {
		_, file, line, _ := runtime.Caller(1)
		log.Errorf("%v:%v, response %v, %v", file, line, code, msg)
	}

	ret := Response{
		Code: int(code),
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, ret)
}

func respondError(c *gin.Context, code status.Code, e error) {
	if e == nil {
		e = errors.New("")
	}
	_, file, line, _ := runtime.Caller(1)
	log.Errorf("%v:%v, response %v, %v", file, line, code, e.Error())

	ret := Response{
		Code: int(code),
		Msg:  code.Text(status.Cn),
		Data: nil,
		More: e.Error(),
	}

	c.JSON(http.StatusOK, ret)
}
