package model

import (
	"apiserver/pkg/auth"
	"apiserver/pkg/constvar"
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

type User struct {
	BaseModel
	Username string `json:"username" gorm:"column:name;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Create() error {
	return DB.Self.Create(&u).Error
}

func (u *User) Delete(id uint64) error {
	user := User{}
	user.BaseModel.ID = id
	return DB.Self.Delete(&user).Error
}

func (u *User) Update() error {
	return DB.Self.Save(&u).Error
}

func Index(id uint64) (*User, error) {
	u := &User{}
	d := DB.Self.Where("id=?", id).First(&u)
	return u, d.Error
}
func List(name string, offset, limit int) ([]*User, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*User, 0)
	var count uint64

	where := fmt.Sprintf("name like '%%%s%%'", name)
	if err := DB.Self.Model(&User{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Encrypt the user password.
func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
