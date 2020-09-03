package handler

import (
	"OnlineCourses/datastore"
	couresError "OnlineCourses/utils/error"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

// ExecutorWithDB is used to create DB connection on request start and handle exception gloabally
func ExecutorWithDB(handler func(http.ResponseWriter, *http.Request, *gorm.DB)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db := datastore.GetDBConnection().Begin()
		defer func() {
			if r := recover(); r != nil {
				db.Rollback()
				RespondWithError(w, r.(error))
			} else {
				db.Commit()
			}
		}()
		handler(w, r, db)
		defer db.Close()
	}
}

// RespondwithJSON : generic handling to send response.
func RespondwithJSON(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError : handle errors in project
func RespondWithError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *couresError.APIError:
		byteArr, _ := json.Marshal(map[string]string{"message": err.Error()})
		RespondwithJSON(w, http.StatusBadRequest, byteArr)
	default:
		byteArr, _ := json.Marshal(map[string]string{"message": "Internal Server Error"})
		RespondwithJSON(w, http.StatusBadRequest, byteArr)
	}
}
