package session

import (
	"os"
	"net/http"
	gsession "github.com/gorilla/sessions"
)

var cookie *gsession.CookieStore 

func CreateStore() {
	key := []byte(os.Getenv("SESSION_KEY"))
	cookie = gsession.NewCookieStore(key)
	cookie.MaxAge(30)
}

func Get(req *http.Request, name string) (*gsession.Session, error) {
	return cookie.Get(req, name)
}
