package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (dao *UserDao) Insert(ctx context.Context, u User) error {
	//存更新时间
	now := time.Now().UnixMilli()
	u.CreateTime = now
	u.UpdateTime = now
	return dao.db.Create(&u).Error
}

// User 对标数据库
// 有人叫model， 也有叫 PO(persistent object)
type User struct {
	Id       int64
	Email    string
	Password string

	//创建时间 -毫秒数
	CreateTime int64
	UpdateTime int64
}
