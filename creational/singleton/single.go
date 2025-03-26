package single

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating...")
			singleInstance = &single{}
		} else {
			fmt.Println("Already created1")
		}
	} else {
		fmt.Println("Already created2")
	}
	return singleInstance
}
