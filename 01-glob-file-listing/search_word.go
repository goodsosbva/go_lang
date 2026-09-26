package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("need to 2 more word. ex) search_wird word filepath") 
		return
	}

	word := os.Args[1]
	files := os.Args[2:] 
	fmt.Println("찾으려는 단어:", word)
	PrintAllFiles(files)
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	for _, path := range files {
			filelist, err := GetFileList(path)
			if err != nil {
				fmt.Println("don`t search the file. err:", err);
				return
			}

			fmt.Println("file list")
			for _, name := range filelist {
				fmt.Println(name)
		}
	}
}
