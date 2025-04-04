package composite

import "fmt"

type File struct {
	Name string
}

func (f *File) Search(keyWord string) {
	fmt.Printf("Searching for %v in File %v\n", keyWord, f.Name)
}

func (f *File) GetName() string {
	return f.Name
}
