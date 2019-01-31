package model

import (
	"sync"
	"time"
)

// BaseModel 模型共有的基础部分
type BaseModel struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-" `
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
}
type UserInfo struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
type UserList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*UserInfo
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
