package model

import "gorm.io/gorm"

var models = []interface{}{
	&CityWay{},
}

//定义表格通用前缀
func tableNamePrefix() string {
	return "tb_"
}

// AutoMigrate 自动迁移
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(models...)
}
