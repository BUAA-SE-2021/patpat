package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Sid      int
	Password string `gorm:"type:varchar(255)"`
}

type JudgeResultUsual struct {
	gorm.Model
	Num    int
	Sid    int
	Name   string `gorm:"type:varchar(255)"`
	Test   string `gorm:"type:varchar(255)"`
	Result int
}

type JudgeResultFormal struct {
	gorm.Model
	Num    int
	Sid    int
	Name   string `gorm:"type:varchar(255)"`
	Test   string `gorm:"type:varchar(255)"`
	Result int
	Tag    string `gorm:"type:varchar(255)"`
}
