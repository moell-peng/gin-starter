package middleware

import (
	"github.com/gin-gonic/gin"
	"moell/internal/app/models"
	"moell/pkg/auth"
	"moell/pkg/database"
	"moell/pkg/response"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := database.DB.Where("id=? and status=0", auth.New().JwtUserId(c)).First(&user).Error; err != nil {
			response.Fail(c, &response.FailResponse{
				Code:    http.StatusUnauthorized,
				Message: "当前员工不存在或已被注销",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
