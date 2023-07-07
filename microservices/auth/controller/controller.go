package controller

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/sssash18/Digicart/microservices/auth/services"
	"github.com/sssash18/Digicart/pkg/common/database"
	"github.com/sssash18/Digicart/pkg/common/models"
)


func SignUp(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(models.Response{
			Status: "error",
			Err:    "invalid parameters: " + err.Error(),
		})
		w.Write(resp)
		return
	}
	user.UserID = uuid.New().String()
	err = services.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(models.Response{
			Status: "error",
			Err:    err.Error(),
		})
		w.Write(resp)
		return
	}
	authToken, err := services.GenerateJWTToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(models.Response{
			Status: "Error",
			Err:    "Unable to generate auth token: " + err.Error(),
		})
		w.Write(resp)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(models.Response{
		Status: "Success",
		Data: map[string]interface{}{
			"auth_token": authToken,
			"user":       user,
		},
	})
	w.Write(resp)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(models.Response{
			Status: "Error",
			Err:    "Invalid parameters: " + err.Error(),
		})
		w.Write(resp)
		return
	}

	if (user.Email == "" && user.Phone == "") || user.Password == "" || (user.Email != "" && user.Phone != "") {
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(models.Response{
			Status: "Error",
			Err:    "Either email or phone is required, and password is mandatory",
		})
		w.Write(resp)
		return
	}

	foundUser := &models.User{}
	if user.Email != "" {
		db.Find(foundUser, "email=?", user.Email)
	}
	if user.Phone != "" {
		db.Find(foundUser, "phone=?", user.Email)
	}

	if foundUser.UserID == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(models.Response{
			Status: "Error",
			Err:    "User with given credentials does not exist",
		})
		w.Write(resp)
		return
	}

	if foundUser.Password == user.Password {
		authToken, err := services.GenerateJWTToken(foundUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp, _ := json.Marshal(models.Response{
				Status: "Error",
				Err:    "Unable to generate auth token: " + err.Error(),
			})
			w.Write(resp)
			return
		}
		w.WriteHeader(http.StatusOK)
		resp, _ := json.Marshal(models.Response{
			Status: "success",
			Data: map[string]interface{}{
				"auth_token": authToken,
			},
		})
		w.Write(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(models.Response{
		Status: "Error",
		Err:    "Incorrect password",
	})
	w.Write(resp)
}

func LogOut(w http.ResponseWriter, r *http.Request) {

}
