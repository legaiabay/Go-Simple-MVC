package structs

import (
	"time"
)

//Mahasiswa -> struct
type Todo struct {
	ID        int        `json:"id"`
	Author    string     `json:"author"`
	Note      string     `json:"note"`
	IDGroup   int        `json:"id_group"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (Todo) TableName() string {
	return "4su_todos"
}
