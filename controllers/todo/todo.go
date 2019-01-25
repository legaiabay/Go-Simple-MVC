package todo

import (
	"gotest.com/go-sandbox/models/model_todo"
	"gotest.com/go-sandbox/structs"

	"github.com/kataras/iris"
)

func ReadAllTodo(c iris.Context) {
	var response structs.Response

	if data, message := model_todo.TodoAll(); message == "success" {
		response.Status = 200
		response.Data = data
		response.Message = message
	} else {
		response.Status = 404
		response.Data = nil
		response.Message = message
	}

	c.JSON(response)
}

func ReadTodoById(c iris.Context) {
	response := structs.Response{}
	var params structs.Todo

	if err := c.ReadJSON(&params); err == nil {
		if data, message := model_todo.TodoGetById(params.ID); message == "success" {
			response.Status = 200
			response.Data = data
			response.Message = message
		} else {
			response.Status = 404
			response.Data = nil
			response.Message = message
		}
	} else {
		response.Status = 400
		response.Data = nil
		response.Message = "wrong or insufficient params"
	}

	c.JSON(response)
}

func CreateTodo(c iris.Context) {
	response := structs.Response{}
	var params structs.Todo

	if err := c.ReadJSON(&params); err == nil {
		if create := model_todo.TodoCreate(params.Author, params.Note, params.IDGroup); create == "success" {
			response.Status = 200
			response.Data = params
			response.Message = create
		} else {
			response.Status = 404
			response.Data = nil
			response.Message = create
		}
	} else {
		response.Status = 400
		response.Data = nil
		response.Message = "Wrong or insufficient params"
	}

	c.JSON(response)
}

func UpdateTodo(c iris.Context) {
	response := structs.Response{}
	var params structs.Todo

	if err := c.ReadJSON(&params); err == nil {
		if update := model_todo.TodoUpdateNote(params.ID, params.Author, params.Note); update == "success" {
			response.Status = 200
			response.Data = params.Note
			response.Message = update
		} else {
			response.Status = 404
			response.Data = nil
			response.Message = update
		}
	} else {
		response.Status = 400
		response.Data = nil
		response.Message = "Wrong or insufficient params"
	}

	c.JSON(response)
}

func DeleteTodo(c iris.Context) {
	response := structs.Response{}
	var params structs.Todo

	if err := c.ReadJSON(&params); err == nil {
		if delete := model_todo.TodoDelete(params.ID); delete == "success" {
			response.Status = 200
			response.Data = params.ID
			response.Message = delete
		} else {
			response.Status = 404
			response.Data = nil
			response.Message = delete
		}
	} else {
		response.Status = 400
		response.Data = nil
		response.Message = "Wrong or insufficient params"
	}

	c.JSON(response)
}
