package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
	// cookie secret key
	// TODO: Cookies cret key
	key   = []byte("qiotoSecretKey")
	store = sessions.NewCookieStore(key) // using cookie store

	// TODO: use redis store
	// store = sessions.NewRedisStore()
)

// create session
func CreateSession(c *gin.Context, userID string) error {
	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		return err
	}

	// save useer data in session
	session.Values["user_id"] = userID
	session.Options = &sessions.Options{
		Path:     "/",      // all routes
		MaxAge:   3600 * 2, // TTL
		HttpOnly: true,     // xss security
	}
	return session.Save(c.Request, c.Writer)
}

// get active session
func GetSessionUserID(c *gin.Context) (string, error) {
	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		return "", err
	}

	// check session data
	userID, ok := session.Values["user_id"].(string)
	if !ok || userID == "" {
		return "", nil // session unfound
	}
	return userID, nil
}

// delete session
func DestroySession(c *gin.Context) error {
	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		return err
	}

	// delete data from session
	session.Options.MaxAge = -1
	return session.Save(c.Request, c.Writer)
}
