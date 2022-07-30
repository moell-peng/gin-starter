package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"moell/internal/app/models"
	"moell/pkg/database"
	"moell/pkg/response"
	"net/http"
)

type AppUpgrade struct {
}

func (e *AppUpgrade) Check(c *gin.Context) {
	var appUpgrade models.AppUpgrade

	result := database.DB.Where("platform", c.Param("platform")).
		Order("version_code desc").
		Select("title", "contents", "version_code", "apk_download_url", "platform", "web_download").
		First(&appUpgrade)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			response.Success(c, &response.SuccessResponse{
				Code: http.StatusOK,
			})
			return
		} else {
			response.Fail(c, &response.FailResponse{
				Code:    http.StatusOK,
				Message: "检查更新异常",
			})
			return
		}
	}

	response.Success(c, &response.SuccessResponse{
		Code: http.StatusOK,
		Data: appUpgrade,
	})
}
