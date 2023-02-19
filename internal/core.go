package internal

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

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
			if !os.IsExist(err) {
				log.Fatalf("error create folder %q: %s\n", folderName, err.Error())
			}
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

func SplitByCount(files []string, count int) {
	// Split files into arrays
	var split [][]string
	for i := 0; i < len(files); i += count {
		split = append(split, files[i:min(i+count, len(files))])
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
