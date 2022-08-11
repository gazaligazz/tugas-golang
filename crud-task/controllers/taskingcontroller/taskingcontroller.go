package taskingcontroller

import (
	"crud-task/models"
	"html/template"
	"net/http"
	"strconv"

	"crud-task/entities"
)

var taskingModel = models.NewTaskingModel()

func Index(response http.ResponseWriter, request *http.Request) {

	tasking, _ := taskingModel.FindAll()

	data := map[string]interface{}{
		"tasking": tasking,
	}

	temp, err := template.ParseFiles("views/tasking/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)

}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/tasking/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var tasking entities.Tasking
		tasking.Task = request.Form.Get("task")
		tasking.Assignee = request.Form.Get("assignee")
		tasking.Deadline = request.Form.Get("deadline")

		var data = make(map[string]interface{})
		data["pesan"] = "Data tasking berhasil disimpan"
		taskingModel.Create(tasking)

		temp, _ := template.ParseFiles("views/tasking/add.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var tasking entities.Tasking
		taskingModel.Find(id, &tasking)

		data := map[string]interface{}{
			"tasking": tasking,
		}

		temp, err := template.ParseFiles("views/tasking/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var tasking entities.Tasking
		tasking.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		tasking.Task = request.Form.Get("task")
		tasking.Assignee = request.Form.Get("assignee")
		tasking.Deadline = request.Form.Get("deadline")

		var data = make(map[string]interface{})
		data["pesan"] = "Data tasking berhasil diperbarui"
		taskingModel.Update(tasking)

		temp, _ := template.ParseFiles("views/tasking/edit.html")
		temp.Execute(response, data)
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {

}
