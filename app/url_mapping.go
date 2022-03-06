package app

import (
	"github/atakanteko/bookstore_users-api/controllers/users_controller"
)

func mapUrls() {
	router.GET("/users/:user_id", users_controller.GetUser)
	router.POST("/users", users_controller.CreateUser)
}
