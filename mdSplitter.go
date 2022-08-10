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
	content, err := os.ReadFile("input.md")
	if err != nil {
		log.Fatal(err);
	}

	fmt.Printf("%s", content)

	fmt.Println("Done!... Probably")
}

// for later use
/* 
i := 1
	for i < 9 {
		f, err := create(fmt.Sprintf("out/%d.md", i))
	
		if err != nil {
			log.Fatal(err)
			return
		}
	
		defer f.Close()
	
		_, err2 := f.WriteString(fmt.Sprintf("# %d\n", i))
	
		if err2 != nil {
			log.Fatal(err2)
			return
		}

		i++
	}
*/