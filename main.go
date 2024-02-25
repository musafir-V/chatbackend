package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/swiggy-private/chatbackend/internal"
	"github.com/swiggy-private/chatbackend/internal/handlers"
	"log"
	"net/http"
)

func main() {
	dao := internal.NewDAO()
	createUserController := handlers.NewCreateUserController(dao)
	loginController := handlers.NewLoginController(dao)
	fetchUsersController := handlers.NewFetchUsersController(dao)
	logoutController := handlers.NewLogoutController(dao)
	sendMessageController := handlers.NewSendMessageController(dao)
	fetchUnreadMessageController := handlers.NewFetchUnreadMessageController(dao)

	http.HandleFunc("/create", createUserController.CreateUserHandler)
	http.HandleFunc("/user", fetchUsersController.FetchUsersHandler)
	http.HandleFunc("/login", loginController.LoginHandler)
	http.HandleFunc("/logout", logoutController.LogoutHandler)
	http.HandleFunc("/sendMessage/", sendMessageController.SendMessageHandler)
	http.HandleFunc("/unreadMessage/", fetchUnreadMessageController.FetchUnreadMessageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
