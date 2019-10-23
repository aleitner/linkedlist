package linkedlist

import (
	"fmt"
	"linkedlist/singly"
)

func main() {
	fmt.Println("Creating List...")
	s := singly.NewSinglyList(1337)

	fmt.Println("Adding item...")
	if err := s.Add("item #1"); err != nil {
		panic(err)
	}

	fmt.Println("Adding item...")
	if err := s.Add("item #2"); err != nil {
		panic(err)
	}

	fmt.Println("Adding item...")
	if err := s.Add("item #3"); err != nil {
		panic(err)
	}

	currentItem, err := s.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Current item: ", currentItem)

	fmt.Println("Moving Next...")
	currentItem, err = s.Next()
	if err != nil {
		panic(err)
	}

	fmt.Println("Current item: ", currentItem)

	fmt.Println("Total List:\n", s)
}
