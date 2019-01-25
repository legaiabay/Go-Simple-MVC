package model_todoGroup

import (
	"strconv"

	"gotest.com/go-sandbox/config"
	"gotest.com/go-sandbox/structs"
)

func TodoGroupAll() ([]structs.TodoGroup, string) {
	var todo []structs.Todo
	var todoGroup structs.TodoGroup
	var todoGroups []structs.TodoGroup

	config.DB.Find(&todoGroup)
	query, err := config.DB.Model(&todoGroup).Related(&todo).Rows()
	defer query.Close()
	if err != nil {
		return todoGroups, err.Error()
	}

	for query.Next() {
		if err := query.Scan(
			&todoGroup.ID,
			&todoGroup.Title,
			&todoGroup.CreatedAt,
			&todoGroup.UpdatedAt,
			&todoGroup.DeletedAt,
		); err == nil {
			todoGroups = append(todoGroups, todoGroup)
		} else {
			return todoGroups, err.Error()
		}
	}

	return todoGroups, "success"
}

func TodoGroupGetById(id int) (structs.TodoGroup, string) {
	var todo []structs.Todo
	var todoGroup structs.TodoGroup

	where := structs.TodoGroup{ID: id}

	config.DB.Find(&todoGroup) //Cari group nya dulu
	query := config.DB.Model(&todoGroup).Related(&todo).Where(&where).Row()

	if query := query.Scan(
		&todoGroup.ID,
		&todoGroup.Title,
		&todoGroup.CreatedAt,
		&todoGroup.UpdatedAt,
		&todoGroup.DeletedAt,
	); query != nil {
		return todoGroup, query.Error()
	}

	return todoGroup, "success"
}

func TodoGroupCreate(title string) string {
	todoGroup := structs.TodoGroup{
		Title: title,
	}

	query := config.DB.Create(&todoGroup)
	if query.Error != nil {
		return query.Error.Error()
	}

	return "success"
}

func TodoGroupUpdateTitle(id int, title string) string {
	todoGroup := structs.TodoGroup{Title: title}
	where := structs.TodoGroup{ID: id}

	query := config.DB.Model(&todoGroup).Where(&where).Update(&todoGroup)
	if query.RowsAffected == 0 {
		return "No todoGroup with ID " + strconv.Itoa(id)
	}

	return "success"
}

func TodoGroupDelete(id int) string {
	todoGroup := structs.TodoGroup{ID: id}
	todo := structs.Todo{IDGroup: id}

	query := config.DB.Delete(&todoGroup, &todoGroup)
	if query.RowsAffected == 0 {
		return "No todoGroup with ID " + strconv.Itoa(id)
	} else {
		config.DB.Delete(&todo, &todo) //Deletes todo in group too
	}

	return "success"
}
