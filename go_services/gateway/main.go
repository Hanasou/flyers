package main

import (
	"fmt"

	"github.com/Hanasou/flyers/go_services/common/shutdown"
)

func teardown() {
	fmt.Println("Bringing down resources")
}

func main() {
	fmt.Println("Hello World")
	shutdown.WaitForInterrupt()
	teardown()
}
