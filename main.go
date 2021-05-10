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
}

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
