package main

import "fmt"

func mainExample() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic recovered: %v\n", r)
		}
	}()

	panic("random panic!!!")

}
