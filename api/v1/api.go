package v1

import (
	"errors"
	"patpat/global"
	"patpat/model"

	"gorm.io/gorm"
)

func Register(sid int, pwd string) (result string) {
	user := model.User{Sid: sid, Password: pwd}
	var queryUser model.User
	if err := global.DB.Where("sid = ?", sid).First(&queryUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		if err = global.DB.Create(&user).Error; err != nil {
			panic(err)
		}
		result = "Success."
	} else {
		result = "Fail. Account has been registered or an unexpected error occurred."
	}
	return result
}
