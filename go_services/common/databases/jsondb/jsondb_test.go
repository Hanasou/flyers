package jsondb_test

import (
	"log"
	"testing"

	"github.com/Hanasou/flyers/go_services/common/databases"
	"github.com/Hanasou/flyers/go_services/common/databases/jsondb"
)

func TestJsonUnmarshal(t *testing.T) {
	jsondb.GenerateSampleTodos()
}

func TestGenerateTestData(t *testing.T) {
	jsondb.GenerateSampleData("todo", 10)
	jsondb.GenerateSampleData("user", 10)
}

func TestGetAll(t *testing.T) {
	jsonDb, err := databases.NewDriver("json")
	if err != nil {
		log.Println("JsonDb test could not create JsonDb driver")
		return
	}

	contents, err := jsonDb.GetAll("todo")
	if err != nil {
		log.Println("Could not provided Json file")
		return
	}

	for _, v := range contents {
		log.Println(v)
	}
}
