package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (res BaseResponse) Success(c *gin.Context) {
	c.JSON(http.StatusOK, BaseResponse{
		Code: http.StatusOK,
		Msg:  "success",
		Data: res.Data,
	})
}

func (res BaseResponse) Error(c *gin.Context) {
	c.JSON(http.StatusBadRequest, BaseResponse{
		Code: res.Code,
		Msg:  res.Msg,
		Data: nil,
	})
}
