package entity

import (
	"encoding/json"
	"fmt"
	"io"
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

// func init() {
// 	fmt.Print("a")
// 	fmt.Println(users)
// 	tempUsers := ReadJson("users.json")
// 	for _, value := range tempUsers {
// 		users[value.UserName] = value
// 	}
// 	fmt.Print("b")
// 	fmt.Println(users)
// }

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

// read json file
func ReadJson(filePath string) []User {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	ms := make([]User, 0)
	decoder := json.NewDecoder(file)
	for {
		m := new(User)
		if err := decoder.Decode(m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		ms = append(ms, *m)
	}
	file.Close()
	return ms
}

//convert string to json format
func ToJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// fmt.Println("new string is :" + string(bytes))
	return string(bytes)
}

//write to json file
func WriteJson(contents string, destination string) {
	file, _ := os.OpenFile("./json_files/users.json", os.O_WRONLY|os.O_CREATE, 0)
	encoder := json.NewEncoder(file)

	for _, v := range users {
		if v.UserName != "" {
			encoder.Encode(v)
		}
	}

	file.Close()
}

func usernameIsUnique(registerName string) bool {
	_, exist := users[registerName]
	if exist {
		log.Fatal("this username has been occupied")
		return false
	} else {
		return true
	}
}

func Register(username string, password string) bool {
	// fmt.Print("a")
	// fmt.Println(users)
	tempUsers := ReadJson("./json_files/users.json")
	// fmt.Print("read content :")
	// fmt.Println(tempUsers)
	for _, value := range tempUsers {
		users[value.UserName] = value
	}
	// fmt.Print("b")
	// fmt.Println(users)
	if usernameIsUnique(username) {
		new_user := NewUser(username, password)
		// fmt.Println(users)
		users[username] = new_user
		// fmt.Println(users)
		temp := ToJson(users)
		// fmt.Println(temp)
		WriteJson(temp, "users.json")
		return true
	} else {
		return false
	}
}

// data := []byte(contents)
// destination = "./users.json"
// err := ioutil.WriteFile(destination, data, 0644)
// if err != nil {
// 	log.Fatal(err)
// }
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
