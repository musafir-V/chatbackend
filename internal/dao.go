package internal

import (
	"database/sql"
	"fmt"
)

type Dao struct {
	db *sql.DB
}

func NewDAO() *Dao {
	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/test_db")
	if err != nil {
		panic(err)
	}
	return &Dao{db: db}
}

func (d *Dao) CreateUser(username, password string) error {
	_, err := d.db.Exec("INSERT INTO user (username, password) VALUES (?, ?)", username, password)
	return err
}
func (d *Dao) DoesUserExistFromUsername(username string) bool {
	row, err := d.db.Query("SELECT username FROM user WHERE username = ?", username)
	if err != nil {
		return false
	}
	for row.Next() {
		return true
	}
	return false
}
func (d *Dao) GetLoginStatus(username string) bool {
	row, err := d.db.Query("SELECT logged_in FROM user WHERE username = ?", username)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(row)
	var loggedIn bool
	row.Next()
	err = row.Scan(&loggedIn)
	if err != nil {
		return false
	}
	return loggedIn
}
func (d *Dao) DoesUserExist(username, password string) bool {
	row, err := d.db.Query("SELECT username FROM user WHERE username = ? and password = ?", username, password)
	if err != nil {
		return false
	}
	for row.Next() {
		return true
	}
	return false
}
func (d *Dao) UpdateLogin(username string, login bool) error {
	_, err := d.db.Exec("UPDATE user SET logged_in = ? WHERE username = ?", login, username)
	return err
}
func (d *Dao) GetAllUsers() ([]string, error) {
	rows, err := d.db.Query("SELECT username FROM user")
	if err != nil {
		return nil, err
	}
	users := make([]string, 0)
	for rows.Next() {
		var user string
		// Scan the row into the User object
		err := rows.Scan(&user)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if user != "" {
			users = append(users, user)
		}
	}
	return users, nil
}
func (d *Dao) InsertMessage(sender, receiver, message string) error {
	_, err := d.db.Exec("INSERT INTO messages (sender, receiver, message, is_read) VALUES (?, ?, ?, ?)", sender, receiver, message, false)
	return err
}
func (d *Dao) GetUnreadMessages(username string) ([]UnreadMessage, error) {
	row, err := d.db.Query("SELECT sender, message FROM messages WHERE receiver = ?", username)
	if err != nil {
		return nil, err
	}
	messages := make(map[string][]string)
	resp := make([]UnreadMessage, 0)
	for row.Next() {
		var sender, message string
		// Scan the row into the User object
		err := row.Scan(&sender, &message)
		if err != nil {
			continue
		}
		messages[sender] = append(messages[sender], message)
	}
	for sender, message := range messages {
		resp = append(resp, UnreadMessage{From: sender, Texts: message})
	}
	return resp, nil
}
