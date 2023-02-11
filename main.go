package main

import (
	"fmt"
	"os"
	"time"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/lbwise/LMS/db"
	"github.com/lbwise/LMS/routes"
)

var Log *log.Logger 


func main() {
	file, _ := os.OpenFile("./log/log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()
	Log = log.New(file, "lms-app: ", log.Default().Flags())
	server := gin.New()
	DB, err := db.ConnectDB()
	db.ResetDB(DB)	
	if err != nil {
		Log.Fatal(err)
		panic(err.Error())
	}
	Log.Println("SERVER STARTING UP")
	Log.Println("CONNECTED TO DB NOW")
	fmt.Printf("-------- LISTENING ON localhost:8080 --------\n\n")

	server.Use(func(c *gin.Context) {
		msg := fmt.Sprintf("lms-app: %s %s %s", c.Request.Method, c.Request.URL, time.Now().Format("2006-01-02"))
		Log.Println(msg)
	})
	server.GET("/home", getHome)
	router := server.Group("/")
	addRoutes(router)

	err = server.Run(":8080")
	if err != nil {
		Log.Fatal(err)
		panic(err.Error())
	}
}


func addRoutes(router *gin.RouterGroup) {
	routes.AuthRoutes(router.Group("/auth"))
	routes.UserRoutes(router.Group("/user"))
	routes.CourseRoutes(router.Group("/study"))

}


func getHome(c *gin.Context) {
	c.String(201, "WELCOME HOME")
}
