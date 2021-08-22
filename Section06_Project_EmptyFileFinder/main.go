package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("USAGE: [DIR TO SEARCH IN]")
	}

	path := args[0]

	fi, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	emptyFiles := make([]byte, 0, calculateSize(fi))

	for _, f := range fi {
		if !f.IsDir() && f.Size() == 0 {
			n := f.Name() + "\n"
			emptyFiles = append(emptyFiles, n...)
			fmt.Println(f.Name())
		}
	}

	if len(emptyFiles) == 0 {
		log.Fatal("Empty Files Not Found")
	}


	err = ioutil.WriteFile("result.txt", emptyFiles, 0755)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

}

func calculateSize(fi []fs.FileInfo) (size int) {
	for _, f := range fi {
		if !f.IsDir() && f.Size() == 0 {
			size += len(f.Name()) + 1
		}
	}

	return
}
