package composite

import "fmt"

type Folder struct {
	Components []Component
	Name       string
}

func (f *Folder) Search(keyWord string) {
	fmt.Printf("Searching for %v in %v folder\n", keyWord, f.Name)
	for _, component := range f.Components {
		component.Search(keyWord)
	}
}

func (f *Folder) Add(c Component) {
	f.Components = append(f.Components, c)
}
