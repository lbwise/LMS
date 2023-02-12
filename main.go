package main

import (
	"fmt"
	"time"
	"log"

	"github.com/gin-gonic/gin"
	sess "github.com/lbwise/LMS/session"
	"github.com/lbwise/LMS/db"
	lg "github.com/lbwise/LMS/log"
	"github.com/lbwise/LMS/routes"
)

var logger *log.Logger			// shouldnt have to import log pkg

// want to init all variables
func init() {
	sess.CreateStore()
	logger, _ = lg.CreateLog()
	// db.ResetDB(DB)
}


func main() {
	server := gin.New()
	_, err := db.ConnectDB()
	
	if err != nil {
		logger.Fatal(err)
		panic(err.Error())
	}
	logger.Println("SERVER STARTING UP")
	logger.Println("CONNECTED TO DB NOW")
	fmt.Printf("-------- LISTENING ON localhost:8080 --------\n\n")

	server.Use(func(c *gin.Context) {
		session, _ := sess.Get(c.Request, "user-session")
		session.Values["LOGGEDIN"] = false
	})

	server.Use(func(c *gin.Context) {
		logger.Println(c.Request.Method, c.Request.URL, time.Now().Format("2006-01-02"))
	})



	server.GET("/home", getHome)
	router := server.Group("/api")
	addRoutes(router)

	err = server.Run(":8080")
	if err != nil {
		logger.Fatal(err)
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
