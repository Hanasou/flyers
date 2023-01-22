package memdb

import (
	"github.com/Hanasou/flyers/go_services/common/graph/model"
	"github.com/Hanasou/flyers/go_services/common/utils"
)

// This will be an in-memory key value store that holds some data
// This will mostly be for testing purposes, to test APIs before we connect to a remote database service

type Entry struct {
	data any
}

type Db struct {
	Store map[string]Entry
}

type TodoDb struct {
	Store map[string]*model.Todo
}

type UserDb struct {
	Store map[string]*model.User
}

func NewUserStore() *UserDb {
	newStore := &UserDb{
		map[string]*model.User{},
	}
	return newStore
}

func NewTodoStore() *TodoDb {
	newStore := &TodoDb{
		map[string]*model.Todo{},
	}
	return newStore
}

func UpsertElement(db *Db, key string, data any) {
	db.Store[key] = Entry{
		data,
	}
}

func UpsertTodoElement(db *TodoDb, key string, todo *model.Todo) {
	db.Store[key] = todo
}

func UpsertUserElement(db *UserDb, key string, user *model.User) {
	db.Store[key] = user
}

func PopulateUserStore(db *UserDb) {
	// Generate users
	nUsers := 5
	userIds := []string{}
	userNames := []string{"Bob", "James", "Eric", "Todd", "Howard"}
	for i := 0; i < nUsers; i++ {
		userIds = append(userIds, utils.GenerateId(16))
	}

	for i := 0; i < nUsers; i++ {
		// New user
		newUser := &model.User{
			ID:   userIds[i],
			Name: userNames[i],
		}
		UpsertUserElement(db, newUser.ID, newUser)
	}
}

func PopulateToDoStore(todoDb *TodoDb, userDb *UserDb) {

	// Let's just spray the users into a slice for now
	usersList := []*model.User{}
	for _, v := range userDb.Store {
		usersList = append(usersList, v)
	}

	// Generate todo objects
	nToDo := 10
	toDoIds := []string{}
	for i := 0; i < nToDo; i++ {
		toDoIds = append(toDoIds, utils.GenerateId(16))
	}
	toDoText := []string{}
	for i := 0; i < nToDo; i++ {
		toDoText = append(toDoText, utils.GenerateId(32))
	}

	for i := 0; i < nToDo; i++ {
		newTodo := &model.Todo{
			ID:   toDoIds[i],
			Text: toDoText[i],
			Done: utils.GenerateNumber(1, 2) == 1,
			User: usersList[utils.GenerateNumber(1, 5)],
		}
		UpsertTodoElement(todoDb, newTodo.ID, newTodo)
	}
}
