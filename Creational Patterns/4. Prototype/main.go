package main

//! Important :
// Prototype is a creational design pattern that allows cloning objects,
// even complex ones, without coupling to their specific classes.

import (
	"fmt"
	"prototype/domain"
	"prototype/helpers"
)

func main() {
    file1 := &helpers.File{Name: "File1"}
    file2 := &helpers.File{Name: "File2"}
    file3 := &helpers.File{Name: "File3"}

    folder1 := &helpers.Folder{
        Children: []domain.Inode{file1},
        Name:      "Folder1",
    }

    folder2 := &helpers.Folder{
        Children: []domain.Inode{folder1, file2, file3},
        Name:      "Folder2",
    }
    fmt.Println("\nPrinting hierarchy for Folder2")
    folder2.Print("  ")

    cloneFolder := folder2.Clone()
    fmt.Println("\nPrinting hierarchy for clone Folder")
    cloneFolder.Print("  ")
}