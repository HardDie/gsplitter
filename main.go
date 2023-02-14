package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
)

type Config struct {
	Count *int
}

func ReadConfig() *Config {
	// Parse config
	cfg := &Config{
		Count: flag.Int("count", 0, "How many files should be in one folder"),
	}
	flag.Parse()

	// Print help if config invalid or empty
	if *cfg.Count < 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}
	return cfg
}

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

func main() {
	cfg := ReadConfig()
	files := ReadFiles()

	// Split files into arrays
	var split [][]string
	for i := 0; i < len(files); i += *cfg.Count {
		split = append(split, files[i:min(i+*cfg.Count, len(files))])
	}

	// Format for folder name
	countOfFolders := len(split)
	printFmt := fmt.Sprintf("%%0%dd", len(strconv.Itoa(countOfFolders)))

	// Create folders and move files
	for i, list := range split {
		// Create folder
		folderName := fmt.Sprintf(printFmt, i)
		err := os.Mkdir(folderName, 0755)
		if err != nil {
			log.Fatalf("error create folder %q: %s\n", folderName, err.Error())
		}

		// Move files into folder
		for _, file := range list {
			err = os.Rename(file, path.Join(folderName, file))
			if err != nil {
				log.Fatalf("error move file %q into folder %s: %s\n", file, folderName, err.Error())
			}
		}
	}

	log.Println("Done!")
}
