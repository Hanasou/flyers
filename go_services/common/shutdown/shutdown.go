package shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func WaitForInterrupt() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	fmt.Println("Shutting Down")
}
