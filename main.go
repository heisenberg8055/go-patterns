package main

import (
	"fmt"

	abs "github.com/heisenberg8055/go-patterns/behavioral/iterator"
)

func main() {
	user1 := &abs.User{Name: "name1", Age: 10}

	user2 := &abs.User{Name: "name2", Age: 20}

	userCollection := abs.UserCollection{Users: []*abs.User{user1, user2}}

	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		fmt.Printf("User: %+v\n", iterator.GetNext())
	}
}
