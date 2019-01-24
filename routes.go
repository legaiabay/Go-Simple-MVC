package main

import (
	"github.com/kataras/iris"
	"gotest.com/go-sandbox/controllers/mahasiswa"
)

func DefineRoutes(app *iris.Application) {
	party := app.Party("/api")
	{
		party.Get("/mahasiswa", mahasiswa.ReadMahasiswa)
		party.Get("/mahasiswa/nim", mahasiswa.ReadMahasiswaByNim)
		party.Post("/mahasiswa", mahasiswa.CreateMahasiswa)
		party.Put("/mahasiswa", mahasiswa.UpdateMahasiswa)
		party.Delete("/mahasiswa", mahasiswa.DeleteMahasiswa)
	}
}
