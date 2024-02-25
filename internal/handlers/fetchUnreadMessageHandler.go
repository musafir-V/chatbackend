package handlers

import (
	"github.com/swiggy-private/chatbackend/internal"
	"net/http"
	"strings"
)

type FetchUnreadMessageController interface {
	FetchUnreadMessageHandler(w http.ResponseWriter, r *http.Request)
}

type fetchUnreadMessageController struct {
	dao *internal.Dao
}

func NewFetchUnreadMessageController(dao *internal.Dao) FetchUnreadMessageController {
	return &fetchUnreadMessageController{dao: dao}
}

func (c fetchUnreadMessageController) FetchUnreadMessageHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathParts := strings.Split(path, "/")
	if len(pathParts) != 4 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	username := pathParts[len(pathParts)-2]
	messages, err := c.dao.GetUnreadMessages(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	internal.WriteJson(w, internal.GetUnreadMessagesResponse{Data: messages, BaseResponse: internal.BaseResponse{Status: "success"}}, http.StatusOK)

}
