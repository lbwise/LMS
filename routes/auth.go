package routes

import (
	"fmt"
	"io"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	Username string `json: username`
	Password string `json: password`
}


func AuthRoutes(router *gin.RouterGroup) {
	router.GET("/login", getLogin)
	router.POST("/login", postLogin)
	router.GET("/register", getRegister)
	router.POST("/register", postRegister)
}


func getLogin(c *gin.Context) {
	c.String(200, "THIS IS THE LOGIN PAGE")
}


func postLogin(c *gin.Context) {
	var user UserLogin
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(data, &user)	
	if err != nil {
		panic(err.Error())
	}
	msg := fmt.Sprintf("username: %s\tpassword: %s\n", user.Username, user.Password)
	c.String(200, msg)
}


func getRegister(c *gin.Context) {
	return
}


func postRegister(c *gin.Context) {
	return
}