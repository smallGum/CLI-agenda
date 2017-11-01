package entity

import (
	"fmt"
	"os"
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
func NewMeeting(title, start, end string, parts []string) {
	if !(validateTitle(title)) {
		os.Exit(1)
	}
	if !(validateParticipators(parts)) {
		os.Exit(1)
	}
	startTime, ok1 := getTime(start)
	endTime, ok2 := getTime(end)
	if (!ok1) || (!ok2) {
		os.Exit(1)
	}
	if !(validateTime(startTime, endTime)) {
		os.Exit(1)
	}
	if !(validateNoConflicts(parts, startTime, endTime)) {
		os.Exit(1)
	}

	m := &Meeting{
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
}

// check if title has existed
func validateTitle(title string) bool {
	if AllMeetings.allMeetings[title] != nil {
		errors.ErrorMsg(GetCurrentUser().UserName, "meeting \""+title+"\" has existed. expected another title.")
		return false
	}
	return true
}

// check if all the participators have registered
func validateParticipators(parts []string) bool {
	for _, part := range parts {
		flag := false

		for _, user := range users {
			if part == user.UserName {
				flag = true
			}
		}

		if !flag {
			errors.ErrorMsg(GetCurrentUser().UserName, "meeting participator "+part+" has not registered.")
			return false
		}
	}
	return true
}

// check if start time is less than end time
func validateTime(start, end time.Time) bool {
	if start.After(end) || start.Equal(end) {
		errors.ErrorMsg(GetCurrentUser().UserName, "invalid start time, which should be less than end time")
		return false
	}
	return true
}

// check if there are confilts
func validateNoConflicts(parts []string, start, end time.Time) bool {
	for _, part := range parts {
		for _, ms := range AllMeetings.onesMeetings[part] {
			if !(end.Before(ms.StartTime) || end.Equal(ms.StartTime) ||
				start.After(ms.EndTime) || start.Equal(ms.EndTime)) {
				errors.ErrorMsg(GetCurrentUser().UserName, "participator "+part+" has meeting time conflict.")
				return false
			}
		}
	}
	return true
}

// -----------------------------------------------------
// helpful function
// -----------------------------------------------------

// convert string to time.Time
func getTime(t string) (time.Time, bool) {
	tmpTime, err := time.Parse("2006-01-02", t)
	if err != nil {
		errors.ErrorMsg(GetCurrentUser().UserName, "invalid time format: "+t)
		return time.Time{}, false
	}

	return tmpTime, true
}

// ------------------------------------------------------
// query meetings methods
// ------------------------------------------------------

// GetMeetings show meetings between time interval [start, end]
func GetMeetings(start, end string) {
	startTime, ok1 := getTime(start)
	endTime, ok2 := getTime(end)
	if (!ok1) || (!ok2) {
		os.Exit(1)
	}
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
