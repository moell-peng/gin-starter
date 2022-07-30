package repository

import (
	"moell/pkg/database"
)

type Helper struct {
}

func (h *Helper) DeleteById(model interface{}, id string) error {
	return database.DB.Delete(model, id).Error
}

func (h *Helper) UpdateByMap(model interface{}, id string, fields map[string]interface{}) error {
	err := database.DB.Model(model).Where("id=?", id).Updates(fields).Error

	return err
}
