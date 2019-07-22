package task

import "errors"

var ErrTaskNotExist = errors.New("task does not exist")

type ID string

type DataAccess interface {
	Get(id ID) (Task, error)
	Post(t Task) (ID, error)
	Put(id ID, t Task) error
	Delete(id ID) error
}


