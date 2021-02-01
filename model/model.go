package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Sid      int    `json:"sid"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}

type JudgeResultUsual struct {
	gorm.Model
	Num    int    `json:"num"`
	Sid    int    `json:"sid"`
	Name   string `gorm:"type:varchar(255)" json:"name"`
	Test   string `gorm:"type:varchar(255)" json:"test"`
	Result int    `json:"result"`
}

type JudgeResultFormal struct {
	gorm.Model
	Num    int    `json:"num"`
	Sid    int    `json:"sid"`
	Name   string `gorm:"type:varchar(255)" json:"name"`
	Test   string `gorm:"type:varchar(255)" json:"test"`
	Result int    `json:"result"`
	Tag    string `gorm:"type:varchar(255)" json:"tag"`
}
