package algo

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleIntance *single

func getInstance() *single {
	if singleIntance == nil {
		once.Do(
			func() {
				singleIntance = &single{}
			})
	} else {
		fmt.Println("Instance had been created")
	}
	return singleIntance
}
