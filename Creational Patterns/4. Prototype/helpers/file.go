package helpers

import (
	"fmt"
	"prototype/domain"
)

type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *File) Clone() domain.Inode {
	return &File{Name: f.Name + "_clone"}
}