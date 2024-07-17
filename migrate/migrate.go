package main

import (
	"CRAZY/config"
	"CRAZY/model/shop"
	"CRAZY/model/system"
	"CRAZY/utils/db"
	"os"
)

func main() {

	// 启动mysql
	_db, _ := db.StartMysql(config.DbConfig.Dsn, config.DbConfig.MaxIdle, config.DbConfig.MaxOpen, config.DbConfig.LogMode)
	_db.Set("gorm:table_options", "CHARSET=utf8mb4 ENGINE=InnoDB").
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

	os.Exit(1)
}
