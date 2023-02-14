package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type Config struct {
	Count       *int
	ByExtension *bool
}

func ReadConfig() *Config {
	// Parse config
	cfg := &Config{
		Count:       flag.Int("count", 0, "How many files should be in one folder"),
		ByExtension: flag.Bool("by_ext", false, "Split files by extension"),
	}
	flag.Parse()

	// One of the option should be passed
	if *cfg.Count < 1 && *cfg.ByExtension == false {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Do not allow both arguments to be used at the same time
	if *cfg.Count >= 1 && *cfg.ByExtension {
		log.Println("Choose only one option:")
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

func SplitByExt(files []string) {
	const Unknown = "unknown"
	folders := make(map[string][]string)

	// Split files by extensions
	for _, file := range files {
		ext := strings.ToLower(path.Ext(file))
		if ext == "" {
			folders[Unknown] = append(folders[Unknown], file)
			continue
		}
		ext = ext[1:]
		folders[ext] = append(folders[ext], file)
	}

	// Create folders and move files
	for folderName, list := range folders {
		// Create folder
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
}

func SplitByCount(files []string, cfg *Config) {
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
}

func main() {
	cfg := ReadConfig()
	files := ReadFiles()

	if *cfg.Count > 0 {
		SplitByCount(files, cfg)
	}
	if *cfg.ByExtension {
		SplitByExt(files)
	}

	log.Println("Done!")
}
