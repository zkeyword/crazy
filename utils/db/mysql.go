package db

import (
	"CRAZY/model"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// StartMysql 初始化mysql
func StartMysql(dsn string, maxIdle int, maxOpen int, LogMode bool) (err error) {
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic("mysql connect error: %v" + err.Error())
	}

	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpen)
	db.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)
	db.LogMode(LogMode)
	db.Set("gorm:table_options", "CHARSET=utf8mb4 ENGINE=InnoDB").
		AutoMigrate(
			// sys
			&model.User{},
			&model.UserRole{},
			&model.Permission{},
			&model.Role{},
			&model.RolePermission{},
			&model.Other{},
			&model.UserAddress{},
			// shop
			&model.ShopCart{},
			&model.ShopOrder{},
		)

	user := &model.User{
		Username: "admin",
		Password: "d36dd63cfd", // admin
		Status:   1,
		Level:    0,
		ParentID: 0,
		RealName: "ADMIN",
	}

	permission := &model.Permission{
		Name:   "全部",
		Key:    "all",
		Status: 1,
	}

	role := &model.Role{
		Name: "管理员",
	}

	userRole := &model.UserRole{
		UserID: 1,
		RoleID: 1,
	}

	rolePermission := &model.RolePermission{
		RoleID:         1,
		PermissionKeys: "all",
	}

	db.FirstOrCreate(user)
	db.FirstOrCreate(permission)
	db.FirstOrCreate(role)

	db.FirstOrCreate(userRole)
	db.FirstOrCreate(rolePermission)

	return
}

// GetMysql 获取mysql连接
func GetMysql() *gorm.DB {
	return db
}

// CloseMysql 关闭mysql
func CloseMysql() {
	if db != nil {
		db.Close()
	}
}
