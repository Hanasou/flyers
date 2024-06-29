package main

import (
	"fmt"

	"github.com/Hanasou/flyers/go_services/common/graphql"
)

func main() {
	fmt.Println("Hello from Gateway")
	graphql.StartGraphqlServer()
}
