package repository

import (
	"github.com/hiiamtrong/todo-simple/db"
	"github.com/hiiamtrong/todo-simple/models"
)

type TodoRepo struct{}

func (TodoRepo) AddTodo(todo *models.Todo) error {
	if err := db.DB.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (TodoRepo) GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	if err := db.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
