package repositories

import (
	"moell/internal/app/models"
	"moell/pkg/database"
	"moell/pkg/repository"
	"moell/pkg/utils/paginate"
)

type User struct {
	repository.Helper
}

func (u *User) Count() int64 {
	var count int64
	database.DB.Model(&models.User{}).Count(&count)
	return count
}

func (u *User) Get() ([]models.User, error) {
	var items []models.User
	result := database.DB.Find(&items)
	return items, result.Error
}

func (u *User) Paginate(paginateParam *paginate.Param) ([]models.User, int64, error) {
	var items []models.User
	total := u.Count()

	result := database.DB.Scopes(paginate.ORMScope(paginateParam)).Find(&items)

	return items, total, result.Error
}

func (u *User) FirstById(id string) (*models.User, error) {
	var m *models.User

	result := database.DB.First(&m, id)

	return m, result.Error
}

func (u *User) Create(m *models.User) error {
	result := database.DB.Create(m)

	return result.Error
}

func (u *User) Update(id string, fields map[string]interface{}) error {
	return u.UpdateByMap(&models.User{}, id, fields)
}

func (u *User) Delete(id string) error {
	return u.DeleteById(&models.User{}, id)
}
