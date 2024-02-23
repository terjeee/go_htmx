package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/terjeee/go_htmx/api"
	"github.com/terjeee/go_htmx/internal/tools"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resWriter http.ResponseWriter, request *http.Request) {
		var username string = request.URL.Query().Get("username")
		var token = request.Header.Get("Authorization")
		var error error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(resWriter, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, error = tools.NewDatabase()
		if error != nil {
			api.InternalErrorHandler(resWriter)
			return
		}

		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(resWriter, UnAuthorizedError)
			return
		}

		next.ServeHTTP(resWriter, request)
	})
}
