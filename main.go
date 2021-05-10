package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	check_dir()
	/*
		folder_to_watch := "./testfolder/"
		file_to_Check := "testfile"

		filename, modifiedtime := new_file_finder(file_to_Check)
		fmt.Printf("File %s was last modified %s \n", filename, modifiedtime)

		path, modifiedtime := find_new_file(folder_to_watch)
		fmt.Printf("path %s was last modified %s \n", path, modifiedtime)
	*/
}

func check_dir() {
	files, err := ioutil.ReadDir("./testfolder")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), " ", file.ModTime(), " ", file.ModTime().Unix())

	}
}

func find_new_file(path string) (string, time.Time) {
	// get last modified time
	file, err := os.Stat(path)
	if err != nil {
		//fmt.Println("there was an error \n", err)
		//fmt.Print("---------------")
		log.Fatal(err)
	}

	modifiedtime := file.ModTime()
	return path, modifiedtime

}

func new_file_finder(new_file string) (string, time.Time) {

	filename := new_file //file with newest modification date

	// get last modified time
	file, err := os.Stat(filename)
	if err != nil {
		//fmt.Println("there was an error \n", err)
		//fmt.Print("---------------")
		log.Fatal(err)
	}

	modifiedtime := file.ModTime()
	return filename, modifiedtime

}
