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

	if err == nil {
		db.DB().SetMaxIdleConns(maxIdle)
		db.DB().SetMaxOpenConns(maxOpen)
		db.DB().SetConnMaxLifetime(time.Duration(30) * time.Minute)
	}

	return
}

// GetMysql 获取mysql连接
func GetMysql() *gorm.DB {
	db.Set("gorm:table_options", "CHARSET=utf8mb4 ENGINE=InnoDB").
		AutoMigrate(
			&model.User{},
			&model.UserRole{},
			&model.Permission{},
			&model.Role{},
			&model.RolePermission{},
		)
	return db
}

// CloseMysql 关闭mysql
func CloseMysql() {
	if db != nil {
		db.Close()
	}
}
