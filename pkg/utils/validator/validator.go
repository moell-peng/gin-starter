package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type RequestForm interface {
	Messages() govalidator.MapData
	Rules() govalidator.MapData
}

type NoMessageForm interface {
	Rules() govalidator.MapData
}

func Validate(c *gin.Context, form RequestForm) url.Values {
	opts := govalidator.Options{
		Request:  c.Request,
		Rules:    form.Rules(),
		Messages: form.Messages(),
	}

	return govalidator.New(opts).Validate()
}

func DefaultMessageValidate(c *gin.Context, form NoMessageForm) url.Values {
	opts := govalidator.Options{
		Request: c.Request,
		Rules:   form.Rules(),
	}

	return govalidator.New(opts).Validate()
}
