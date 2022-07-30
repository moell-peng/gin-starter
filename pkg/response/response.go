package response

import (
	"github.com/gin-gonic/gin"
	"moell/pkg/utils/paginate"
	"net/http"
)

const (
	SuccessStatus = "success"
	FailStatus    = "fail"
)

type SuccessResponse struct {
	Code    int            `json:"code"`
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    *paginate.Meta `json:"meta,omitempty"`
}

type FailResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func (f *FailResponse) Error() string {
	return f.Message
}

func NewFail(message string, code int) error {
	return &FailResponse{
		Code:    code,
		Status:  FailStatus,
		Message: message,
		Errors:  nil,
	}
}

func NotContent(c *gin.Context) {
	Success(c, &SuccessResponse{
		Code: http.StatusNoContent,
	})
}

func Created(c *gin.Context) {
	Success(c, &SuccessResponse{
		Code: http.StatusCreated,
	})
}

func Data(c *gin.Context, data interface{}) {
	Success(c, &SuccessResponse{
		Code: http.StatusOK,
		Data: data,
	})
}

func Success(c *gin.Context, r *SuccessResponse) {
	r.Status = SuccessStatus
	if r.Code == 0 {
		r.Code = http.StatusOK
	}
	if r.Message == "" {
		r.Message = "成功"
	}

	c.JSON(r.Code, r)
}

func UnprocessableEntity(c *gin.Context, errors interface{}) {
	Fail(c, &FailResponse{
		Code:    http.StatusUnprocessableEntity,
		Status:  FailStatus,
		Message: "",
		Errors:  errors,
	})
}

func BadRequest(c *gin.Context, message string) {
	Fail(c, NewFail(message, http.StatusBadRequest))
}

func Forbidden(c *gin.Context, message string) {
	Fail(c, NewFail(message, http.StatusForbidden))
}

func NotFound(c *gin.Context, message string) {
	Fail(c, NewFail(message, http.StatusNotFound))
}

func InternalServerError(c *gin.Context, message string) {
	Fail(c, NewFail(message, http.StatusInternalServerError))
}

func BadGateway(c *gin.Context, message string) {
	Fail(c, NewFail(message, http.StatusBadGateway))
}

func Fail(c *gin.Context, err error) {
	var response *FailResponse

	if e, ok := err.(*FailResponse); err != nil && ok {
		response = e
	} else {
		var message string
		if err == nil {
			message = "服务器发生错误"
		} else {
			message = err.Error()
		}

		response = &FailResponse{
			Code:    500,
			Status:  FailStatus,
			Message: message,
		}
	}

	c.JSON(response.Code, response)
}
