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

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error get current directory: %s\n", err.Error())
	}
	for folderName, list := range folders {
		moveFiles(wd, folderName, list)
	}
}

func SplitByCount(files []string, count, offset int) {
	// Split files into arrays
	var split [][]string
	for i := 0; i < len(files); i += count {
		split = append(split, files[i:min(i+count, len(files))])
	}

	// Format for folder name
	countOfFolders := len(split)
	if offset > 0 {
		countOfFolders += offset - 1
	}
	printFmt := fmt.Sprintf("%%0%dd", len(strconv.Itoa(countOfFolders)))

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error get current directory: %s\n", err.Error())
	}
	for i, list := range split {
		folderName := fmt.Sprintf(printFmt, i+offset)
		moveFiles(wd, folderName, list)
	}
}

func SplitByDate(files []os.FileInfo) {
	folders := make(map[string][]string)

	// Split files by date
	for _, file := range files {
		date := file.ModTime().Format("2006-01-02")
		folders[date] = append(folders[date], file.Name())
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error get current directory: %s\n", err.Error())
	}
	for folderName, list := range folders {
		moveFiles(wd, folderName, list)
	}
}

func SplitByFirstLetter(files []string) {
	const Unknown = "unknown"
	folders := make(map[string][]string)

	// Split files by first letter
	for _, file := range files {
		if file == "" {
			folders[Unknown] = append(folders[Unknown], file)
			continue
		}
		firstLetter := strings.ToUpper(string([]rune(file)[0]))
		if firstLetter == "" {
			folders[Unknown] = append(folders[Unknown], file)
			continue
		}
		folders[firstLetter] = append(folders[firstLetter], file)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error get current directory: %s\n", err.Error())
	}
	for folderName, list := range folders {
		moveFiles(wd, folderName, list)
	}
}

func moveFiles(wd, folderName string, files []string) {
	exist, err := IsFolderExists(folderName)
	if err != nil {
		log.Fatalf("error check folder %q: %s\n", folderName, err.Error())
	}

	tmpFolderName := folderName
	if !exist {
		tmpFolderName, err = os.MkdirTemp(wd, folderName)
		if err != nil {
			log.Fatalf("error create tmp folder %q: %s\n", folderName, err.Error())
		}
	}

	// Move files into folder
	for _, file := range files {
		err = os.Rename(file, path.Join(tmpFolderName, file))
		if err != nil {
			log.Fatalf("error move file %q into folder %s: %s\n", file, tmpFolderName, err.Error())
		}
	}

	if !exist {
		// Rename tmp folder into result folder
		err = os.Rename(tmpFolderName, folderName)
		if err != nil {
			log.Fatalf("error rename tmp folder %q into folder %s: %s\n", tmpFolderName, folderName, err.Error())
		}
	}
}
