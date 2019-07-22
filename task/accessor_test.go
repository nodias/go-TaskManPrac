package task

import (
	"reflect"
	"testing"
)

func Test_NewMemoryDataAccess(t *testing.T) {
	cases := []struct {
		title   string
		args    []interface{}
		wantRes MemoryDataAccess
		wantErr error
	}{
		{
			"sucess",
			[]interface{}{},
			MemoryDataAccess{tasks: map[ID]task.Task{}, nextID: int64(1)},
			nil,
		},
	}
	for _, c := range cases {
		res := NewMemoryDataAccess()
		if reflect.TypeOf(c.wantRes) == reflect.TypeOf(res) {
			t.Errorf("Test_MemoryDataAccess_Get - %s \n expect res: %v \n but res: %v", c.title, c.wantRes, res)
		}
	}
}
func Test_MemoryDataAccess_Get(t *testing.T) {
	memoryDataAccess := MemoryDataAccess{
		tasks: map[ID]task.Task{
			"1": {
				"laundry",
				task.TODO,
				nil,
				nil,
			},
		},
		nextID: int64(2),
	}

	cases := []struct {
		title   string
		args    []interface{}
		wantRes task.Task
		wantErr error
	}{
		{
			"sucess",
			[]interface{}{ID("1")},
			task.Task{
				"laundry",
				task.TODO,
				nil,
				nil,
			},
			nil,
		}, {
			"ErrTaskNotExist",
			[]interface{}{ID("2")},
			task.Task{},
			ErrTaskNotExist,
		},
	}
	for _, c := range cases {
		idmock := c.args[0]
		res, err := memoryDataAccess.Get(idmock.(ID))
		if !reflect.DeepEqual(res, c.wantRes) || !reflect.DeepEqual(err, c.wantErr) {
			t.Errorf("Test_MemoryDataAccess_Get - %s \n expect res: %s, err: %s \n but res: %s, err: %s", c.title, c.wantRes, c.wantErr, res, err)
		}
	}
}

func Test_MemoryDataAccess_Put(t *testing.T) {
	memoryDataAccess := MemoryDataAccess{
		tasks: map[ID]task.Task{
			"1": {
				"laundry",
				task.TODO,
				nil,
				nil,
			},
		},
		nextID: int64(2),
	}

	cases := []struct {
		title   string
		args    []interface{}
		wantRes task.Task
		wantErr error
	}{
		{
			"sucess",
			[]interface{}{
				ID("1"),
				task.Task{"laundry", task.DONE, nil, nil},
			},
			task.Task{"laundry", task.DONE, nil, nil},
			nil,
		},
		{
			"ErrTaskNotExist",
			[]interface{}{
				ID("2"),
				task.Task{"laundry", task.DONE, nil, nil},
			},
			task.Task{},
			ErrTaskNotExist,
		},
	}
	for _, c := range cases {
		idmock := c.args[0].(ID)
		taskmock := c.args[1].(task.Task)
		err := memoryDataAccess.Put(idmock, taskmock)
		res, _ := memoryDataAccess.Get(idmock)
		if !reflect.DeepEqual(res, c.wantRes) || !reflect.DeepEqual(err, c.wantErr) {
			t.Errorf("Test_MemoryDataAccess_Put - %s \n expect res: %s, err: %s \n but res: %s, err: %s", c.title, c.wantRes, c.wantErr, res, err)
		}
	}
}

func Test_MemoryDataAccess_Post(t *testing.T) {
	memoryDataAccess := MemoryDataAccess{
		tasks:  map[ID]task.Task{},
		nextID: int64(1),
	}

	cases := []struct {
		title   string
		args    []interface{}
		wantRes task.Task
		wantErr error
	}{
		{
			"sucess",
			[]interface{}{
				task.Task{"laundry", task.TODO, nil, nil},
			},
			task.Task{"laundry", task.TODO, nil, nil},
			nil,
		},
	}
	for _, c := range cases {
		taskmock := c.args[0].(task.Task)
		id, err := memoryDataAccess.Post(taskmock)
		res, _ := memoryDataAccess.Get(id)
		if !reflect.DeepEqual(res, c.wantRes) || !reflect.DeepEqual(err, c.wantErr) {
			t.Errorf("Test_MemoryDataAccess_Post - %s \n expect res: %s, err: %s \n but res: %s, err: %s", c.title, c.wantRes, c.wantErr, res, err)
		}
	}
}

func Test_MemoryDataAccess_Delete(t *testing.T) {
	memoryDataAccess := MemoryDataAccess{
		tasks: map[ID]task.Task{
			"1": {
				"laundry",
				task.TODO,
				nil,
				nil,
			},
		},
		nextID: int64(2),
	}

	cases := []struct {
		title   string
		args    []interface{}
		wantRes task.Task
		wantErr error
	}{
		{
			"sucess",
			[]interface{}{
				ID("1"),
			},
			task.Task{},
			nil,
		},
		{
			"ErrTaskNotExist",
			[]interface{}{
				ID("2"),
			},
			task.Task{},
			ErrTaskNotExist,
		},
	}
	for _, c := range cases {
		idmock := c.args[0].(ID)
		err := memoryDataAccess.Delete(idmock)
		res, _ := memoryDataAccess.Get(idmock)
		if !reflect.DeepEqual(res, c.wantRes) || !reflect.DeepEqual(err, c.wantErr) {
			t.Errorf("Test_MemoryDataAccess_Post - %s \n expect res: %s, err: %s \n but res: %s, err: %s", c.title, c.wantRes, c.wantErr, res, err)
		}
	}
}
