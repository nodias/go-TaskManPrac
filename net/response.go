package net

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrTaskNotExist = errors.New("task does not exist")

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
		if vt == ErrTaskNotExist.Error() {
			r.Err = ErrTaskNotExist
			return nil
		}
		r.Err = errors.New(vt)
		return nil
	default:
		return errors.New("ResponseErr unmarshalJSON failed")
	}
}
