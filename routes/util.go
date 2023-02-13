package routes

import (
	"encoding/json"
	"io"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	lg "github.com/lbwise/LMS/log"
)

func getData(obj *any, c *gin.Context) (any, error) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		lg.Logger.Println(err)
		return nil, errors.New("Unable to read given data")
	}
	json.Unmarshal(data, obj)
	return data, nil
}


func checkLoggedIn(c *gin.Context) {
	loggedIn := getLoggedIn(c)
	if !loggedIn {
		c.Redirect(303, "/api/auth/login")
	}
	c.Next()
}

func getLoggedIn(c *gin.Context) bool {
	var loggedIn bool
	sess, err := getSession(c)
	if err != nil {
		panic(err.Error())
	}
	val, ok := sess.Values["LOGGEDIN"]
	if ok {
		loggedIn = val.(bool)
	}
	return loggedIn
}

func getSession(c *gin.Context) (*sessions.Session, error) {
	sess, ok := c.Get("session")
	session := sess.(*sessions.Session)
	if !ok {
		return nil, errors.New("Could not retreive session")
	}
	return session, nil
}