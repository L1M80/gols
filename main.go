package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir(".")

	if err != nil {
		panic(err)
	}

	showNames(files)
}

func showNames(files []os.FileInfo) {
	for _, file := range files {
		fmt.Printf("%s ", file.Name())
	}
}
