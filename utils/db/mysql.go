package db

import (
	"CRAZY/model/cms"
	"CRAZY/model/shop"
	"CRAZY/model/system"
	"errors"
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

	_db, _ := db.DB()
	if err := _db.Ping(); err != nil {
		panic(err)
	}

	_db.SetMaxIdleConns(maxIdle)
	_db.SetMaxOpenConns(maxOpen)
	_db.SetConnMaxLifetime(time.Hour * 8)

	db.Set("gorm:table_options", "CHARSET=utf8mb4 ENGINE=InnoDB").
		AutoMigrate(
			// sys
			&system.User{},
			&system.UserRole{},
			&system.Permission{},
			&system.Role{},
			&system.RolePermission{},
			&system.Other{},

			// shop
			&shop.UserAddress{},
			&shop.ShopCart{},
			&shop.ShopOrder{},

			// cms
			&cms.Post{},
		)

	// user := &system.User{
	// 	Username:    "admin",
	// 	Password:    "d36dd63cfd", // admin
	// 	Status:      1,
	// 	LoginStatus: 1,
	// 	RealName:    "ADMIN",
	// }

	// permission := &system.Permission{
	// 	Name:   "全部",
	// 	Key:    "All",
	// 	Status: 1,
	// }

	// role := &system.Role{
	// 	Name: "管理员",
	// }

	// userRole := &system.UserRole{
	// 	UserID: 1,
	// 	RoleID: 1,
	// }

	// rolePermission := &system.RolePermission{
	// 	RoleID:         1,
	// 	PermissionKeys: "All",
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
	if db == nil {
		return nil, InitErr
	}
	return db, nil
}
