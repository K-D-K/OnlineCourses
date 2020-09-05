package user

import (
	"OnlineCourses/controller/user"
	"OnlineCourses/handler"
	"OnlineCourses/models"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

// GET_ALL users of the ORG
func GET_ALL(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	userInstance := user.INSTANCE(db)
	users := userInstance.GetUsers()

	byteArr, _ := json.Marshal(users)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

// POST a user in the database
func POST(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	userInstance := user.INSTANCE(db)
	users := []models.User{}
	decoder := json.NewDecoder(r.Body)
	decoder.Token()

	for decoder.More() {
		user := models.User{}
		err := decoder.Decode(&user)
		if err != nil {
			panic(err)
		}
		userInstance.Create(&user)
		users = append(users, user)
	}

	byteArr, _ := json.Marshal(users)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}
