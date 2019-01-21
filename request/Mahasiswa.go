package request

type Mahasiswa struct {
	Nama  string `json:nama`
	Kelas string `json:kelas`
	NIM   int    `json:nim`
}
