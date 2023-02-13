package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lbwise/LMS/db"
	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserProfile struct {
	ID string `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserRegister struct {
	UserProfile
	Password string `json:"password"`
}


func AuthRoutes(router *gin.RouterGroup) {
	router.GET("/login", func(c *gin.Context) { c.String(200, "PLEASE LOG IN DUDE")})
	router.POST("/login", postLogin)
	router.POST("/register", postRegister)
}


func postLogin(c *gin.Context) {
	var user UserLogin
	var profile UserProfile
	var password string
	data, _ := io.ReadAll(c.Request.Body)
	err := json.Unmarshal(data, &user)	
	if err != nil {
		panic(err.Error())
	}
	profile.Email = user.Email
	query := fmt.Sprintf(`SELECT user_id, name, password FROM users WHERE email='%s';`, user.Email)
	err = db.DB.QueryRow(query).Scan(&profile.ID, &profile.Name, &password)
	if err != nil {
		panic(err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		c.Redirect(303, "/api/auth/login")
		return
	} else if err != nil {
		panic("AN ERROR OCCURED")
	}
	sess, err := getSession(c)
	sess.Values["LOGGEDIN"] = true
	if err != nil {
		panic(err.Error())
	}
	sess.Save(c.Request, c.Writer)
	c.JSON(202, profile)
}


func postRegister(c *gin.Context) {
	var newUser UserRegister
	data, err  := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(data, &newUser)
	if err != nil {
		panic(err.Error())
	}
	created := time.Now().Format("2006-01-02")
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	query := fmt.Sprintf(`INSERT INTO users (name, email, password, created_on) VALUES ('%s', '%s', '%s', '%s');`,
		newUser.Name,
		newUser.Email,
		hashedPwd,
		created,
	)
	_, err = db.DB.Exec(query)
	if err != nil {
		panic(err.Error())
	}
	c.String(201, "Signed up successfully")
}