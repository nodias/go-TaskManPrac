package main

import (
	"net/http"
	"net/url"
)

func Example_myhandler() {
	req := http.Request{}
	u, _ := url.Parse("http://localhost:7000/api/v1/task/nodias")
	req.URL = u
	myhandler(nil, &req)
	//Output:
	//nodias <nil>
	//[]
}
