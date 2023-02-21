package jsondb

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/Hanasou/flyers/go_services/common/graph/model"
	"github.com/Hanasou/flyers/go_services/common/utils"
)

// For this, we're going to store data in Json format in .json files
// The pure in memory database wasn't exactly working out since it's hard to make it generic with static types

type JsonDb struct {
}

const (
	basePath = "./tables/"
)

func (db *JsonDb) GetAll(table string) ([]map[string]interface{}, error) {
	// Get all will retrieve generic elements from the file, or "table" you provide
	fullPath := basePath + table + ".json"
	byteContents, err := utils.OpenFile(fullPath)
	if err != nil {
		log.Println("JsonDb could not open file under listed path: " + fullPath)
		return nil, err
	}
	objects := []map[string]interface{}{}
	err = json.Unmarshal(byteContents, &objects)
	if err != nil {
		log.Println("JsonDb could not unmarshal data successfully into slice of maps")
		return nil, err
	}

	return objects, nil
}

func GenerateSampleData(objectType string, amount int) error {
	methodName := "generateSample" + objectType

	// Create a list of generic data.
	// We're going to compromise type safety here because it's hard
	data := []any{}

	// Populate data with objects depending on what we passed in
	for i := 0; i < amount; i++ {
		iString := strconv.Itoa(i)
		if objectType == "todo" {
			newTodo := &model.Todo{
				ID:   utils.GenerateId(16),
				Text: "Text for ToDo Number" + iString,
				Done: false,
				User: nil,
			}
			data = append(data, newTodo)
		} else if objectType == "user" {
			newUser := &model.User{
				ID:   utils.GenerateId(16),
				Name: "NewName " + iString,
			}
			data = append(data, newUser)
		} else {
			return errors.New(methodName + " could not find provided object type")
		}
	}

	return nil
}

// This function will write a file containing a list of ToDo objects
func GenerateSampleTodos() {
	methodName := "generateSampleTodos"
	// Create a list of ToDos
	todos := []*model.Todo{}

	for i := 0; i < 10; i++ {
		newTodo := &model.Todo{
			ID:   utils.GenerateId(16),
			Text: "Text for ToDo Number" + strconv.Itoa(i),
			Done: false,
			User: nil,
		}
		todos = append(todos, newTodo)
	}

	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		log.Println(methodName + " failed to marshal todos into Json bytes")
		return
	}

	fullPath := basePath + "todo.json"
	file, err := os.Create(fullPath)
	if err != nil {
		log.Println(methodName + " failed to create file")
		return
	}
	defer file.Close()

	_, err = file.Write(jsonBytes)
	if err != nil {
		log.Println(methodName + " generateSampleTodos failed to write Json bytes to file")
		return
	}

	log.Println("Todos generated")
}
