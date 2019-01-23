package mMahasiswa

import (
	"fmt"
	"log"
	"strconv"

	"gotest.com/go-sandbox/structs"

	"gotest.com/go-sandbox/config"
)

//MahasiswaAll -> Get all mahasiswa
func MahasiswaAll() []structs.Mahasiswa {
	var mahasiswa structs.Mahasiswa
	var mahasiswaArray []structs.Mahasiswa

	rows, err := config.DG.Model(&mahasiswa).Select("*").Rows()
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		if err := rows.Scan(
			&mahasiswa.ID,
			&mahasiswa.Nama,
			&mahasiswa.Kelas,
			&mahasiswa.NIM,
			&mahasiswa.CreatedAt,
			&mahasiswa.UpdatedAt,
			&mahasiswa.DeletedAt,
		); err == nil {
			mahasiswaArray = append(mahasiswaArray, mahasiswa)
		} else {
			fmt.Println(err.Error())
		}
	}

	return mahasiswaArray
}

func MahasiswaCreate(nama string, kelas string, nim int) bool {
	mahasiswa := structs.Mahasiswa{
		Nama:  nama,
		Kelas: kelas,
		NIM:   nim,
	}

	config.DG.AutoMigrate(&structs.Mahasiswa{})

	data := config.DG.Create(&mahasiswa)
	if data.Error != nil {
		log.Print(data.Error)
		return false
	}

	return true
}

func MahasiswaUpdateNama(nim int, nama string) bool {
	mahasiswa := structs.Mahasiswa{
		Nama: nama,
	}

	data := config.DG.Model(&mahasiswa).Where("nim = " + strconv.Itoa(nim)).Update(&mahasiswa)
	if data.Error != nil {
		log.Print(data.Error)
		return false
	}

	return true
}

func MahasiswaDelete(nim int) bool {
	var mahasiswa structs.Mahasiswa

	data := config.DG.Model(&mahasiswa).Where("nim = " + strconv.Itoa(nim)).Delete(&mahasiswa)
	if data.Error != nil {
		log.Print(data.Error)
		return false
	}

	return true
}
