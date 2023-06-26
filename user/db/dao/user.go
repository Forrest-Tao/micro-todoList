package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"user/db/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{DB: NewDBClient(ctx)}
}
func (dao *UserDao) FindUserByUserName(UserName string) (r *model.User, err error) {
	err = dao.Model(&model.User{}).Where("user_name=?", UserName).Find(&r).Error
	return
}
func (dao *UserDao) Create(info *model.User) error {
	return dao.Model(&model.User{}).Create(&info).Error
}
