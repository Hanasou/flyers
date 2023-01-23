package memdb

import (
	"encoding/json"

	"github.com/Hanasou/flyers/go_services/common/utils"
)

// This will be an in-memory key value store that holds some data
// This will mostly be for testing purposes, to test APIs before we connect to a remote database service

type database struct {
	tables []*table
}

type table struct {
	schema interface{}
	rows   map[string]*row
}

type row struct {
	data []byte
}

func CreateDatabase() *database {
	return &database{
		tables: []*table{},
	}
}

func CreateTable(schema interface{}) *table {
	return &table{
		schema: schema,
		rows:   map[string]*row{},
	}
}

func (db *database) AddTable(t *table) {
	db.tables = append(db.tables, t)
}

func CreateRow(data interface{}) (*row, string) {
	jsonMap := map[string]interface{}{}
	jsonData, err := json.Marshal(data)
	utils.PanicErrorHandle(err, "Error marshalling JSON")
	err = json.Unmarshal(jsonData, &jsonMap)
	utils.PanicErrorHandle(err, "Error unmarshalling JSON")

	// If key exists in map
	rowId := ""
	if v, ok := jsonMap["ID"]; ok {
		if s, ok := v.(string); ok {
			rowId = s
		} else {
			rowId = utils.GenerateId(16)
		}
	} else {
		rowId = utils.GenerateId(16)
	}
	return &row{jsonData}, rowId
}

func (t *table) UpsertRow(data interface{}) {
	row, rowId := CreateRow(data)
	t.rows[rowId] = row
}
