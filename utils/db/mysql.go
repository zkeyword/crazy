package db

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	InitErr = errors.New("not initialized.")
)

// StartMysql 初始化mysql
func StartMysql(dsn string, maxIdle int, maxOpen int, LogMode bool) (*gorm.DB, error) {
	defaultConfig := mysql.Config{
		DSN:                       dsn, // data source name
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	// connect
	var err error
	db, err = gorm.Open(mysql.New(defaultConfig), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if LogMode {
		db = db.Debug()
	}

	db_, _ := db.DB()
	if err := db_.Ping(); err != nil {
		panic(err)
	}

	db_.SetMaxIdleConns(maxIdle)
	db_.SetMaxOpenConns(maxOpen)
	db_.SetConnMaxLifetime(time.Hour * 8)

	// db.Set("gorm:table_options", "CHARSET=utf8mb4 ENGINE=InnoDB").
	// 	AutoMigrate(
	// 		// sys
	// 		&model.User{},
	// 		&model.UserRole{},
	// 		&model.Permission{},
	// 		&model.Role{},
	// 		&model.RolePermission{},
	// 		&model.Other{},
	// 		&model.UserAddress{},
	// 		// shop
	// 		&model.ShopCart{},
	// 		&model.ShopOrder{},
	// 	)

	// user := &model.User{
	// 	Username: "admin",
	// 	Password: "d36dd63cfd", // admin
	// 	Status:   1,
	// 	Level:    0,
	// 	ParentID: 0,
	// 	RealName: "ADMIN",
	// }

	// permission := &model.Permission{
	// 	Name:   "全部",
	// 	Key:    "all",
	// 	Status: 1,
	// }

	// role := &model.Role{
	// 	Name: "管理员",
	// }

	// userRole := &model.UserRole{
	// 	UserID: 1,
	// 	RoleID: 1,
	// }

	// rolePermission := &model.RolePermission{
	// 	RoleID:         1,
	// 	PermissionKeys: "all",
	// }

	// db.FirstOrCreate(user)
	// db.FirstOrCreate(permission)
	// db.FirstOrCreate(role)

	// db.FirstOrCreate(userRole)
	// db.FirstOrCreate(rolePermission)

	return db, nil
}

// GetMysql 获取mysql连接
func GetMysql() (*gorm.DB, error) {
	fmt.Println(db)
	if db == nil {
		return nil, InitErr
	}
	return db, nil
}
