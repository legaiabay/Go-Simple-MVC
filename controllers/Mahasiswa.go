package controllers

import (
	"fmt"
	"log"
	"strconv"

	"gotest.com/go-sandbox/models"
	"gotest.com/go-sandbox/request"
	"gotest.com/go-sandbox/structs"

	"github.com/kataras/iris"
)

// CreateMahasiswa -> Buat data mahasiswa baru
func CreateMahasiswa(c iris.Context) {
	response := structs.Response{}
	var params request.Mahasiswa

	if err := c.ReadJSON(&params); err == nil {
		nama := params.Nama
		kelas := params.Kelas
		nim := params.NIM

		fmt.Println("nama : " + nama)

		if create := models.MahasiswaCreate(nama, kelas, nim); create {
			response.Status = "Ok"
			response.Data = nama
			response.Message = "Create success!"
		} else {
			response.Status = "Error"
			response.Data = "nil"
			response.Message = "Create error!"
		}

	} else {
		fmt.Println("error")
		log.Print(err)
	}

	c.JSON(response)
}

//ReadMahasiswa -> Liat semua data mahasiswa
func ReadMahasiswa(c iris.Context) {

	response := structs.Response{
		Status:  "OK",
		Data:    models.MahasiswaAll(),
		Message: "Success",
	}

	c.JSON(response)
}

//UpdateMahasiswa -> Update mahasiswa by ID
func UpdateMahasiswa(c iris.Context) {
	id := c.Params().GetIntDefault("id", 0)

	c.JSON(iris.Map{
		"message": "update : " + strconv.Itoa(id),
	})
}

//DeleteMahasiswa -> Hapus mahasiswa by ID
func DeleteMahasiswa(c iris.Context) {
	id := c.Params().GetIntDefault("id", 0)

	c.JSON(iris.Map{
		"message": "delete : " + strconv.Itoa(id),
	})
}
