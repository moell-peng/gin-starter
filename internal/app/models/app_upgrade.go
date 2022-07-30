package models

import (
	"gorm.io/datatypes"
	"moell/pkg/model"
)

type AppUpgrade struct {
	model.Model
	Title          string         `gorm:"type:varchar(50);not null;" json:"title"`
	Contents       datatypes.JSON `json:"contents"`
	VersionCode    int            `gorm:"type:int(5)" json:"version_code"`
	ApkDownloadUrl string         `gorm:"type:varchar(255);" json:"apk_download_url"`
	Force          string         `gorm:"type:varchar(10);" json:"force"`
	Platform       string         `gorm:"type:varchar(20);" json:"platform"`
	WebDownload    string         `gorm:"type:varchar(10)" json:"web_download"`
}
