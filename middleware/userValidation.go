package middleware

import (
	"OnlineCourses/datastore"
	"OnlineCourses/handler"
	"OnlineCourses/models"
	"OnlineCourses/utils/constants"
	"OnlineCourses/utils/error"
	"context"
	"net/http"
)

// UserValidation for authentication
func UserValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("x-auth-token")
		if len(authToken) == 0 {
			handler.RespondWithError(w, error.GetAPIError("Auth token not present"))
			return
		}
		// In live all user details need to cached in redis. so that we can avoid redundant DB calls.
		// Temp code for validation
		if authToken == "Invalid" {
			handler.RespondWithError(w, error.GetAPIError("Invalid auth token"))
			return
		}

		r = setUserInfoInContext(authToken, r)
		next.ServeHTTP(w, r)
	})
}

func setUserInfoInContext(authToken string, r *http.Request) *http.Request {
	ctx := r.Context()
	user := getUserInfo(authToken)
	return r.WithContext(context.WithValue(ctx, constants.UserInfoKey, user))
}

func getUserInfo(authToken string) models.User {
	// In feature it will be used cache query. temporraily hitting DB
	// Assuming auth token as name :P
	user := models.User{}
	db := datastore.GetDBConnection()
	// Error case deliberately ignored here
	db.Where("name = ?", authToken).First(&user)
	return user
}
