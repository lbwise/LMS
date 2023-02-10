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



func main() {
	file, _ := os.OpenFile("./log/log.go", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()
	logger := log.New(file, "lms-app: ", log.Default().Flags())
	server := gin.New()
	_, err := db.ConnectDB()
	// db.ResetDB(DB)
	if err != nil {
		panic(err.Error())
	}

	logger.Println("CONNECTED TO DB NOW")
	fmt.Printf("-------- LISTENING ON localhost:8080 --------\n\n")

	server.Use(printLog)
	server.GET("/home", getHome)
	router := server.Group("/")
	addRoutes(router)

	err = server.Run(":8080")
	if err != nil {
		panic(err)
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


func printLog(c *gin.Context) {
	fi, err := os.OpenFile("./log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("COULDN'T OPEN LOG")
		panic(err)
	}
	msg := fmt.Sprintf("lms-app: %s %s %s\n", c.Request.Method, c.Request.URL, time.Now().Format("2006-01-02"))
	fi.WriteString(msg)
}