package singlylinkedlist

import (
	"errors"
	"fmt"
	"sync"
)

type Singly struct {
	ID int
	head *Item
	current *Item
	mtx sync.Mutex
}

type Item struct {
	data string
	next *Item
}

func New(id int) *Singly {
	return &Singly{
		ID: id,
	}
}

func (s *Singly) Add(data string) error {
	item := &Item{
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

	if s.current == nil {
		s.current = s.head
	}

	return nil
}

func (s Singly) Current() (string, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.current == nil {
		// NB: We could also just set s.current to s.head
		return "", errors.New("No current item")
	}

	return s.current.data, nil
}

func (s Singly) Next() (string, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.current == nil {
		// NB: We could also just set s.current to s.head
		return "", errors.New("No current item")
	}

	s.current = s.current.next

	return s.current.data, nil
}

func (s Singly) Beginning() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.current = s.head
	return nil
}

func (s Singly) String() string {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	stringVal := "[ "

	item := s.head
	for  {
		stringVal = fmt.Sprintf("%s \"%s\",", stringVal, item.data)

		if item.next == nil {
			break
		}

		item = item.next
	}

	stringVal = fmt.Sprintf("%s%s", stringVal, " ]")
	return stringVal
}