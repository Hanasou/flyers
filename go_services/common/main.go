package main

import (
	"fmt"

	"github.com/Hanasou/flyers/go_services/common/memdb"
)

func main() {
	fmt.Println(memdb.NewKv().Store)
	memdb.Hello()
}
