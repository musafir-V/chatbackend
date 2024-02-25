package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/swiggy-private/chatbackend/internal"
	"io/ioutil"
	"net/http"
)

type CreateUserController interface {
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
}

type createUserController struct {
	dao *internal.Dao
}

func NewCreateUserController(dao *internal.Dao) CreateUserController {
	return &createUserController{dao: dao}
}

func (c createUserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user struct {
		Username string `json:"username"`
		Password string `json:"passcode"`
	}
	err = json.Unmarshal(body, &user)
	if c.dao.DoesUserExistFromUsername(user.Username) {
		internal.WriteJson(w, internal.BaseResponse{Message: "User already exists", Status: "failure"}, http.StatusBadRequest)
		return
	}
	err = c.dao.CreateUser(user.Username, user.Password)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	internal.WriteJson(w, internal.BaseResponse{Message: "User created successfully", Status: "success"}, http.StatusOK)
	fmt.Println("User created successfully")
}
