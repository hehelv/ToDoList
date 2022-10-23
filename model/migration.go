package model

func migration() {
	_ = DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{})
}
