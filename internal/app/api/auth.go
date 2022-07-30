package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"moell/internal/app/models"
	"moell/internal/app/requests"
	"moell/pkg/auth"
	"moell/pkg/database"
	"moell/pkg/response"
	"moell/pkg/utils/validator"
	"net/http"
)

type Auth struct {
}

func (u *Auth) Register(c *gin.Context) {
	result := validator.Validate(c, &requests.RegisterRequest{})
	if len(result) > 0 {
		response.UnprocessableEntity(c, result)
		return
	}

	var count int64
	database.DB.Model(&models.User{}).Where("username = ? ", c.PostForm("username")).Count(&count)
	if count > 0 {
		response.Fail(c, &response.FailResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "用户名已被占用",
			Errors: map[string][]string{
				"username": []string{"用户名已被占用"},
			},
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	if err != nil {
		response.Forbidden(c, err.Error())
		return
	}

	user := models.User{
		Name:     c.PostForm("name"),
		Username: c.PostForm("username"),
		Password: string(hash),
	}
	if err := database.DB.Create(&user).Error; err != nil {
		response.Forbidden(c, err.Error())
		return
	}

	token, err := auth.New().CreateUserToken(user.ID, "user")

	response.Success(c, &response.SuccessResponse{
		Code: http.StatusCreated,
		Data: map[string]interface{}{
			"token": token,
			"user":  user,
		},
	})
}

func (u *Auth) Login(c *gin.Context) {
	result := validator.Validate(c, &requests.LoginRequest{})
	if len(result) > 0 {
		response.UnprocessableEntity(c, result)
		return
	}

	var user models.User
	if err := database.DB.First(&user, "username = ? and status = 0 ", c.PostForm("username")).Error; err != nil {
		response.NotFound(c, "用户不存在")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password")))
	if err != nil {
		response.Forbidden(c, "密码有误")
		return
	}

	token, err := auth.New().CreateUserToken(user.ID, "user")

	response.Success(c, &response.SuccessResponse{
		Data: map[string]interface{}{
			"token": token,
			"user":  user,
		},
	})
}

func (u *Auth) ChangePassword(c *gin.Context) {
	result := validator.Validate(c, &requests.ChangePasswordRequest{})
	if len(result) > 0 {
		response.UnprocessableEntity(c, result)
		return
	}

	var user models.User
	if err := database.DB.First(&user, "id = ? ", auth.New().JwtUserId(c)).Error; err != nil {
		response.NotFound(c, "非法请求")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("old_password")))
	if err != nil {
		response.Forbidden(c, "旧密码有误")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	if err != nil {
		response.Forbidden(c, err.Error())
		return
	}

	if err := database.DB.Model(&user).Update("password", string(hash)).Error; err != nil {
		response.Forbidden(c, err.Error())
		return
	}

	response.NotContent(c)
}

func (u *Auth) FrozenAccount(c *gin.Context) {
	result := validator.Validate(c, &requests.FrozenAccountRequest{})
	if len(result) > 0 {
		response.UnprocessableEntity(c, result)
		return
	}

	var user models.User
	if err := database.DB.First(&user, "id = ? ", auth.New().JwtUserId(c)).Error; err != nil {
		response.NotFound(c, "非法请求")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password"))); err != nil {
		response.Forbidden(c, "旧密码有误")
		return
	}

	if err := database.DB.Model(&user).Update("status", 1).Error; err != nil {
		response.Forbidden(c, err.Error())
		return
	}

	response.NotContent(c)
}
