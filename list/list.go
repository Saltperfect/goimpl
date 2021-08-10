package list

import "fmt"

type linknode struct {
	val  interface{}
	next *linknode
}

type Linklist struct {
	Head *linknode
	curr *linknode
	size int
}

func Newlist() *Linklist {
	return &Linklist{}
}

func (l *Linklist) Put(i interface{}) error {
	node := linknode{val: i}
	if l.Head == nil {
		l.Head = &node
	}
	if l.curr != nil {
		l.curr.next = &node
	}
	l.curr = &node
	l.size++
	return nil
}

func (l *Linklist) Print() error {
	if l.Head == nil {
		return fmt.Errorf("empty list")
	}
	tmp := l.Head
	for i := 0; i < l.size; i++ {
		fmt.Printf("%+v,", tmp.val)
		tmp = tmp.next
	}
	return nil
}

func (l *Linklist) Size() (int, error) {
	return l.size, nil
}

func (l *Linklist) Empty() (bool, error) {
	return l.size == 0, nil
}

func (l *Linklist) ValueAt(index int) (interface{}, error) {
	if index >= l.size {
		return nil, fmt.Errorf("index out of bound")
	}
	tmp := l.Head
	for i := 0; i < index; i++ {
		tmp = tmp.next
	}
	return tmp.val, nil
}
func (l *Linklist) Insert(index int, val interface{}) error {
	if index >= l.size {
		return fmt.Errorf("index out of bound")
	}
	tmp := l.Head
	for i := 1; i < index; i++ {
		tmp = tmp.next
	}
	node := linknode{val: val, next: tmp.next}
	tmp.next = &node
	l.size++
	return nil
}
