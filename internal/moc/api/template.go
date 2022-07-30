package api

var Template = `package api

import (
	"github.com/gin-gonic/gin"
	"moell/internal/{{.AppDir}}/requests"
	"moell/internal/{{.MSRDir}}/services"
	"moell/pkg/response"
	"moell/pkg/utils/validator"
)

type {{.Name}} struct {

}

func({{.AbbrName}} *{{.Name}}) Index(ctx *gin.Context) {
	result, err := services.New{{.Name}}().Query(ctx)
	if err != nil {
		response.Fail(ctx, err)
		return
	}

	response.Success(ctx, result)
}

func({{.AbbrName}} *{{.Name}}) Store(ctx *gin.Context) {
	result := validator.Validate(ctx, &requests.{{.Name}}Request{})

	if len(result) > 0 {
		response.UnprocessableEntity(ctx, result)
		return
	}

	if err := services.New{{.Name}}().Store(ctx); err != nil {
		response.Fail(ctx, err)
		return
	}

	response.Created(ctx)
}

func({{.AbbrName}} *{{.Name}}) Update(ctx *gin.Context) {
	result := validator.Validate(ctx, &requests.{{.Name}}Request{})

	if len(result) > 0 {
		response.UnprocessableEntity(ctx, result)
		return
	}

	if err := services.New{{.Name}}().Update(ctx); err != nil {
		response.Fail(ctx, err)
		return
	}

	response.NotContent(ctx)
}

func({{.AbbrName}} *{{.Name}}) Destroy(ctx *gin.Context) {
	if err := services.New{{.Name}}().Delete(ctx); err != nil {
		response.Fail(ctx, err)
		return
	}

	response.NotContent(ctx)
}`

var RequestTemplate = `package requests

import "github.com/thedevsaddam/govalidator"

type {{.Name}}Request struct {

}

func (r *{{.Name}}Request) Rules() govalidator.MapData {
	return govalidator.MapData{

	}
}

func (r *{{.Name}}Request) Messages() govalidator.MapData {
	return govalidator.MapData{
		
	}
}`
