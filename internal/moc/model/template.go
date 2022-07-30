package model

var Template = `package models

import "moell/pkg/model"

type {{.Name}} struct {
	model.Model
}`
