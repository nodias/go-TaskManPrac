package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/nodias/go-TaskManPrac/task"
)

var pathPrefix = "/api/v1/task/"

func main() {
	http.HandleFunc(pathPrefix, myhandler)
	log.Fatal(http.ListenAndServe(":7000", nil))
}

func myhandler(w http.ResponseWriter, r *http.Request) {
	getID := func() (ID, error) {
		id := ID(r.URL.Path[len(pathPrefix):])
		if id == "" {
			return id, errors.New("apiHandler: ID is empty")
		}
		return id, nil
	}

	getTasks := func() ([]task.Task, error) {
		var result []task.Task
		err := r.ParseForm()
		if err!= nil{
			return nil, err
		}
		encodedTasks, ok := r.PostForm["task"]
		if !ok {
			return nil, errors.New("task parameter expected")
		}
		for _, encodedTask := range encodedTasks{
			var t task.Task
			if err:= json.Unmarshal([]byte(encodedTask), &t); err != nil {
				return nil, err
			}
			result = append(result, t)
		}
		return result, nil
	}

	switch r.Method {
	case "GET":
		panic("not implement")
	case "POST":
		panic("not implement")
	case "PUT":
		panic("not implement")
	case "DELETE":
		panic("not implement")
	}

	fmt.Println(getID())
	fmt.Println(getTasks())

}
