package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/nodias/go-TaskManPrac/task"
	"github.com/unrolled/render"
	"log"
	"net/http"
)


var renderer *render.Render

func init(){
	renderer = render.New()
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	t, err := m.Get(id)
	err = renderer.HTML(w, http.StatusOK, "task", &Response{
		Id:   id,
		Task: t,
		Err:  ResponseErr{err},
	})
	if err != nil {
		log.Println(err)
		return
	}
}

var m = task.NewInmemoryAccessor()

func getTasks(r *http.Request) ([]task.Task, error) {
	var result []task.Task
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	encodedTasks, ok := r.PostForm["task"]
	if !ok {
		return nil, errors.New("task parameter expected")
	}
	for _, encodedTask := range encodedTasks {
		var t task.Task
		if err := json.Unmarshal([]byte(encodedTask), &t); err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func apiGetHandler(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	t, err := m.Get(id)
	err = json.NewEncoder(w).Encode(Response{
		Id:   id,
		Task: t,
		Err:  ResponseErr{err},
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func apiPutHandler(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	tasks, err := getTasks(r)
	if err != nil {
		log.Println(err)
		return
	}
	for _, t := range tasks {
		err := m.Put(id, t)
		err = json.NewEncoder(w).Encode(Response{
			Id:   id,
			Task: t,
			Err:  ResponseErr{err},
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func apiPostHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := getTasks(r)
	if err != nil {
		log.Println(err)
		return
	}
	for _, t := range tasks {
		id, err := m.Post(t)
		err = json.NewEncoder(w).Encode(Response{
			Id:   id,
			Task: t,
			Err:  ResponseErr{err},
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func apiDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := task.ID(mux.Vars(r)["id"])
	err := m.Delete(id)
	err = json.NewEncoder(w).Encode(Response{
		Id:  id,
		Err: ResponseErr{err},
	})

}
