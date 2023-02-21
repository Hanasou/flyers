package databases

import (
	"errors"

	"github.com/Hanasou/flyers/go_services/common/databases/jsondb"
)

type Driver interface {
	GetAll(table string) ([]map[string]interface{}, error)
}

func NewDriver(driverType string) (Driver, error) {
	if driverType == "json" {
		return &jsondb.JsonDb{}, nil
	}
	return nil, errors.New("database not supported")
}
