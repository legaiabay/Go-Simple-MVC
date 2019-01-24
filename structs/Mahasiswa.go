package structs

import "time"

//Mahasiswa -> struct
type Mahasiswa struct {
	ID        int        `json:"id"`
	Nama      string     `json:"nama"`
	Kelas     string     `json:"kelas"`
	NIM       int        `json:"nim"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (Mahasiswa) TableName() string {
	return "4sa_mahasiswa"
}
