package main

import (
	"fmt"
	"net/http"

	"crud-task/controllers/taskingcontroller"
)

func main() {
	http.HandleFunc("/", taskingcontroller.Index)
	http.HandleFunc("/tasking", taskingcontroller.Index)
	http.HandleFunc("/tasking/index", taskingcontroller.Index)
	http.HandleFunc("/tasking/add", taskingcontroller.Add)
	http.HandleFunc("/tasking/edit", taskingcontroller.Edit)
	http.HandleFunc("/tasking/delete", taskingcontroller.Delete)

	fmt.Println("Server running di port 8080 ...")
	http.ListenAndServe(":8080", nil)

}
