package main

import (
	"github.com/kataras/iris"
	"gotest.com/go-sandbox/controllers"
)

func DefineRoutes(app *iris.Application) {
	party := app.Party("/api")
	{
		party.Get("/mahasiswa", controllers.ReadMahasiswa)
		party.Post("/mahasiswa", controllers.CreateMahasiswa)
	}
}
