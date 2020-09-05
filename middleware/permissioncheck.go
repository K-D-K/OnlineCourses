package middleware

import (
	"OnlineCourses/handler"
	"OnlineCourses/models"
	"OnlineCourses/utils/constants"
	"OnlineCourses/utils/error"
	"net/http"
)

// ValidateAuthorPermissions whether the user is author or not
// Here we check role only. permission for course, section, lessons will be handled in handler
func ValidateAuthorPermissions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userInfo, ok := r.Context().Value(constants.UserInfoKey).(models.User)
		if !ok {
			handler.RespondWithError(w, error.GetAPIError("User Info is nill. Unexpected behavior"))
			return
		}
		if userInfo.Role != 1 {
			handler.RespondWithError(w, error.GetAPIError("Permission denied"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
