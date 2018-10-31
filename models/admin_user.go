package models

type AdminUser struct {
	Model
	Username 		string		`gorm:"type:varchar(30);unique_index" json:"username,omitempty"`
	Name			string		`gorm:"type:varchar(30);not null;default:''" json:"name,omitempty"`
	Password		string		`gorm:"type:varchar(64);not null;default:''" json:"password,omitempty"`
}
