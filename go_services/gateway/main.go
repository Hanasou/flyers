package main

import (
	"fmt"

	"github.com/Hanasou/flyers/go_services/common/shutdown"
	"github.com/Hanasou/flyers/go_services/common/utils"
)

func main() {
	fmt.Println("Hello from Gateway")

	utils.HelloFromUtils()

	shutdown.WaitForInterrupt()
}
