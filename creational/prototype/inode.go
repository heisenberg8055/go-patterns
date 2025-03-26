package prototype

// Prototype Interface
type Inode interface {
	print(string)
	clone() Inode
}
