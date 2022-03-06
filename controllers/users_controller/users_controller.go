package users_controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github/atakanteko/bookstore_users-api/domain/users"
	"io"
	"log"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		//TODO
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		log.Fatal("Error unmarshalling", err.Error())
	}
	fmt.Println(user)
	c.String(http.StatusNotImplemented, "implement me!")
}
