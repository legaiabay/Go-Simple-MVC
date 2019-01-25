package structs

import (
	"time"
)

//Mahasiswa -> struct
type TodoGroup struct {
	Todo      []Todo     `json:"todo" gorm:"foreignkey:IDGroup"`
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (TodoGroup) TableName() string {
	return "4su_todo_groups"
}
