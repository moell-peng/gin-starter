package handler

import (
	"github.com/gin-gonic/gin"
	"moell/internal/app/requests"
	"moell/internal/app/services"
	"moell/pkg/response"
	"moell/pkg/utils/validator"
)

type User struct {
}

func (u *User) Index(ctx *gin.Context) {
	result, err := services.NewUser().Query(ctx)
	if err != nil {
		response.Fail(ctx, err)
		return
	}

	response.Success(ctx, result)
}

func (u *User) Store(ctx *gin.Context) {
	result := validator.Validate(ctx, &requests.UserRequest{Action: "create"})

	if len(result) > 0 {
		response.UnprocessableEntity(ctx, result)
		return
	}

	if err := services.NewUser().Store(ctx); err != nil {
		response.Fail(ctx, err)
		return
	}

	response.Created(ctx)
}

func (u *User) Update(ctx *gin.Context) {
	result := validator.Validate(ctx, &requests.UserRequest{})

	if len(result) > 0 {
		response.UnprocessableEntity(ctx, result)
		return
	}

	if err := services.NewUser().Update(ctx); err != nil {
		response.Fail(ctx, err)
		return
	}

	response.NotContent(ctx)
}

func (u *User) Destroy(ctx *gin.Context) {
	if err := services.NewUser().Delete(ctx); err != nil {
		response.Fail(ctx, err)
		return
	}

	response.NotContent(ctx)
}
