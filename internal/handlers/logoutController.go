package handlers

import (
	"encoding/json"
	"github.com/swiggy-private/chatbackend/internal"
	"io/ioutil"
	"net/http"
)

type LogoutController interface {
	LogoutHandler(w http.ResponseWriter, r *http.Request)
}

type logoutController struct {
	dao *internal.Dao
}

func NewLogoutController(dao *internal.Dao) LogoutController {
	return &logoutController{dao: dao}
}

func (c logoutController) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user struct {
		Username string `json:"username"`
	}
	err = json.Unmarshal(body, &user)
	// Update the login status in MySQL
	err = c.dao.UpdateLogin(user.Username, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	internal.WriteJson(w, internal.BaseResponse{Status: "success"}, http.StatusOK)
}
