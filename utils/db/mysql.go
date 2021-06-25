package db

import (
	"CRAZY/model"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// StartMysql 初始化mysql
func StartMysql(dsn string, maxIdle, maxOpen int) (err error) {
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic("mysql connect error: %v" + err.Error())
	}

	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpen)
	db.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)
	db.LogMode(true)
	db.Set("gorm:table_options", "CHARSET=utf8mb4 ENGINE=InnoDB").
		AutoMigrate(
			&model.User{},
			&model.UserRole{},
			&model.Permission{},
			&model.Role{},
			&model.RolePermission{},
			&model.Other{},
		)

		// user := &model.User{
		// 	Username: "admin2",
		// 	Password: "d36dd63cfd",
		// 	Status:   1,
		// 	Level:    0,
		// 	ParentID: 0,
		// }

		// userRole := &model.UserRole{
		// 	UserID:  1,
		// 	RoleIDs: "1,2",
		// }

		// permission := &model.Permission{
		// 	Name:   "全部2",
		// 	Key:    "all2",
		// 	Status: 1,
		// }

		// role := &model.Role{
		// 	Name: "管理员2",
		// }

		// rolePermission := &model.RolePermission{
		// 	RoleID:         2,
		// 	PermissionKeys: "all,all2",
		// }

		// db.Create(user)
		// db.Create(userRole)
		// db.Create(permission)
		// db.Create(role)
		// db.Create(rolePermission)

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
