package model

import "gorm.io/gorm"

// MakeMigrate 迁移数据表
func MakeMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{}) // 用户表
}
