package main

import (
	"fmt"
	"time"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	sess "github.com/gorilla/sessions"
	"github.com/lbwise/LMS/db"
	lg "github.com/lbwise/LMS/log"
	"github.com/lbwise/LMS/routes"
)

var logger *log.Logger			// shouldnt have to import log pkg
var cookie *sess.CookieStore 

// want to init all variables
func init() {
	logger, _ = lg.CreateLog()
	key := []byte(os.Getenv("SESSION_KEY"))
	fmt.Println(string(key))
	cookie = sess.NewCookieStore(key)
	cookie.MaxAge(300)
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
		session, _ := cookie.Get(c.Request, "user-session")
		c.Set("session", session)
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
