package models

import (
	"log"
	"strconv"

	"gotest.com/go-sandbox/structs"

	"gotest.com/go-sandbox/config"
)

//MahasiswaAll -> Get all mahasiswa
func MahasiswaAll() []structs.Mahasiswa {
	var mahasiswa structs.Mahasiswa
	var mahasiswaArray []structs.Mahasiswa

	db := config.DBConnect()
	defer db.Close()

	data, err := db.Query("SELECT * FROM mahasiswa")
	if err != nil {
		log.Print(err)
	}

	for data.Next() {
		if err := data.Scan(&mahasiswa.ID, &mahasiswa.Nama, &mahasiswa.Kelas, &mahasiswa.NIM); err != nil {
			log.Fatal(err.Error())

		} else {
			mahasiswaArray = append(mahasiswaArray, mahasiswa)
		}
	}

	return mahasiswaArray
}

func MahasiswaCreate(nama string, kelas string, nim int) bool {
	db := config.DBConnect()
	defer db.Close()

	data, err := db.Query("INSERT INTO mahasiswa (nama, kelas, nim) VALUES ( '" + nama + "','" + kelas + "','" + strconv.Itoa(nim) + "')")
	if err != nil {
		log.Print(err)
		return false
	}

	if data == nil {
		log.Print("data kosong")
		return false
	}

	return true
}

func MahasiswaUpdateNama(nama string) {

}

func MahasiswaDelete(nim int) {

}
