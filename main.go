package main

import (
	"fmt"
	"os"
	"time"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/lbwise/LMS/db"
)



func main() {
	file, _ := os.OpenFile("./log/log.go", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()
	logger := log.New(file, "lms-app: ", log.Default().Flags())
	server := gin.New()
	DB, err := db.ConnectDB()
	logger.Println("CONNECTED TO DB NOW")
	db.ResetDB(DB)
	fmt.Printf("-------- LISTENING ON localhost:8080 --------\n\n")
	if err != nil {
		panic(err.Error())
	}
	server.Use(printLog)
	server.GET("/home", getHome)
	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}
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