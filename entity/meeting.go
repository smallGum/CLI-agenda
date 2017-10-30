package entity

import (
	"time"
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
type Meetings struct {
	meetings     map[string]*Meeting            // key: title, value: address of the Meeting entity that has this title
	onesMeetings map[string]map[string]*Meeting // key: user name, value: the meetings the user has participated
}

// -----------------------------------------------------
// Meeting structure methods definition
// -----------------------------------------------------
