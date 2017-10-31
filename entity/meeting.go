package entity

import (
	"fmt"
	"time"

	"github.com/jack-cheng/CLI-agenda/errors"
)

// ---------------------------------------------------
// data structures definition
// ---------------------------------------------------

// Meeting one meeting entity
type Meeting struct {
	Title         string
	Participators []string
	StartTime     time.Time
	EndTime       time.Time
	Sponsor       string
}

// Meetings all the meetings
type meetings struct {
	allMeetings  map[string]*Meeting            // key: title, value: address of the Meeting entity that has this title
	onesMeetings map[string]map[string]*Meeting // key: user name, value: the meetings the user has participated
}

// AllMeetings only one meetings instance can be accessed
var AllMeetings *meetings

// -----------------------------------------------------
// Meeting structure methods definition
// -----------------------------------------------------

// NewMeeting create a new meeting and add to AllMeetings
func (m *Meeting) NewMeeting(title, start, end string, parts []string) {
	m.validateTitle(title)
	m.validateParticipators(parts)
	startTime := getTime(start)
	endTime := getTime(end)
	m.validateTime(startTime, endTime)
	m.validateNoConflicts(parts, startTime, endTime)

	m = &Meeting{
		Title:         title,
		Participators: parts,
		StartTime:     startTime,
		EndTime:       endTime,
		Sponsor:       GetCurrentUser().UserName,
	}

	AllMeetings.allMeetings[title] = m
	for _, part := range parts {
		if AllMeetings.onesMeetings[part] == nil {
			AllMeetings.onesMeetings[part] = make(map[string]*Meeting)
		}

		AllMeetings.onesMeetings[part][title] = m
	}
	if AllMeetings.onesMeetings[m.Sponsor] == nil {
		AllMeetings.onesMeetings[m.Sponsor] = make(map[string]*Meeting)
	}

	AllMeetings.onesMeetings[m.Sponsor][title] = m

	m.successCreation(title)
}

// check if title has existed
func (m *Meeting) validateTitle(title string) {
	if AllMeetings.allMeetings[title] != nil {
		errors.ErrorMsg("meeting \"" + title + "\" has been created! Please use another title")
	}
}

// check if all the participators have registered
func (m *Meeting) validateParticipators(parts []string) {
	for _, part := range parts {
		flag := false

		for _, user := range users {
			if part == user.UserName {
				flag = true
			}
		}

		if !flag {
			errors.ErrorMsg("participator " + part + " has not registered!")
		}
	}
}

// check if start time is less than end time
func (m *Meeting) validateTime(start, end time.Time) {
	if start.After(end) || start.Equal(end) {
		errors.ErrorMsg("start time of meeting must be less than end time!")
	}
}

// check if there are confilts
func (m *Meeting) validateNoConflicts(parts []string, start, end time.Time) {
	for _, part := range parts {
		for _, ms := range AllMeetings.onesMeetings[part] {
			if !(end.Before(ms.StartTime) || end.Equal(ms.StartTime) ||
				start.After(ms.EndTime) || start.Equal(ms.EndTime)) {
				errors.ErrorMsg("participator " + part + " has another meeting which conflicts this meeting!")
			}
		}
	}
}

// display success message
func (m *Meeting) successCreation(title string) {
	fmt.Printf("meeting %s is successfully created!\n", title)
}

// -----------------------------------------------------
// helpful function
// -----------------------------------------------------

// convert string to time.Time
func getTime(t string) time.Time {
	tmpTime, err := time.Parse("2006-01-02", t)
	if err != nil {
		errors.ErrorMsg("invalid time format: " + t)
	}

	return tmpTime
}

// ------------------------------------------------------
// query meetings methods
// ------------------------------------------------------

// GetMeetings show meetings between time interval [start, end]
func GetMeetings(start, end string) {
	startTime := getTime(start)
	endTime := getTime(end)
	curUser := GetCurrentUser().UserName
	flag := false

	fmt.Println(curUser + "'s meetings between " + start + " and " + end + ": ")
	for _, m := range AllMeetings.onesMeetings[curUser] {
		if !(m.StartTime.After(endTime) || m.EndTime.Before(startTime)) {
			flag = true
			fmt.Println("title: " + m.Title)
			fmt.Printf("participators: %v\n", m.Participators)
			fmt.Println("start time: " + m.StartTime.Format("2006-01-02"))
			fmt.Println("end time: " + m.EndTime.Format("2006-01-02"))
			fmt.Println("sponsor: " + m.Sponsor)
		}
	}

	if !flag {
		fmt.Println("none.")
	}
}

// -----------------------------------------------------
// initial and save methods
// -----------------------------------------------------

// InitAllMeetings initialize AllMeetings
func InitAllMeetings() {
	ms := loadAllMeetings()

	AllMeetings = new(meetings)
	AllMeetings.allMeetings = make(map[string]*Meeting)
	AllMeetings.onesMeetings = make(map[string]map[string]*Meeting)
	for _, m := range ms {
		AllMeetings.allMeetings[m.Title] = &m

		for _, person := range m.Participators {
			if AllMeetings.onesMeetings[person] == nil {
				AllMeetings.onesMeetings[person] = make(map[string]*Meeting)
			}

			AllMeetings.onesMeetings[person][m.Title] = &m
		}
	}
}

// SaveAllMeetings save AllMeetings
func SaveAllMeetings() {
	wirteAllMeetings()
}
