package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"moell/config"
	"moell/internal/app/models"
	"moell/internal/app/routes"
	"moell/pkg/database"
	"moell/pkg/logger"
)

func main() {
	if err := config.Load("config/config.yaml"); err != nil {
		fmt.Println("Failed to load configuration")
		return
	}

	logger.InitLogger()

	db, err := database.InitDB()
	if err != nil {
		fmt.Println("err open databases")
		return
	}

	_ = db.AutoMigrate(
		&models.User{},
		&models.AppUpgrade{},
	)

	gin.SetMode(config.Get().GinMode)

	router := gin.Default()

	routes.InitRouter(router)
	/*if err := endless.ListenAndServe(config.Get().Addr, router); err != nil {
		fmt.Println(err.Error())
	}*/

	_ = router.Run(config.Get().Addr)
}
