// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"shop/base"
	"shop/logger"
)

func init() {
	orm.RegisterModel(new(base.ShopUser))
}

// AddUser 创建用户
func AddUser(m *base.ShopUser) (ID int64, err error) {
	o := orm.NewOrm()
	ID, err = o.Insert(m)
	if err != nil {
		logger.Logger.Errorf("add user info failed! err:[%v].", err)
		return ID, err
	}
	return ID, nil
}

// GetUser 查询用户通过ID
func GetUserByID(ID int) (user base.ShopUser, err error) {
	o := orm.NewOrm()
	userTable := new(base.ShopUser)
	err = o.QueryTable(userTable).Filter("ID", ID).Filter("Deleted", 0).One(&user)

	return user, err
}

// GetUserByOpenID 根据openID获取用户信息
func GetUserByOpenID(openID string) (user base.ShopUser, err error) {
	o := orm.NewOrm()
	userTable := new(base.ShopUser)
	err = o.QueryTable(userTable).Filter("OpenID", openID).One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(user *base.ShopUser) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(user)
	if err != nil {
		return err
	}
	return nil
}

// GetUserIsExistByOpenID 根据openID判断用户是否存在
func GetUserIsExistByOpenID(openID string) bool {
	o := orm.NewOrm()
	userTable := new(base.ShopUser)
	isOK := o.QueryTable(userTable).Filter("OpenID", openID).Exist()
	return isOK
}
