package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/swiggy-private/chatbackend/internal"
	"io/ioutil"
	"net/http"
)

type LoginController interface {
	LoginHandler(w http.ResponseWriter, r *http.Request)
}

type loginController struct {
	dao *internal.Dao
}

func NewLoginController(dao *internal.Dao) LoginController {
	return &loginController{dao: dao}
}

func (c loginController) LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	// Check if user exists
	if !c.dao.DoesUserExist(user.Username, user.Password) {
		internal.WriteJson(w, internal.BaseResponse{Message: "User does not exist for login", Status: "failure"}, http.StatusBadRequest)
		return
	}

	// Insert into MySQL
	err = c.dao.UpdateLogin(user.Username, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	internal.WriteJson(w, internal.BaseResponse{Status: "success"}, http.StatusOK)
	fmt.Println("login successfully")
}
