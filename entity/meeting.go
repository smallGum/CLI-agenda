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

// only one meetings instance can be accessed
var AllMeetings *meetings = nil

// -----------------------------------------------------
// Meeting structure methods definition
// -----------------------------------------------------

// NewMeeting create a new meeting and add to AllMeetings
func (m *Meeting) NewMeeting(title, start, end, sponsor string, parts []string) {
	if AllMeetings == nil {
		AllMeetings.initAllMeetings()
	}

	m.validateTitle(title)
	m.validateParticipators(parts)
	startTime := m.getTime(start)
	endTime := m.getTime(end)
	m.validateTime(startTime, endTime)
	m.validateNoConflicts(parts, startTime, endTime)

	m = &Meeting{
		Title:         title,
		Participators: parts,
		StartTime:     startTime,
		EndTime:       endTime,
		Sponsor:       getSponsor(),
	}

	AllMeetings.allMeetings[title] = m
	for _, part := range parts {
		AllMeetings.onesMeetings[part][title] = m
	}

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

		for _, user := range RegisteredUsers {
			if part == user.UserName {
				flag = true
			}
		}

		if !flag {
			errors.ErrorMsg("participator " + part + " has not registered!")
		}
	}
}

// convert string to time.Time
func (m *Meeting) getTime(t string) time.Time {
	tmpTime, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		errors.ErrorMsg("invalid time format: " + t)
	}

	return tmpTime
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
// meetings structure methods definition
// -----------------------------------------------------
