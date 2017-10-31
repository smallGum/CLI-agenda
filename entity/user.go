package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Tel      string `json:"telephone"`
}

var users = map[string]User{}
var currentUser = NewUser("guest", "")

func GetCurrentUser() User {
	return currentUser
}

func NewUser(username string, password string) User {
	var user User
	user.UserName = username
	user.Password = password
	user.Email = ""
	user.Tel = ""
	return user
}

//convert string to json format
func ToJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}

// read json file
func ReadJson(filePath string) []User {
	var users []User
	//filePath := "../users.json"
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err) // error in read file
	}
	if err := json.Unmarshal(bytes, &users); err != nil {
		log.Fatal(err) // fail to unmarshal
	}
	return users
}

//write to json file
func WriteJson(contents string, destination string) {
	data := []byte(contents)
	//destination="./users.json"
	err := ioutil.WriteFile(destination, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// func (user *User) logout() {
//
// }
//
// func (user *User) lookupAllUser() {
//
// }
//
// func (user *User) cancelAccount() {
//
// }
//
// func (user *User) lookupMeetings() {
//
// }
//
// func (user *User) cancelMeeting() {
//
// }
//
// func (user *User) quitMeeting() {
//
// }
//
// func (user *User) clearAllMeetings() {
//
// }
