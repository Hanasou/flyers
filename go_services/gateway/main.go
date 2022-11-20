package main

import (
	"fmt"

	"github.com/Hanasou/flyers/go_services/common/shutdown"
	"github.com/Hanasou/flyers/go_services/common/utils"

	"github.com/Hanasou/flyers/go_services/common/changeme"
)

func main() {
	fmt.Println("Hello from Gateway")

	utils.HelloFromUtils()
	utils.HelloFromUtils2()

	changeme.TestFunction()

	shutdown.WaitForInterrupt()
}
