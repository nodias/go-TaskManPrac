package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nodias/go-TaskManPrac/task"
)

type ResponseErr struct {
	Err error `json:"err"`
}

func (r ResponseErr) MarshalJSON() ([]byte, error) {
	if r.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, r.Err)), nil
}

func (r *ResponseErr) UnmarshalJSON(b []byte) error {
	var v interface{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	if v == nil {
		r.Err = nil
		return nil
	}
	switch vt := v.(type) {
	case string:
		if vt == task.ErrTaskNotExist.Error() {
			r.Err = task.ErrTaskNotExist
			return nil
		}
		r.Err = errors.New(vt)
		return nil
	default:
		return errors.New("ResponseErr unmarshalJSON failed")
	}
}

type Response struct {
	Id   task.ID     `json:"id,omitempty"`
	Task task.Task   `json:"task"`
	Err  ResponseErr `json:"err"`
}
