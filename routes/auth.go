package routes

import (
	"fmt"
	"io"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lbwise/LMS/db"
)

type UserLogin struct {
	Username string `json: "username"`
	Password string `json: "password"`
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
	var name string
	data, _ := io.ReadAll(c.Request.Body)
	err := json.Unmarshal(data, &user)	
	if err != nil {
		panic(err.Error())
	}
	query := fmt.Sprintf(`SELECT name FROM users WHERE password='%s'`, user.Password)
	err = db.DB.QueryRow(query).Scan(&name)
	if err != nil {
		panic(err.Error())
	}
	c.String(200, name)
}


func getRegister(c *gin.Context) {
	return
}


func postRegister(c *gin.Context) {
	return
}