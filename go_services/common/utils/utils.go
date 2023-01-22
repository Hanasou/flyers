package utils

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
)

// Utilities package
func HelloFromUtils() {
	fmt.Println("Hello from Utils")
}

func HelloFromUtils2() {
	fmt.Println("Hello from Utils 2")
}

// New method to test something else out
func NewHelloFromUtils() {
	fmt.Println("New Hello from Utils Function")
}

// GenericErrorHandle will simply panic on errors
func PanicErrorHandle(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

// GenerateNumber
func GenerateNumber(min int, max int) int {
	return rand.Intn(max-min) + min
}

// GenerateId will generate a string of random characters
func GenerateId(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	PanicErrorHandle(err)

	return hex.EncodeToString(b)
}
