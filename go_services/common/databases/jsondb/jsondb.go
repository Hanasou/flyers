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
	basePath = "/home/rzhang/git_repos/flyers/go_services/common/databases/jsondb/tables/"
)

func (db *JsonDb) GetAll(table string) ([]byte, error) {
	// Get all will retrieve generic elements from the file, or "table" you provide
	fullPath := basePath + table + ".json"
	byteContents, err := utils.OpenFile(fullPath)
	if err != nil {
		log.Println("JsonDb could not open file under listed path: " + fullPath)
		return nil, err
	}

	return byteContents, nil
}

func (db *JsonDb) AddElement(table string) error {
	return nil
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

	// Marshal this data into json bytes
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(methodName + " failed to marshal todos into Json bytes")
		return err
	}

	// Create file
	fullPath := basePath + objectType + ".json"
	file, err := os.Create(fullPath)
	if err != nil {
		log.Println(methodName + " failed to create file")
		return err
	}
	defer file.Close()

	// Write to file
	_, err = file.Write(jsonBytes)
	if err != nil {
		log.Println(methodName + " generateSampleTodos failed to write Json bytes to file")
		return err
	}
	return nil
}
