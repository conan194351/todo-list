package response

import (
	"github.com/conan194351/todo-list.git/pkg/dto/response"
	"github.com/conan194351/todo-list.git/pkg/errs"
	"github.com/gin-gonic/gin"
)

func DefaultResponse(ctx *gin.Context, res response.Response) {
	ctx.JSON(res.Status, res)
}

func SetHttpStatusError(ctx *gin.Context, e errs.ErrorCode, er interface{}) {
	err := errs.ResponseError(e)
	res := response.Response{
		Status:  err.HttpStatusCode,
		Message: err.Message,
		Error:   er,
	}
	DefaultResponse(ctx, res)
}

func SetHttpStatusOK(ctx *gin.Context, status int, message string, data interface{}) {
	res := response.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	DefaultResponse(ctx, res)
}
