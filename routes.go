package main

import (
	"github.com/kataras/iris"
	"gotest.com/go-sandbox/controllers/todo"
	"gotest.com/go-sandbox/controllers/todoGroup"
)

func DefineRoutes(app *iris.Application) {
	party := app.Party("/api")
	{
		party.Get("/todo", todo.ReadAllTodo)
		party.Get("/todo/id", todo.ReadTodoById)
		party.Post("/todo/create", todo.CreateTodo)
		party.Put("/todo/update", todo.UpdateTodo)
		party.Delete("/todo/delete", todo.DeleteTodo)

		party.Get("/todo/group", todoGroup.ReadAllTodoGroup)
		party.Get("/todo/group/id", todoGroup.ReadAllTodoByGroupId)
		party.Post("/todo/group/create", todoGroup.CreateTodoGroup)
		party.Put("/todo/group/update", todoGroup.UpdateTodoGroup)
		party.Delete("/todo/group/delete", todoGroup.DeleteTodoGroup)
	}
}
