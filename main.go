package main

import (
	"fmt"
	"ll/singlylinkedlist"
)

func main() {
	fmt.Println("Creating List...")
	singly := singlylinkedlist.New(1337)

	fmt.Println("Adding item...")
	if err := singly.Add("item #1"); err != nil {
		panic(err)
	}

	fmt.Println("Adding item...")
	if err := singly.Add("item #2"); err != nil {
		panic(err)
	}

	fmt.Println("Adding item...")
	if err := singly.Add("item #3"); err != nil {
		panic(err)
	}

	currentItem, err := singly.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Current item: ", currentItem)

	fmt.Println("Moving Next...")
	currentItem, err = singly.Next()
	if err != nil {
		panic(err)
	}

	fmt.Println("Current item: ", currentItem)

	fmt.Println("Total List:\n", singly)
}
