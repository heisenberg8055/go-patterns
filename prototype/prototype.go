package prototype

import "fmt"

func Prototype() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{name: "Folder1", children: []Inode{file1}}

	folder2 := &Folder{name: "Folder2", children: []Inode{folder1, file2, file3}}

	fmt.Println("Folder2 structure")

	folder2.print("  ")

	cloned := folder2.clone()

	fmt.Println("Cloned Folder2 structure")
	cloned.print("  ")
}
