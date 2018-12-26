package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Document struct {
	filename string
	filepath string
}

var InvertedIndex map[string][]int
var Documents []Document

func CreateIndex(pathtodir string) (bool, error) {
	fmt.Println("Building the inverted index.")
	dirf, err := os.Open(pathtodir)
	if err != nil {
		return false, err
	}
	files, err := dirf.Readdir(-1)
	if err != nil {
		return false, err
	}
	if len(files) <= 0 {
		fmt.Println("Empty dataset provided, building a empty index.")
		return true, nil
	}
	//	indexed := 0
	for _, file := range files {
		if !file.IsDir() {
			var fipath = filepath.Join(pathtodir, file.Name())
			indexFile(fipath)
		}
	}

	return true, nil
}
