package migrate

import (
	"belajarweb/initializers"
	"belajarweb/models"
)

// func init(){
// 	initializers.LoadEnvVariables()
// 	initializers.ConnectToDB()
// }

func Migrate(){
	initializers.DB.AutoMigrate(&models.Post{})
}