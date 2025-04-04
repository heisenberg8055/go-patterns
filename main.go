package main

import (
	abs "github.com/heisenberg8055/go-patterns/structural/composite"
)

func main() {

	file1 := &abs.File{Name: "File1"}
	file2 := &abs.File{Name: "File2"}
	file3 := &abs.File{Name: "File3"}

	folder1 := &abs.Folder{Name: "Folder1"}
	folder1.Add(file1)

	folder2 := &abs.Folder{Name: "Folder2"}

	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("test")
}
