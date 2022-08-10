package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Thanks https://stackoverflow.com/a/59961623
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		 return nil, err
	}
	return os.Create(p)
}

func main() {
	f, err := create("out/test.md")
	
	if err != nil {
		log.Fatal(err)
		return
	}

	defer f.Close()

	_, err2 := f.WriteString("old falcon\n")

	if err2 != nil {
		log.Fatal(err2)
		return
	}

	fmt.Println("Done!... Probably")
}