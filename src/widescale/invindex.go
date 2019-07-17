package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

type Document struct {
	FileName string `json:"file_name"`
	Filepath string `json:"file_path"`
}

// InvertedIndex data structure mapping from tokens(words) to documents.
var InvertedIndex map[string][]int

// Documents is a splice containg information about of all the documents.
var Documents []Document

// DocIndexed stores count of total documents indexed.
var DocIndexed int

/*
CreateIndex : creates the index :P.
*/
func CreateIndex(pathtodir string) error {
	log.Println("Building the inverted index.")
	InvertedIndex = make(map[string][]int)
	dirf, err := os.Open(pathtodir)
	if err != nil {
		return err
	}
	files, err := dirf.Readdir(-1)
	if err != nil {
		log.Println("An error occured while reading the directory.")
		return err
	}
	if len(files) <= 0 {
		log.Println("Empty dataset provided, building a empty index.")
		return nil
	}
	for _, file := range files {
		if !file.IsDir() {
			var fipath = filepath.Join(pathtodir, file.Name())
			err := indexFile(fipath)
			if err != nil {
				log.Println("Unable to index", fipath, ".")
				return err
			}
		}
	}
	log.Println("Built an index of", DocIndexed, "documents.")
	return nil
}

/*
Fetches the file and adds its words to the inverted index.
*/
func indexFile(fipath string) error {
	file, err := os.Open(fipath)
	if err != nil {
		return err
	}
	DocIndexed++
	Documents = append(Documents, Document{filepath.Base(fipath), fipath})
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)
wordLoop:
	for fileScanner.Scan() {
		word := Normalize(fileScanner.Text())
		pl := InvertedIndex[word]
		for _, docIndex := range pl {
			if docIndex == DocIndexed-1 {
				continue wordLoop
			}
		}
		InvertedIndex[word] = append(pl, DocIndexed-1)
	}
	return nil
}
