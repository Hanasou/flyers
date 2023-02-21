package main

import (
	"github.com/Hanasou/flyers/go_services/common/graphql"
	"github.com/Hanasou/flyers/go_services/common/shutdown"
)

func main() {
	graphql.StartGraphqlServer()
	shutdown.WaitForInterrupt()
}
