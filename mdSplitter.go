package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Thanks https://stackoverflow.com/a/59961623
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		 return nil, err
	}
	return os.Create(p)
}

func main() {
	argsWithProg := os.Args
	argsWithoutProg := argsWithProg[1:]
	if len(argsWithoutProg) < 1 {
		log.Fatal("Expected input file as argument")
	}

	rawMD, err := os.ReadFile(argsWithoutProg[0])
	if err != nil {
		log.Fatal(err);
	}

	sliced := strings.Split(string(rawMD),"---")

	/* This is the 
		```
		Author: Igor bedesqui
		Code: ib
		``` 
		from the frontmatter, so you'll need to add the --- before and after. I don't include them in case you wanna add more stuff manually (which I will 👺)
	*/
	frontMatterContent := sliced[1];

	chapters := strings.Split(sliced[2], "##")
	
	for i, chapterContent:= range chapters {
		if i == 0 {
			continue;
		}

		f, err := create(fmt.Sprintf("out/%d.md", i))
		if err != nil {
			log.Fatal(err)
			return
		}
		defer f.Close()
		_, err2 := f.WriteString(fmt.Sprintf("---%sChapter: %d\n---\n## %s\n", frontMatterContent, i,  chapterContent))
		if err2 != nil {
			log.Fatal(err2)
			return
		}
	}

	fmt.Println("\nDone!... Probably")
}