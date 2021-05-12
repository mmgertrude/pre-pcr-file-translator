//source: https://stackoverflow.com/questions/45578172/golang-find-most-recent-file-by-date-and-time
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//global variables:
var path = "./testfolder/"

func main() {
	newestFile := new_file_discover()
	fmt.Println(newestFile)

	//fetch requisition name from unilab API ?
	unilab_name_or_file := "name returned from unilab"
	new_req_name_file := unilab_name_or_file
	fmt.Println(new_req_name_file)

	file_to_write := req_file_modifier(new_req_name_file, newestFile)
	//write function to save file_to_write
	fmt.Println(file_to_write)

}

//has no testcode
func new_file_discover() string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var newestFile string
	var newestTime int64 = 0
	for _, f := range files {
		fi, err := os.Stat(path + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		currTime := fi.ModTime().Unix()
		if currTime > newestTime {
			newestTime = currTime
			newestFile = f.Name()
		}
	}
	return (newestFile)
}

//has no test code
func req_file_modifier(new_req_name_file, newestFile string) string {
	//do stuff ie
	// replace requisition id in newestfile with requisition name
	//return modified file name
	fmt.Println(new_req_name_file, newestFile)
	return "file_to_save"

}
