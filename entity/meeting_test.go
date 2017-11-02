package entity

import (
	"testing"
	"time"
)

func TestValidateTitle(t *testing.T) {
	AllMeetings = new(meetings)
	AllMeetings.allMeetings = make(map[string]Meeting)
	AllMeetings.allMeetings["test"] = Meeting{}

	if flag := validateTitle("test"); flag != false {
		t.Error("validateTitle test failure!")
	} else {
		t.Log("validateTitle test pass")
	}
}

func TestValidateParticipators(t *testing.T) {
	parts := []string{"none", "nothing"}

	if flag := validateParticipators(parts); flag != false {
		t.Error("validateParticipators test failure!")
	} else {
		t.Log("validateParticipators test pass")
	}
}

func TestValidateTime(t *testing.T) {
	start, _ := time.Parse("2006-01-02", "2017-08-10")
	end, _ := time.Parse("2006-01-02", "2017-07-10")

	if flag := validateTime(start, end); flag != false {
		t.Error("validateTime test failure!")
	} else {
		t.Log("validateTime test pass")
	}
}

func TestValidateNoConflicts(t *testing.T) {
	start, _ := time.Parse("2006-01-02", "2017-07-10")
	end, _ := time.Parse("2006-01-02", "2017-08-10")
	AllMeetings = new(meetings)
	AllMeetings.onesMeetings = make(map[string]map[string]Meeting)
	m := Meeting{
		Title:         "nothing",
		Participators: []string{"yes"},
		StartTime:     start,
		EndTime:       end,
		Sponsor:       "none",
	}
	AllMeetings.onesMeetings["none"] = make(map[string]Meeting)
	AllMeetings.onesMeetings["yes"] = make(map[string]Meeting)
	AllMeetings.onesMeetings["none"]["nothing"] = m
	AllMeetings.onesMeetings["yes"]["nothing"] = m
	parts1 := []string{"yes"}
	parts2 := []string{"none"}
	startTime, _ := time.Parse("2006-01-02", "2017-07-12")
	endTime, _ := time.Parse("2006-01-02", "2017-07-15")

	if flag := validateNoConflicts(parts1, startTime, endTime); flag != false {
		t.Error("validateNoConflicts test failure!")
	} else {
		t.Log("validateNoConflicts test pass")
	}

	if flag := validateNoConflicts(parts2, startTime, endTime); flag != false {
		t.Error("validateNoConflicts test failure!")
	} else {
		t.Log("validateNoConflicts test pass")
	}
}

func TestGetTime(t *testing.T) {
	t1 := "2017-08"
	t2 := "2017-08-39"
	t3 := "2017-14-22"

	if _, flag := getTime(t1); flag != false {
		t.Error("getTime test failure!")
	} else {
		t.Log("getTime test pass")
	}

	if _, flag := getTime(t2); flag != false {
		t.Error("getTime test failure!")
	} else {
		t.Log("getTime test pass")
	}

	if _, flag := getTime(t3); flag != false {
		t.Error("getTime test failure!")
	} else {
		t.Log("getTime test pass")
	}
}
