package singly

import (
	"errors"
	"fmt"
	"sync"
)

type SinglyList struct {
	ID      int
	head    *SinglyItem
	current *SinglyItem
	mtx     sync.Mutex
}

type SinglyItem struct {
	data string
	next *SinglyItem
}

func NewSinglyList(id int) *SinglyList {
	return &SinglyList{
		ID: id,
	}
}

func (s *SinglyList) Add(data string) error {
	item := &SinglyItem{
		data: data,
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.head == nil {
		s.head = item
	} else {
		nodeToCheck := s.head
		for nodeToCheck.next != nil {
			nodeToCheck = nodeToCheck.next
		}

		nodeToCheck.next = item
	}

	return nil
}

func (s SinglyList) Current() (string, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.current == nil {
		// NB: We could also just set s.current to s.head
		return "", errors.New("No current item")
	}

	return s.current.data, nil
}

func (s *SinglyList) Next() (string, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.current == nil {
		// NB: We could also just set s.current to s.head
		return "", errors.New("No current item")
	}

	if s.current.next == nil{
		return "", errors.New("No more items in list")
	}

	s.current = s.current.next

	return s.current.data, nil
}

func (s *SinglyList) Beginning() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.current = s.head
	return nil
}

func (s SinglyList) String() string {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	stringVal := "[ "

	item := s.head
	for {
		stringVal = fmt.Sprintf("%s \"%s\",", stringVal, item.data)

		if item.next == nil {
			break
		}

		item = item.next
	}

	stringVal = fmt.Sprintf("%s%s", stringVal, " ]")
	return stringVal
}
