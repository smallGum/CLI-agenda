package entity

import (
	"encoding/json"
	"io"
	"os"

	"github.com/jack-cheng/CLI-agenda/errors"
)

// read all meetings from the file
func loadAllMeetings() []Meeting {
	file, err := os.OpenFile("meetings.json", os.O_RDONLY, 0)
	if err != nil {
		errors.ErrorMsg("cannot open file meetings.json!")
	}

	ms := make([]Meeting, 0)
	decoder := json.NewDecoder(file)
	for {
		m := new(Meeting)
		if err := decoder.Decode(m); err == io.EOF {
			break
		} else if err != nil {
			errors.ErrorMsg("reading meetings error: " + err.Error())
		}

		ms = append(ms, *m)
	}

	file.Close()
	return ms
}

// write all meetings to the file
func wirteAllMeetings() {
	file, _ := os.OpenFile("meetings.json", os.O_WRONLY|os.O_CREATE, 0)
	encoder := json.NewEncoder(file)

	for _, v := range AllMeetings.allMeetings {
		if v.Title != "" {
			encoder.Encode(*v)
		}
	}

	file.Close()
}
