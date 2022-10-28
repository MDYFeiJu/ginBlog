package model

import (
	"fmt"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(400);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	data.Password = BcryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

// GetUsers 获取用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var count int64
	err = db.Find(&users).Count(&count).Error
	if pageNum == -1 {
		err = db.Limit(pageSize).Offset(-1).Find(&users).Error
	} else {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, count
}

// BcryptPw 用户密码加密
func BcryptPw(pw string) string {
	b := []byte(pw)
	fmt.Printf("%v", b)
	password, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	psw := string(password)
	return psw
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	err = db.Where("id=?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditUser 编辑用户信息
func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&User{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func ValidatePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	fmt.Printf("%v", err)
	if err != nil {
		return false
	}
	return true
}

func CheckLogin(name, password string) int {
	var user User
	db.Where("username=?", name).First(&user)
	if user.ID == 0 {
		return errmsg.ErrorUserNotExist
	}
	if !ValidatePasswords(user.Password, password) {
		return errmsg.ErrorPasswordWrong
	}

	//if user.Role != 0 {
	//	return errmsg.UserNoRight
	//}
	return errmsg.SUCCESS
}
