package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/nodias/go-TaskManPrac/task"
)

const pathPrefix = "/api/v1/task/"
const htmlPrefix = "/task/"

var memoryDataAccess task.DataAccess

func init() {
	memoryDataAccess = task.NewMemoryDataAccess()
}

func main() {
	http.HandleFunc(pathPrefix, myhandler)
	http.HandleFunc(htmlPrefix, htmlHandler)
	log.Println("Server ON!!")
	log.Fatal(http.ListenAndServe(":7000", nil))
}

func myhandler(w http.ResponseWriter, r *http.Request) {
	getID := func() (task.ID, error) {
		id := task.ID(r.URL.Path[len(pathPrefix):])
		if id == "" {
			return id, errors.New("apiHandler: ID is empty")
		}
		return id, nil
	}
	getTasks := func() ([]task.Task, error) {
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
	switch r.Method {
	case "GET":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		t, err := memoryDataAccess.Get(id)
		err = json.NewEncoder(w).Encode(Response{
			id,
			t,
			ResponseErr{err},
		})
		if err != nil {
			log.Println(err)
			return
		}
	case "POST":
		tasks, err := getTasks()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			id, err := memoryDataAccess.Post(t)
			err = json.NewEncoder(w).Encode(Response{
				id,
				t,
				ResponseErr{err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "PUT":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		tasks, err := getTasks()
		if err != nil {
			log.Println(err)
			return
		}
		for _, task := range tasks {
			err := memoryDataAccess.Put(id, task)
			err = json.NewEncoder(w).Encode(Response{
				Id:   id,
				Task: task,
				Err:  ResponseErr{err},
			})
			if err != nil {
				log.Println("err")
				return
			}
		}
	case "DELETE":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		err = memoryDataAccess.Delete(id)
		err = json.NewEncoder(w).Encode(Response{
			Id:  id,
			Err: ResponseErr{err},
		})
		if err != nil {
			log.Println(err)
			return
		}

	}
}

var tmpl = template.Must(template.ParseGlob("net/html/*.html"))

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println(r.Method, "method is not supported")
	}
	getID := func() (task.ID, error) {
		id := task.ID(r.URL.Path[len(htmlPrefix):])
		if id == "" {
			return "", errors.New("htmlHandler : ID is empty")
		}
		return id, nil
	}
	id, err := getID()
	if err != nil {
		log.Println(err)
		return
	}
	t, err := memoryDataAccess.Get(id)
	err = tmpl.ExecuteTemplate(w, "task.html", &Response{
		Id:   id,
		Task: t,
		Err:  ResponseErr{err},
	})
	if err != nil {
		log.Println(err)
		return
	}
}
