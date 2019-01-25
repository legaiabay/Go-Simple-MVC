package todoGroup

import (
	"github.com/kataras/iris"
	"gotest.com/go-sandbox/models/model_todoGroup"
	"gotest.com/go-sandbox/structs"
)

func ReadAllTodoGroup(c iris.Context) {
	var response structs.Response

	if data, message := model_todoGroup.TodoGroupAll(); message == "success" {
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

func ReadAllTodoByGroupId(c iris.Context) {
	response := structs.Response{}
	var params structs.TodoGroup

	if err := c.ReadJSON(&params); err == nil {
		if data, message := model_todoGroup.TodoGroupGetById(params.ID); message == "success" {
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

func CreateTodoGroup(c iris.Context) {
	response := structs.Response{}
	var params structs.TodoGroup

	if err := c.ReadJSON(&params); err == nil {
		if create := model_todoGroup.TodoGroupCreate(params.Title); create == "success" {
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

func UpdateTodoGroup(c iris.Context) {
	response := structs.Response{}
	var params structs.TodoGroup

	if err := c.ReadJSON(&params); err == nil {
		if update := model_todoGroup.TodoGroupUpdateTitle(params.ID, params.Title); update == "success" {
			response.Status = 200
			response.Data = params.Title
			response.Message = update
		} else {
			response.Status = 404
			response.Data = nil
			response.Message = update
		}
	} else {
		response.Status = 200
		response.Data = nil
		response.Message = "Wrong or insufficient params"
	}

	c.JSON(response)
}

func DeleteTodoGroup(c iris.Context) {
	response := structs.Response{}
	var params structs.TodoGroup

	if err := c.ReadJSON(&params); err == nil {
		if delete := model_todoGroup.TodoGroupDelete(params.ID); delete == "success" {
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
