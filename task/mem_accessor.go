package task

import "fmt"

type MemoryDataAccess struct {
	tasks  map[ID]Task
	nextID int64
}

func NewMemoryDataAccess() *MemoryDataAccess {
	return &MemoryDataAccess{
		tasks:  map[ID]Task{},
		nextID: int64(1),
	}
}

func (m *MemoryDataAccess) Get(id ID) (Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return Task{}, ErrTaskNotExist
	}
	return t, nil
}

func (m *MemoryDataAccess) Put(id ID, t Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	m.tasks[id] = t
	return nil
}

func (m *MemoryDataAccess) Post(t Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}
