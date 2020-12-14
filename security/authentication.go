package security

import (
	"net/http"

	logrus "github.com/sirupsen/logrus"
)

//var log = logrus.New()

//BasicAuthMiddleware ... BasicAuthMiddleware
func BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !checkUsernameAndPassword(user, pass) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		logrus.Infof("Authenticated user: %s", user)
		handler(w, r)
	}
}

func checkUsernameAndPassword(username, password string) bool {
	return username == "abc" && password == "123"
}
