package handlers

import (
	"github.com/swiggy-private/chatbackend/internal"
	"net/http"
)

type FetchUsersController interface {
	FetchUsersHandler(w http.ResponseWriter, r *http.Request)
}

type fetchUsersController struct {
	dao *internal.Dao
}

func NewFetchUsersController(dao *internal.Dao) FetchUsersController {
	return &fetchUsersController{dao: dao}
}

func (c fetchUsersController) FetchUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := c.dao.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	internal.WriteJson(w, internal.GetUserResponse{Data: users, Status: "success"}, http.StatusOK)
}
