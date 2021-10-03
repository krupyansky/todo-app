package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/krupyansky/todo-app/pkg/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type TodoList interface {
	Create(userId int, list entity.TodoList) (int, error)
	GetAll(userId int) ([]entity.TodoList, error)
	GetById(userId, listId int) (entity.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input entity.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item entity.TodoItem) (int, error)
	GetAll(userId, listId int) ([]entity.TodoItem, error)
	GetById(userId, itemId int) (entity.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input entity.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
