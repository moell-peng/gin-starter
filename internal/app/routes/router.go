package routes

import (
	"github.com/gin-gonic/gin"
	"moell/internal/app/api"
	"moell/internal/app/middleware"
)

func InitRouter(router *gin.Engine) *gin.Engine {

	v1 := router.Group("/api/v1")
	{
		auth := new(api.Auth)
		v1.POST("/login", auth.Login)
		v1.POST("/register", auth.Register)

		appUpgrade := new(api.AppUpgrade)
		v1.GET("/check-app-upgrade/:platform", appUpgrade.Check)

		v1.Use(middleware.JWT("user"))
		v1.Use(middleware.Auth())

		v1.PATCH("/change-password", auth.ChangePassword)
		v1.PATCH("/frozen-account", auth.FrozenAccount)

		user := new(api.User)
		v1.GET("/user", user.Index)
		v1.POST("/user", user.Store)
		v1.PATCH("/user/:id", user.Update)
		v1.DELETE("/user/:id", user.Destroy)
	}

	return router

}
