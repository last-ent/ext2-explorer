package main

import (
	"fmt"
	"log"

	"github.com/last-ent/ext2-explorer/ext2"
)

func main() {

	fs := ext2.NewFSReader("/home/entux/Documents/Code/fsfs/linux.ex2", ext2.DEFAULT_BLOCK_SIZE)

	for _, k := range fs.RootDir.Dentries {
		fmt.Println(k.Name, fmt.Sprintf("(%s)", k.FileType))
	}

	fileData, err := fs.ReadFile("/", "abc.txt")
	if err != nil {
		log.Println("Error:", err.Error())
		return
	}
	fmt.Println(string(fileData))
}
