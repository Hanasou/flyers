package main

import (
	"fmt"

	"github.com/Hanasou/flyers/go_services/common/shutdown"
)

func main() {
	fmt.Println("Hello from Gateway")

	shutdown.WaitForInterrupt()
}
