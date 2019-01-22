package mMahasiswa

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

	/*
		db := config.DBConnectLama()
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
	*/

	rows, err := config.DG.Table("mahasiswa").Select("id, nama, kelas, nim").Rows()
	defer rows.Close()

	if err != nil {
		log.Print(err)
		return nil
	}

	for rows.Next() {
		if err := rows.Scan(&mahasiswa.ID, &mahasiswa.Nama, &mahasiswa.Kelas, &mahasiswa.NIM); err != nil {
			log.Fatal(err.Error())
		} else {
			mahasiswaArray = append(mahasiswaArray, mahasiswa)
		}
	}

	return mahasiswaArray
}

func MahasiswaCreate(nama string, kelas string, nim int) bool {
	/*
		db := config.DBConnectLama()
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
	*/

	mahasiswa := structs.Mahasiswa{
		Nama:  nama,
		Kelas: kelas,
		NIM:   nim,
	}

	data := config.DG.Table("mahasiswa").Create(&mahasiswa)
	if data == nil {
		log.Print("data kosong")
		return false
	}

	return true
}

func MahasiswaUpdateNama(nim int, nama string) bool {
	/*
		db := config.DBConnectLama()
		defer db.Close()

		data, err := db.Query("UPDATE mahasiswa SET nama = '" + nama + "' WHERE nim = '" + strconv.Itoa(nim) + "'")
		if err != nil {
			log.Print(err)
			return false
		}

		if data == nil {
			log.Print("data kosong")
			return false
		}
	*/

	mahasiswa := structs.Mahasiswa{Nama: nama}

	data := config.DG.Table("mahasiswa").Where("nim = " + strconv.Itoa(nim)).Update(&mahasiswa)
	if data == nil {
		log.Print("data kosong")
		return false
	}

	return true
}

func MahasiswaDelete(nim int) bool {
	/*
		db := config.DBConnectLama()
		defer db.Close()

		data, err := db.Query("DELETE FROM mahasiswa WHERE nim = '" + strconv.Itoa(nim) + "'")
		if err != nil {
			log.Print(err)
			return false
		}

		if data == nil {
			log.Print("data kosong")
			return false
		}

		return true
	*/

	data := config.DG.Table("mahasiswa").Where("nim = " + strconv.Itoa(nim)).Delete(structs.Mahasiswa{})
	if data == nil {
		log.Print("data kosong")
		return false
	}

	return true
}
