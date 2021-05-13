package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/caarlos0/env"
)

// Setting environmental variables
type session struct {
	InputFolder     string `env:"INPUT_FOLDER" envDefault:"testInput/"`
	OutputFolder    string `env:"OUTPUT_FOLDER" envDefault:"testInput/testOutput/"`
	ProcessedFolder string `env:"PROCESSED_FOLDER" envDefault:"testInput/testProcessed/"`
	ErrorFolder     string `env:"ERROR_FOLDER" envDefault:"testInput/testError/"`
	ApiCallUrl      string `env:"API_CALL_URL" envDefault:"u700courier.test.mgmlab.net/analysis_by_name/Q"`
}

func main() {
	cfg := session{}
	env.Parse(&cfg)

	// If folders does not exist we should make them
	folders := []string{cfg.InputFolder, cfg.OutputFolder, cfg.ProcessedFolder, cfg.ErrorFolder}
	err := check_and_create_folders(folders)
	if err != nil {
		log.Fatal(err)
	}

	// Getting slice with files in the folder
	fileSlice, err := file_discoverer(cfg.InputFolder)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileSlice)

	// Loop over files to process them
	var processingError error
	for _, file := range fileSlice {

		// Move file according to if we had an error or not
		if processingError == nil {
			err = file_mover(file, cfg.InputFolder, cfg.ProcessedFolder)
		} else {
			err = file_mover(file, cfg.InputFolder, cfg.ErrorFolder)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Checks if a folder exists and if it does not create one
func check_and_create_folders(folders []string) error {
	var err error
	// For each folder in the list
	for _, folder := range folders {
		_, err := os.Stat(folder)
		// Check if it exists
		if os.IsNotExist(err) {
			// If not make the folder
			err = os.Mkdir(folder, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// Returning a slice with the files in a folder
func file_discoverer(InputFolder string) ([]string, error) {
	// Making an empty slice
	var fileSlice []string
	// Reading all the files in a folder
	files, err := ioutil.ReadDir(InputFolder)
	if err != nil {
		return fileSlice, err
	}
	// Looping through files
	for _, file := range files {
		// If they are not a folder add them to the list of files
		if !file.IsDir() {
			fileSlice = append(fileSlice, file.Name())
		}
	}
	return fileSlice, err
}

// Move file from one folder to another folder
func file_mover(file string, oldPath string, newPath string) error {
	err := os.Rename(oldPath+"/"+file, newPath+"/"+file)
	return err
}
