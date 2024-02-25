package handlers

import (
	"encoding/json"
	"github.com/swiggy-private/chatbackend/internal"
	"io/ioutil"
	"net/http"
	"strings"
)

type SendMessageController interface {
	SendMessageHandler(w http.ResponseWriter, r *http.Request)
}

type sendMessageController struct {
	dao *internal.Dao
}

func NewSendMessageController(dao *internal.Dao) SendMessageController {
	return &sendMessageController{dao: dao}
}

func (c sendMessageController) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathParts := strings.Split(path, "/")
	if len(pathParts) != 4 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	username := pathParts[len(pathParts)-2]
	if !c.dao.DoesUserExistFromUsername(username) {
		internal.WriteJson(w, internal.BaseResponse{Message: "User doesnot exist", Status: "failure"}, http.StatusBadRequest)
		return
	}
	if !c.dao.GetLoginStatus(username) {
		internal.WriteJson(w, internal.BaseResponse{Message: "User is not logged in", Status: "failure"}, http.StatusBadRequest)
		return
	}
	var message struct {
		To   string `json:"to"`
		Text string `json:"text"`
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.dao.InsertMessage(username, message.To, message.Text)
	internal.WriteJson(w, internal.BaseResponse{Status: "success"}, http.StatusOK)
}
