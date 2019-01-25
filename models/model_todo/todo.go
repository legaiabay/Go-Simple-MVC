package model_todo

import (
	"strconv"

	"gotest.com/go-sandbox/structs"

	"gotest.com/go-sandbox/config"
)

//MahasiswaAll -> Get all mahasiswa
func TodoAll() ([]structs.Todo, string) {
	var todo structs.Todo
	var todos []structs.Todo

	query, err := config.DB.Model(&todo).Rows()
	defer query.Close()
	if err != nil {
		return todos, err.Error()
	}

	for query.Next() {
		if err := query.Scan(
			&todo.ID,
			&todo.Author,
			&todo.Note,
			&todo.IDGroup,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.DeletedAt,
		); err == nil {
			todos = append(todos, todo)
		} else {
			return todos, err.Error()
		}
	}

	return todos, "success"
}

//MahasiswaGetByNim -> Get mahasiswa by NIM
func TodoGetById(id int) (structs.Todo, string) {
	todo := structs.Todo{ID: id}

	query := config.DB.Find(&todo, &todo).Row()

	if data := query.Scan(
		&todo.ID,
		&todo.Author,
		&todo.Note,
		&todo.IDGroup,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.DeletedAt,
	); data != nil {
		return todo, "No todo with ID " + strconv.Itoa(id)
	}

	return todo, "success"
}

func TodoCreate(author string, note string, id_group int) string {
	todo := structs.Todo{
		Author:  author,
		Note:    note,
		IDGroup: id_group,
	}

	where := structs.TodoGroup{
		ID: id_group,
	}

	query := config.DB.Where(&where).Create(&todo)
	if query.Error != nil {
		return query.Error.Error()
	}

	return "success"
}

func TodoUpdateNote(id int, author string, note string) string {
	todo := structs.Todo{
		Author: author,
		Note:   note,
	}
	where := structs.Todo{ID: id}

	query := config.DB.Model(&todo).Where(&where).Update(&todo)
	if query.RowsAffected == 0 {
		return "No note with ID " + strconv.Itoa(id)
	}

	return "success"
}

func TodoDelete(id int) string {
	todo := structs.Todo{ID: id}

	query := config.DB.Delete(&todo, &todo)
	if query.RowsAffected == 0 {
		return "No todo with ID " + strconv.Itoa(id)
	}

	return "success"
}
