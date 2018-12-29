package main

import (
	"bufio"
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
var DocIndexed int

/*
CreateIndex : creates the index :P.
*/
func CreateIndex(pathtodir string) error {
	fmt.Println("Building the inverted index.")
	InvertedIndex = make(map[string][]int)
	dirf, err := os.Open(pathtodir)
	if err != nil {
		return err
	}
	files, err := dirf.Readdir(-1)
	if err != nil {
		return err
	}
	if len(files) <= 0 {
		fmt.Println("Empty dataset provided, building a empty index.")
		return nil
	}
	for _, file := range files {
		if !file.IsDir() {
			var fipath = filepath.Join(pathtodir, file.Name())
			err := indexFile(fipath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func indexFile(fipath string) error {
	file, err := os.Open(fipath)
	if err != nil {
		return err
	}
	DocIndexed++
	Documents = append(Documents, Document{file.Name(), fipath})
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)
wordLoop:
	for fileScanner.Scan() {
		word := Normalize(fileScanner.Text())
		fmt.Println(word)
		pl := InvertedIndex[word]
		for _, docIndex := range pl {
			if docIndex == DocIndexed {
				continue wordLoop
			}
		}
		InvertedIndex[word] = append(pl, DocIndexed)
	}
	return nil
}
