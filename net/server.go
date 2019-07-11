package net

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/nodias/go-TaskManPrac/task"
)

var m = NewMemoryDataAccess()

const pathPrefix = "api/v1/task/"

func main() {
	http.HandleFunc(pathPrefix, apiHandler)
	log.Fatal(http.ListenAndServe(":8887", nil))
}

func apiHandler(w http.ResponseWriter, req *http.Request) {
	getID := func() (ID, error) {
		id := ID(req.URL.Path[len(pathPrefix):])
		if id == "" {
			return id, errors.New("apiHandler: ID is empty")
		}
		return id, nil
	}
	getTasks := func() ([]task.Task, error) {
		var result []task.Task
		if err := req.ParseForm(); err != nil {
			return nil, err
		}
		encodedTasks, ok := req.PostForm["task"]
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
	switch req.Method {
	case "GET":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		t, err := m.Get(id)
		err = json.NewEncoder(w).Encode(Response{
				Id : id,
				Tasks : t,
				Err : ResponseErr{err},
		})
		if err != nil{
			log.Println(err)
		}
	case "PUT":
		//id, err := getID()
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
		//t, err := getTasks()
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
		//err = m.Put(id, t)

	case "POST":
		panic("unimplemented")
	case "DELETE":
		panic("unimplemented")
	}
}
