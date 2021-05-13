package main

import (
	"os"
	"testing"

	"github.com/spf13/afero"
)

// A test to check if we can discover new files
func TestCheck_and_create_folders(t *testing.T) {
	// Declaring a backend to use Afero
	appFS := afero.NewMemMapFs()
	folders := []string{"testfolder1/", "testfolder2/"}

	//setting up system
	//making a new folder
	appFS.MkdirAll(folders[0], os.ModePerm)
	//using check_and_create_folder to see if it works:
	err := check_and_create_folders(folders)
	if err != nil {
		t.Fatalf("%error", err)
	}
	for _, folder := range folders {
		_, err := appFS.Stat(folder)
		//if the folder was not created give an error:
		if os.IsNotExist(err) {
			t.Fatalf("folder %s was not created", folder)
		}
	}

}

// Create a test file in a folder and check if the program notices
