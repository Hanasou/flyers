package utils

import (
	"log"
	"os"
)

// This should provide functions for reading the contents of files

func OpenFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Println("Cannot open file")
		return nil, err
	}
	defer file.Close()

	// Create a byte slice with enough capacity to hold the file
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println("Cannot get stats for file")
		return nil, err
	}
	buffer := make([]byte, fileInfo.Size())

	// Read contents of file into buffer
	_, err = file.Read(buffer)
	if err != nil {
		log.Println("Cannot put contents of file into buffer")
		return nil, err
	}

	return buffer, nil
}
