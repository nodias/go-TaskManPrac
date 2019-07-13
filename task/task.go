package task

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

func (s status) String() string {
	switch s {
	case UNKNOWN:
		return "UNKNOWN"
	case TODO:
		return "TODO"
	case DONE:
		return "DONE"
	default:
		return ""
	}
}

func (s status) MarshalJSON() ([]byte, error) {
	str := s.String()
	if str == "" {
		return []byte(fmt.Sprintf(`"%s"`, str)), errors.New("status.MarshalJSON: unknown value")
	}
	return []byte(str), nil
}

func (s *status) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "\"UNKNOWN\"":
		*s = UNKNOWN
	case "\"TODO\"":
		*s = TODO
	case "\"DONE\"":
		*s = DONE
	default:
		return errors.New("status.UnmarshalJSON: unknown value")
	}
	return nil
}

type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	d := Deadline{t}
	return &d
}

func (d Deadline) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, d.Unix(), 10), nil
}

func (d *Deadline) UnmarshalJSON(data []byte) error {
	unix, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.Unix(unix, 0)
	return nil
}

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

func (t Task) String() string {
	check := " "
	if t.Status == DONE {
		check = "v"
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

type IncludeSubTasks Task

func (i IncludeSubTasks) indentString(prefix string) string {
	str := prefix + Task(i).String() + "\n"
	for _, st := range i.SubTasks {
		str += IncludeSubTasks(st).indentString(prefix + "  ")
	}
	return str
}

func (i IncludeSubTasks) String() string {
	return i.indentString("")
}
