package internal

import (
	"fmt"
	"log"
	"os"
)

func ReadFiles() []string {
	// Open current folder
	dir, err := os.ReadDir(".")
	if err != nil {
		log.Fatal("error read dir:", err.Error())
	}

	// Read list of all files in current directory
	var files []string
	for _, file := range dir {
		if file.IsDir() {
			fmt.Println("Directory skipped:", file.Name())
			continue
		}
		files = append(files, file.Name())
	}
	return files
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
