package main

import (
	"fmt"

	"github.com/Hanasou/flyers/go_services/common/memdb"
	"github.com/Hanasou/flyers/go_services/common/shutdown"
)

func main() {
	fmt.Println(memdb.NewTodoStore().Store)
	shutdown.WaitForInterrupt()
}
