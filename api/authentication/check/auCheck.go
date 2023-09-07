package check

import (
	"github.com/gin-contrib/sessions"
)

func Check(sess sessions.Session) bool {
	logIn := sess.Get("loggedIn")
	if logIn == false {
		return false
	} else {
		return true
	}
}
