package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/caarlos0/env"
	"github.com/spf13/afero"
)

// Setting environmental variables
type session struct {
	InputFolder     string `env:"INPUT_FOLDER" envDefault:"testInput/"`
	OutputFolder    string `env:"OUTPUT_FOLDER" envDefault:"testInput/testOutput/"`
	ProcessedFolder string `env:"PROCESSED_FOLDER" envDefault:"testInput/testProcessed/"`
	ErrorFolder     string `env:"ERROR_FOLDER" envDefault:"testInput/testError/"`
	ApiCallUrl      string `env:"API_CALL_URL" envDefault:"https://u700courier.test.mgmlab.net/analysis_by_name/Q"`
	appFS			afero.Fs
}

func main() {
	// Setting to use the environmental variables
	cfg := session{}
	env.Parse(&cfg)

	// Setting up for afero
	cfg.appFS = afero.NewOsFs()

	// If folders does not exist we should make them
	folders := []string{cfg.InputFolder, cfg.OutputFolder, cfg.ProcessedFolder, cfg.ErrorFolder}
	err := check_and_create_folders(folders, cfg.appFS)
	if err != nil {
		log.Fatal(err)
	}

	// Getting slice with files in the folder
	fileSlice, err := file_discoverer(cfg.InputFolder, cfg.appFS)
	if err != nil {
		log.Fatal(err)
	}

	// Get data from the restAPI
	responseData, err := get_data(cfg.ApiCallUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Loop over files to process them
	var processingError error
	for _, file := range fileSlice {

		// Move file according to if we had an error or not
		if processingError == nil {
			err = file_mover(file, cfg.InputFolder, cfg.ProcessedFolder, cfg.appFS)
		} else {
			err = file_mover(file, cfg.InputFolder, cfg.ErrorFolder, cfg.appFS)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Checks if a folder exists and if it does not create one
func check_and_create_folders(folders []string, appFS afero.Fs) error {
	var err error
	// For each folder in the list
	for _, folder := range folders {
		_, err := appFS.Stat(folder)
		// Check if it exists
		if os.IsNotExist(err) {
			// If not make the folder
			err = appFS.MkdirAll(folder, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// Returning a slice with the files in a folder
func file_discoverer(InputFolder string, appFS afero.Fs) ([]string, error) {
	// Making an empty slice
	var fileSlice []string
	// Reading all the files in a folder
	files, err := afero.ReadDir(appFS, InputFolder)
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

// Make map with analysis data for MGM
func get_data(ApiCallUrl string) ([]byte, error) {
	var responseData []byte

	// Requesting data for MGM
	response, err := http.Get(ApiCallUrl)
	if err != nil {
		return responseData, err
	}
	// Just getting the body of the response
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return responseData, err
	}
	return responseData, err
}

// Move file from one folder to another folder
func file_mover(file string, oldPath string, newPath string, appFS afero.Fs) error {
	err := appFS.Rename(oldPath+file, newPath+file)
	return err
}
