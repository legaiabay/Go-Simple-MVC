package structs

type Mahasiswa struct {
	ID    int    `json:id`
	Nama  string `json:nama`
	Kelas string `json:kelas`
	NIM   int    `json:nim`
}
