package main

import (
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

	getTasks := func() []task.Task {

		return nil
	}
	fmt.Println(getID())
	fmt.Println(getTasks())

}
