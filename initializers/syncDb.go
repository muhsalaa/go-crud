package initializers

import "github.com/go-crud/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Post{})
}
