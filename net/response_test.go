package net

import (
	"errors"
	"reflect"
	"testing"
)

func Test_ResponseErr_MarshalJSON(t *testing.T) {
	var reMock_nil ResponseErr

	cases := []struct {
		title   string
		args    []interface{}
		wantRes string
		wantErr error
	}{
		{
			"success",
			[]interface{}{nil},
			"null",
			nil,
		},
		{
			"success2",
			[]interface{}{ErrTaskNotExist},
			`"task does not exist"`,
			nil,
		},
	}
	for _, c := range cases {
		if c.args[0] != nil {
			reMock_nil = ResponseErr{
				c.args[0].(error),
			}
		} else {
			reMock_nil = ResponseErr{
				nil,
			}
		}
		res, err := reMock_nil.MarshalJSON()
		if err != nil {
			t.Error(err)
			return
		}
		if !reflect.DeepEqual(string(res), c.wantRes) || !reflect.DeepEqual(err, c.wantErr) {
			t.Errorf("ResponseErr_MarshalJSON - %s \n expect res: %s, err: %s \n but res: %s, err: %s", c.title, c.wantRes, c.wantErr, res, err)
		}
	}
}

func Test_ResponseErr_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		title   string
		args    []interface{}
		wantRes ResponseErr
		wantErr string
	}{
		{
			"success - ErrTaskNotExist",
			[]interface{}{[]byte(`"task does not exist"`)},
			ResponseErr{ErrTaskNotExist},
			"",
		},
		{
			"success - v==nil",
			[]interface{}{[]byte("null")},
			ResponseErr{nil},
			"",
		},
		{
			"error - err := json.Unmarshal",
			[]interface{}{[]byte("nil")},
			ResponseErr{nil},
			"invalid character 'i' in literal null (expecting 'u')",
		},
		{
			"error - string but not exist error",
			[]interface{}{[]byte(`"not exist error"`)},
			ResponseErr{errors.New("not exist error")},
			"",
		},
		{
			"error - interface type error",
			[]interface{}{[]byte("false")},
			ResponseErr{nil},
			"ResponseErr unmarshalJSON failed",
		},
	}
	for _, c := range cases {
		res := ResponseErr{nil}
		err := res.UnmarshalJSON(c.args[0].([]byte))
		if !reflect.DeepEqual(res, c.wantRes) || err != nil && !reflect.DeepEqual(err.Error(), c.wantErr) {
			t.Errorf("ResponseErr_UnmarshalJSON - %s \n expect res: %s, err: %s \n but res: %s, err: %s", c.title, c.wantRes, c.wantErr, res, err.Error())
		}
	}
}
