package mahasiswa

import (
	"gotest.com/go-sandbox/models/mMahasiswa"
	"gotest.com/go-sandbox/structs"

	"github.com/kataras/iris"
)

//ReadMahasiswa -> Liat semua data mahasiswa
func ReadMahasiswa(c iris.Context) {
	var response structs.Response

	if data, message := mMahasiswa.MahasiswaAll(); message == "success" {
		response.Status = "ok"
		response.Data = data
		response.Message = message
	} else {
		response.Status = "error"
		response.Data = nil
		response.Message = message
	}

	c.JSON(response)
}

func ReadMahasiswaByNim(c iris.Context) {
	response := structs.Response{}
	var params structs.Mahasiswa

	if err := c.ReadJSON(&params); err == nil {
		if data, message := mMahasiswa.MahasiswaGetByNim(params.NIM); message == "success" {
			response.Status = "ok"
			response.Data = data
			response.Message = message
		} else {
			response.Status = "error"
			response.Data = nil
			response.Message = message
		}
	} else {
		response.Status = "error"
		response.Data = nil
		response.Message = "wrong or insufficient params"
	}

	c.JSON(response)
}

// CreateMahasiswa -> Buat data mahasiswa baru
func CreateMahasiswa(c iris.Context) {
	response := structs.Response{}
	var params structs.Mahasiswa

	if err := c.ReadJSON(&params); err == nil {
		if create := mMahasiswa.MahasiswaCreate(params.Nama, params.Kelas, params.NIM); create == "success" {
			response.Status = "ok"
			response.Data = params
			response.Message = create
		} else {
			response.Status = "error"
			response.Data = nil
			response.Message = create
		}
	} else {
		response.Status = "error"
		response.Data = nil
		response.Message = "Wrong or insufficient params"
	}

	c.JSON(response)
}

//UpdateMahasiswa -> Update mahasiswa by ID
func UpdateMahasiswa(c iris.Context) {
	response := structs.Response{}
	var params structs.Mahasiswa

	if err := c.ReadJSON(&params); err == nil {
		if update := mMahasiswa.MahasiswaUpdateNama(params.NIM, params.Nama); update == "success" {
			response.Status = "ok"
			response.Data = params.Nama
			response.Message = update
		} else {
			response.Status = "error"
			response.Data = nil
			response.Message = update
		}
	} else {
		response.Status = "error"
		response.Data = nil
		response.Message = "Wrong or insufficient params"
	}

	c.JSON(response)
}

//DeleteMahasiswa -> Hapus mahasiswa by ID
func DeleteMahasiswa(c iris.Context) {
	response := structs.Response{}
	var params structs.Mahasiswa

	if err := c.ReadJSON(&params); err == nil {
		if delete := mMahasiswa.MahasiswaDelete(params.NIM); delete == "success" {
			response.Status = "ok"
			response.Data = params.NIM
			response.Message = delete
		} else {
			response.Status = "error"
			response.Data = nil
			response.Message = delete
		}
	} else {
		response.Status = "error"
		response.Data = nil
		response.Message = "Wrong or insufficient params"
	}

	c.JSON(response)
}
