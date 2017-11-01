package entity

import (
	"encoding/json"
	"fmt"
	"io"
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
var currentUser User

func init() {
	allUsers := ReadJson("./json_files/users.json")
	for _, value := range allUsers {
		users[value.UserName] = value
	}
	guest := NewUser("guest", "guest")
	users["guest"] = guest
	current := ReadJson("./json_files/currentUser.json")
	currentUser = current[0]
}

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
	return string(bytes)
}

//write to json file
func WriteJson(contents string, destination string) {
	file, _ := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0)
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
	if usernameIsUnique(username) {
		new_user := NewUser(username, password)
		users[username] = new_user
		// fmt.Println(users)
		WriteJson("", "./json_files/users.json")
		return true
	} else {
		return false
	}
}

func Login(username string, password string) bool {
	if GetCurrentUser().UserName != "guest" {
		log.Fatal("you have already logged in, to switch to another account," +
			"you must log out first")
	}
	user, exist := users[username]
	if exist {
		if user.Password == password {
			temp := ToJson(user)
			ioutil.WriteFile("./json_files/currentUser.json", []byte(temp), 0644)
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (user User) Logout() bool {
	if user.UserName == "guest" {
		log.Fatal("you haven't logged in!")
		return false
	} else {
		temp := ToJson(users["guest"])
		ioutil.WriteFile("./json_files/currentUser.json", []byte(temp), 0644)
		return true
	}
}

func (user User) LookupAllUser() {
	fmt.Println("there are", len(users), " users:")
	fmt.Println("--------------------------")
	if user.UserName == "guest" {
		log.Fatal("only users loged in have access to this")
		return
	} else {
		for _, user := range users {
			fmt.Println("user:" + user.UserName)
			fmt.Println("email:" + user.Email)
			fmt.Println("tel:" + user.Tel)
			fmt.Println("--------------------------")
		}
	}
}

func (user User) CancelAccount() bool {
	if user.UserName != "guest" {
		user.QuitMeeting()
		user.ClearAllMeetings()
		user.Logout()
		delete(users, user.UserName)
		WriteJson("", "./json_files/users.json")
		return true
	} else {
		log.Fatal("you can not cancel guest public account")
		return false
	}
}

func (user User) CancelMeeting(title string) {
	meeting, exist := AllMeetings.onesMeetings[user.UserName][title]
	// fmt.Println(AllMeetings.onesMeetings[user.UserName])
	if exist {
		//TODO cancel the target meeting
		fmt.Println("cancel meeting " + meeting.Title + " called!")
	} else {
		log.Fatal("meeting under the given title sponsored by current user doesn't exist")
	}
}

//
func (user User) QuitMeeting() {
	//TODO remove user from all meeting's participators
	fmt.Println("quitMeeting Called")
}

func (user User) ClearAllMeetings() {
	delete(AllMeetings.onesMeetings, user.UserName)
}

//创建会议
//增删会议
//查询给定时间段的会议
