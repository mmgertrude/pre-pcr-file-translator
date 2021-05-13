package main

import (
	"os"
	"testing"
	"strconv"
	"github.com/spf13/afero"
)


// A test to check if we crate new folders that does not exist, and not create folders if they do exist
// TODO: Think about changing to table driven testing
func TestCheck_and_create_folders(t *testing.T) {
	// The folders we want to use for testing
	folders := []string{"folderThatShouldExist/", "folderToBeCreated/"}

	// --- Setting up for testing
	// Declaring a backend to use Afero
	appFS := afero.NewMemMapFs()
	// Making the folder that should exist
	err := appFS.MkdirAll(folders[0], os.ModePerm)
	if err != nil {
		t.Fatalf("%error", err)
	}
	folder, err := appFS.Stat(folders[0])
	if err != nil {
		t.Fatalf("%error", err)
	}
	// Recording the time it was created
	createdTime := folder.ModTime().Unix()

	// Using check_and_create_folder to see if it works:
	err = check_and_create_folders(folders, appFS)
	if err != nil {
		t.Fatalf("%error", err)
	}

	// Want to check if folderThatShouldExist/ was not recreated
	folder, err = appFS.Stat(folders[0])
	if err != nil {
		t.Fatalf("%error", err)
	}
	modTime := folder.ModTime().Unix()
	// If the folder now has a different mod time give an error
	if modTime > createdTime{
		t.Fatalf("Error! Folder %s was created even though it already existed", folders[0])
	}

	// Want to check that the folderToBeCreated was created
	_, err = appFS.Stat(folders[1])
	// If the folder was not created give an error:
	if os.IsNotExist(err) {
		t.Fatalf("Error! Folder %s does not exist because it was not created", folders[1])
	}
}


// A test for check if we can discover new files
func TestFile_discoverer(t *testing.T) {
	// The folder
	folder := "test/"

	// --- Setting up for testing
	// Declaring a backend to use Afero
	appFS := afero.NewMemMapFs()
	// Make our test folder
	err := appFS.MkdirAll(folder, os.ModePerm)
	if err != nil {
		t.Fatalf("%error", err)
	}
	// Make some files that we can discover
	for i := 1; i <= 10; i++ {
		err = afero.WriteFile(appFS, folder+strconv.Itoa(i)+".txt", []byte("test"), os.ModePerm)
		if err != nil {
			t.Fatalf("%error", err)
		}
	}

	// Call the function we want to test
	files, err := file_discoverer(folder, appFS)
	if err != nil {
		t.Fatalf("%error", err)
	}

	// Check if the length of the slice is of the length we expect
	if len(files) != 10 {
		t.Fatalf("Error! Wanted 10 files, but %d was created", len(files))
	}
}


// A test to see if we can move files
func TestFile_mover(t *testing.T) {
	// The folders and file
	folders := []string{"oldPath/", "NewPath/"}
	file := "test.txt"

	// --- Setting up for testing
	// Declaring a backend to use Afero
	appFS := afero.NewMemMapFs()
	// Make our test folders
	for _, folder := range folders {
		err := appFS.MkdirAll(folder, os.ModePerm)
		if err != nil {
			t.Fatalf("%error", err)
		}
	}
	// Make a file we can move from the old path
	err := afero.WriteFile(appFS, folders[0]+file, []byte("test"), os.ModePerm)
	if err != nil {
		t.Fatalf("%error", err)
	}

	// Call function we want to test
	err = file_mover(file, folders[0], folders[1], appFS)
	if err != nil {
		t.Fatalf("%error", err)
	}

	// Check if the file still exists in the oldPath
	_, err = appFS.Stat(folders[0]+file)
	if os.IsExist(err) {
		t.Fatalf("Error! %s still exists in %s", file, folders[0])
	}

	// Check if the file was moved to the newPath
	_, err = appFS.Stat(folders[1]+file)
	if os.IsNotExist(err) {
		t.Fatalf("Error! %s was not moved to %s", file, folders[1])
	}
}
