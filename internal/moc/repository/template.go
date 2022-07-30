package repository

var Template = `package repositories

import (
	"moell/internal/{{.DirName}}/models"
	"moell/pkg/database"
	"moell/pkg/repository"
	"moell/pkg/utils/paginate"
)

type {{.Name}} struct {
	repository.Helper
}

func ({{.AbbrName}} *{{.Name}}) Count() int64 {
	var count int64
	database.DB.Model(&models.{{.Name}}{}).Count(&count)
	return count
}

func ({{.AbbrName}} *{{.Name}}) Get() ([]models.{{.Name}}, error) {
	var items []models.{{.Name}}
	result := database.DB.Find(&items)
	return items, result.Error
}

func ({{.AbbrName}} *{{.Name}}) Paginate(paginateParam *paginate.Param) ([]models.{{.Name}}, int64, error) {
	var items []models.{{.Name}}
	total := {{.AbbrName}}.Count()

	result := database.DB.Scopes(paginate.ORMScope(paginateParam)).Find(&items)

	return items, total, result.Error
}


func ({{.AbbrName}} *{{.Name}}) FirstById(id string) (*models.{{.Name}}, error) {
	var m *models.{{.Name}}

	result := database.DB.First(&m, id)

	return m, result.Error
}

func ({{.AbbrName}} *{{.Name}}) Create(m *models.{{.Name}}) error {
	result := database.DB.Create(m)

	return result.Error
}

func ({{.AbbrName}} *{{.Name}}) Update(id string, fields map[string]interface{}) error {
	return {{.AbbrName}}.UpdateByMap(&models.{{.Name}}{}, id ,fields)
}

func ({{.AbbrName}} *{{.Name}}) Delete(id string) error {
	return {{.AbbrName}}.DeleteById(&models.{{.Name}}{}, id)
}`
