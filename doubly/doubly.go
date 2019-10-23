package doubly

import (
	"errors"
	"fmt"
	"sync"
)

type DoublyList struct {
	ID      int
	head    *DoublyItem
	tail    *DoublyItem
	current *DoublyItem
	mtx     sync.Mutex
}

type DoublyItem struct {
	data     string
	next     *DoublyItem
	previous *DoublyItem
}

func NewDoublyList(id int) *DoublyList {
	return &DoublyList{
		ID: id,
	}
}

func (d *DoublyList) Add(data string) error {
	item := &DoublyItem{
		data: data,
	}

	d.mtx.Lock()
	defer d.mtx.Unlock()

	if d.head == nil {
		d.head = item
	} else {
		nodeToCheck := d.tail
		nodeToCheck.next = item
		item.previous = nodeToCheck
	}

	d.tail = item

	return nil
}

func (d DoublyList) Current() (string, error) {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	if d.current == nil {
		// NB: We could also just set d.current to d.head
		return "", errors.New("No current item")
	}

	return d.current.data, nil
}

func (d *DoublyList) Next() (string, error) {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	if d.current == nil{
		// NB: We could also just set d.current to d.head
		return "", errors.New("No current item")
	}

	if d.current.next == nil{
		return "", errors.New("No more items in list")
	}

	d.current = d.current.next

	return d.current.data, nil
}

func (d *DoublyList) Previous() (string, error) {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	if d.current == nil {
		// NB: We could also just set d.current to d.head
		return "", errors.New("No current item")
	}

	if d.current.previous == nil{
		return "", errors.New("No more items in list")
	}

	d.current = d.current.previous

	return d.current.data, nil
}

func (d *DoublyList) Beginning() error {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	d.current = d.head
	return nil
}

func (d DoublyList) String() string {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	stringVal := "[ "

	item := d.head
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
