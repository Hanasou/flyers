package main

import (
	"github.com/Hanasou/flyers/go_services/common/shutdown"
)

func main() {
	shutdown.WaitForInterrupt()
}
