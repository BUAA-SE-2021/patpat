package v1

import (
	"errors"
	"fmt"
	"patpat/global"
	"patpat/model"
	"strconv"

	"gorm.io/gorm"
)

func QueryAUser(sid int) (user model.User, notFound bool) {
	err := global.DB.Where("sid = ?", sid).First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return user, true
	} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
		panic(err)
	} else {
		return user, false
	}
}

func Register(sid int, pwd string) (result string) {
	user := model.User{Sid: sid, Password: pwd}
	_, notFound := QueryAUser(sid)
	if notFound {
		if err := global.DB.Create(&user).Error; err != nil {
			panic(err)
		}
		result = "Registration Success."
	} else {
		result = "Fail. Account has been registered or an unexpected error occurred."
	}
	return result
}

func Login(sid int, pwd string) (ok bool) {
	var result string
	user, notFound := QueryAUser(sid)
	if notFound {
		ok = false
		result = "No such user!"
	} else {
		if user.Password != pwd {
			ok = false
			result = "Wrong password!"
		} else {
			ok = true
			result = "Login Success."
		}
	}
	fmt.Println(result)
	return ok
}

func QueryResult(sid int) (result string) {
	var q []model.JudgeResultFormal
	global.DB.Where("sid = ?", sid).Find(&q)
	result = "您每一次实验的评测情况如下:\n"
	for _, v := range q {
		result += "Num = " + strconv.Itoa(v.Num) + ", 评测点 = " + v.Test[0:len(v.Test)-5] + ", Grade = " + strconv.Itoa(v.Result) + " (" + v.Tag + ")\n"
	}
	return result
}
