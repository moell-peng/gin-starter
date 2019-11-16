package controllers

import (
	"github.com/gin-gonic/gin"
	"gin-gorm-example/models"
	"gin-gorm-example/database"
	"net/http"
	"crypto/md5"
	"encoding/hex"
)

type AdminUser struct {
	Basic
}

func (a *AdminUser) Index(c *gin.Context) {
	var users []models.AdminUser

	database.DB.Select("id, name, username, created_at, updated_at").Order("id").Find(&users)

	a.JsonSuccess(c, http.StatusOK, gin.H{"data": users})
}

func (a *AdminUser) Store(c *gin.Context) {
	var request CreateRequest

	if err := c.ShouldBind(&request); err == nil {
		var count int
		database.DB.Model(&models.AdminUser{}).Where("username = ?", request.Username).Count(&count)

		if (count > 0) {
			a.JsonFail(c, http.StatusBadRequest, "用户名已存在")
			return
		}

		password := []byte(request.Password)
		md5Ctx := md5.New()
		md5Ctx.Write(password)
		cipherStr := md5Ctx.Sum(nil)
		user := models.AdminUser{
			Username: request.Username,
			Name:     request.Name,
			Password: hex.EncodeToString(cipherStr),
		}

		if err := database.DB.Create(&user).Error; err != nil {
			a.JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}

		a.JsonSuccess(c, http.StatusCreated, gin.H{"message": "创建成功"})
	} else {
		a.JsonFail(c, http.StatusBadRequest, err.Error())
	}
}

func (a *AdminUser) Update(c *gin.Context) {
	var request UpdateRequest

	if err := c.ShouldBind(&request); err == nil {
		var user models.AdminUser
		if database.DB.First(&user, c.Param("id")).Error != nil {
			a.JsonFail(c, http.StatusNotFound, "数据不存在")
			return
		}

		user.Name = request.Name

		if err := database.DB.Save(&user).Error; err != nil {
			a.JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}

		a.JsonSuccess(c, http.StatusCreated, gin.H{})
	} else {
		a.JsonFail(c, http.StatusBadRequest, err.Error())
	}
}

func (a *AdminUser) Show(c *gin.Context) {
	var user models.AdminUser

	if database.DB.Select("id, name, username, created_at, updated_at").First(&user, c.Param("id")).Error != nil {
		a.JsonFail(c, http.StatusNotFound, "数据不存在")
		return
	}

	a.JsonSuccess(c, http.StatusCreated, gin.H{"data": user})
}

func (a *AdminUser) Destroy(c *gin.Context) {
	var user models.AdminUser

	if database.DB.First(&user, c.Param("id")).Error != nil {
		a.JsonFail(c, http.StatusNotFound, "数据不存在")
		return
	}

	if err := database.DB.Unscoped().Delete(&user).Error; err != nil {
		a.JsonFail(c, http.StatusBadRequest, err.Error())
		return
	}

	a.JsonSuccess(c, http.StatusCreated, gin.H{})

}

type UpdateRequest struct {
	Name		string `form:"name" json:"name" binding:"required"`
}

type CreateRequest struct {
	Username	string `form:"username" json:"username" binding:"required"`
	Name		string `form:"name" json:"name" binding:"required"`
	Password	string `form:"password" json:"password" binding:"required"`
}



