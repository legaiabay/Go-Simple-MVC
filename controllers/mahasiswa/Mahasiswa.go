package mahasiswa

import (
	"fmt"
	"log"

	"gotest.com/go-sandbox/models/mMahasiswa"
	"gotest.com/go-sandbox/structs"

	"github.com/kataras/iris"
)

// CreateMahasiswa -> Buat data mahasiswa baru
func CreateMahasiswa(c iris.Context) {
	response := structs.Response{}
	var params structs.Mahasiswa

	if err := c.ReadJSON(&params); err == nil {
		nama := params.Nama
		kelas := params.Kelas
		nim := params.NIM

		fmt.Println("nama : " + nama)

		if create := mMahasiswa.MahasiswaCreate(nama, kelas, nim); create {
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
		Data:    mMahasiswa.MahasiswaAll(),
		Message: "Success",
	}

	c.JSON(response)
}

//UpdateMahasiswa -> Update mahasiswa by ID
func UpdateMahasiswa(c iris.Context) {
	response := structs.Response{}
	var params structs.Mahasiswa

	if err := c.ReadJSON(&params); err == nil {
		nama := params.Nama
		nim := params.NIM

		if update := mMahasiswa.MahasiswaUpdateNama(nim, nama); update {
			response.Status = "Ok"
			response.Data = nama
			response.Message = "Update success!"
		} else {
			response.Status = "Error"
			response.Data = "nil"
			response.Message = "Update error!"
		}

	} else {
		fmt.Println("error")
		log.Print(err)
	}

	c.JSON(response)
}

//DeleteMahasiswa -> Hapus mahasiswa by ID
func DeleteMahasiswa(c iris.Context) {
	response := structs.Response{}
	var params structs.Mahasiswa

	if err := c.ReadJSON(&params); err == nil {
		nim := params.NIM

		if delete := mMahasiswa.MahasiswaDelete(nim); delete {
			response.Status = "Ok"
			response.Data = nim
			response.Message = "Delete success!"
		} else {
			response.Status = "Error"
			response.Data = "nil"
			response.Message = "Delete error!"
		}

	} else {
		fmt.Println("error")
		log.Print(err)
	}

	c.JSON(response)
}
