package service

var Template = `package services

import (
	"github.com/gin-gonic/gin"
	"moell/internal/{{.DirName}}/models"
	"moell/internal/{{.DirName}}/repositories"
	"moell/pkg/response"
	"moell/pkg/utils/ginparam"
	"moell/pkg/utils/paginate"
)

type {{.Name}} struct {
	repo *repositories.{{.Name}}
}

func ({{.AbbrName}} *{{.Name}}) Query(c *gin.Context) (*response.SuccessResponse, error){
	paginateParam := paginate.RequestParam(c)

	items, total, err := {{.AbbrName}}.repo.Paginate(paginateParam)

	return &response.SuccessResponse{
		Data:    items,
		Meta:    &paginate.Meta{
			CurrentPage: paginateParam.Page,
			PerPage:     paginateParam.PageSize,
			Total:       total,
		},
	}, err
}

func ({{.AbbrName}} *{{.Name}}) Store(c *gin.Context) error {
	var m models.{{.Name}}
	if err := c.ShouldBind(&m); err != nil {
		return err
	}

	return {{.AbbrName}}.repo.Create(&m);
}

func ({{.AbbrName}} *{{.Name}}) Update(c *gin.Context) error {
	data := ginparam.PostKeysToMap(c, []string{
		
	})

	return {{.AbbrName}}.repo.Update(c.Param("id"), data);
}

func ({{.AbbrName}} *{{.Name}}) Delete(c *gin.Context) error {
	return {{.AbbrName}}.repo.Delete(c.Param("id"));
}

func New{{.Name}}() *{{.Name}} {
	return &{{.Name}}{
		repo: &repositories.{{.Name}}{},
	}
}`
