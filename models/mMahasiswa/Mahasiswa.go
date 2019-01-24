package mMahasiswa

import (
	"strconv"

	"gotest.com/go-sandbox/structs"

	"gotest.com/go-sandbox/config"
)

//MahasiswaAll -> Get all mahasiswa
func MahasiswaAll() ([]structs.Mahasiswa, string) {
	var mahasiswa structs.Mahasiswa
	var mahasiswas []structs.Mahasiswa

	query, err := config.DB.Model(&mahasiswa).Rows()
	defer query.Close()
	if err != nil {
		return mahasiswas, err.Error()
	}

	for query.Next() {
		if err := query.Scan(
			&mahasiswa.ID,
			&mahasiswa.Nama,
			&mahasiswa.Kelas,
			&mahasiswa.NIM,
			&mahasiswa.CreatedAt,
			&mahasiswa.UpdatedAt,
			&mahasiswa.DeletedAt,
		); err == nil {
			mahasiswas = append(mahasiswas, mahasiswa)
		} else {
			return mahasiswas, err.Error()
		}
	}

	return mahasiswas, "success"
}

//MahasiswaGetByNim -> Get mahasiswa by NIM
func MahasiswaGetByNim(nim int) (structs.Mahasiswa, string) {
	mahasiswa := structs.Mahasiswa{NIM: nim}

	query := config.DB.Find(&mahasiswa, &mahasiswa).Row()

	if data := query.Scan(
		&mahasiswa.ID,
		&mahasiswa.Nama,
		&mahasiswa.Kelas,
		&mahasiswa.NIM,
		&mahasiswa.CreatedAt,
		&mahasiswa.UpdatedAt,
		&mahasiswa.DeletedAt,
	); data != nil {
		return mahasiswa, "Can't find mahasiswa with NIM " + strconv.Itoa(nim)
	}

	return mahasiswa, "success"
}

func MahasiswaCreate(nama string, kelas string, nim int) string {
	mahasiswa := structs.Mahasiswa{
		Nama:  nama,
		Kelas: kelas,
		NIM:   nim,
	}

	query := config.DB.Create(&mahasiswa)
	if query.Error != nil {
		return query.Error.Error()
	}

	return "success"
}

func MahasiswaUpdateNama(nim int, nama string) string {
	mahasiswa := structs.Mahasiswa{Nama: nama}
	where := structs.Mahasiswa{NIM: nim}

	query := config.DB.Model(&mahasiswa).Where(&where).Update(&mahasiswa)
	if query.RowsAffected == 0 {
		return "Can't find mahasiswa with NIM " + strconv.Itoa(nim)
	}

	return "success"
}

func MahasiswaDelete(nim int) string {
	mahasiswa := structs.Mahasiswa{NIM: nim}

	query := config.DB.Delete(&mahasiswa, &mahasiswa)
	if query.RowsAffected == 0 {
		return "Can't find mahasiswa with NIM " + strconv.Itoa(nim)
	}

	return "success"
}
