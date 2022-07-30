package paginate

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type Meta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
}

type Param struct {
	Page		int		`json:"page"`
	PageSize	int		`json:"page_size"`
}

func RequestParam(c *gin.Context) *Param {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize <= 0 {
		pageSize = 15
	}

	return &Param{
		Page:     page,
		PageSize: pageSize,
	}
}

func ORMScope(param *Param) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if param.Page == 0 {
			param.Page = 1
		}

		offset := (param.Page - 1) * param.PageSize

		return db.Offset(offset).Limit(param.PageSize)
	}
}