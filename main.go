package main

import (
	"concurrency/service"
	"fmt"
)

func main() {

	err := service.Init()
	if err != nil {
		fmt.Println("unexpected error , terminating ", err)
	}

}
