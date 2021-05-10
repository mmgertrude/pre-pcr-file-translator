import "testing"

func TestNew_file_finder(t *testing.T) {
	//if folder_to_watch is empty
	no_file, _ := new_file_finder(" ")

	if no_file != " " {
		t.Error("new_file_finder failed, expected %v, got %v", " ", no_file) // to indicate test failed
	}

	//test for existing new file
	/*
		folder_to_watch := testfile
		filename, _ := new_file_finder(folder_to_watch){
			t.Error("new_file_finder failed, expected %v, got %v", filename, " ") // to indicate test failed
		}
	*/

}
