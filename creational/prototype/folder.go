package prototype

import "fmt"

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indent string) {
	fmt.Println(indent + f.name)
	for _, i := range f.children {
		i.print(indent + indent)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChild []Inode
	for _, i := range f.children {
		copy := i.clone()
		tempChild = append(tempChild, copy)
	}
	cloneFolder.children = tempChild
	return cloneFolder
}
