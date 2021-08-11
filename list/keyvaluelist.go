package list

import (
	"fmt"
)

type KeyValue struct {
	Key int
	Value interface{}
}

type keyValueNode struct {
	val  KeyValue
	next *keyValueNode
}

type KeyValueLinklist struct {
	Head *keyValueNode
	curr *keyValueNode
	size int
}

func NewKeyValuelist() *KeyValueLinklist {
	return &KeyValueLinklist{}
}

func (l *KeyValueLinklist) PutKeyValue(key int, i interface{}) error {
	node := keyValueNode{val: KeyValue{Key: key, Value: i}}
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

func (l *KeyValueLinklist) PrintKeyValueList() error {
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

func (l *KeyValueLinklist) KeyValueSize() (int, error) {
	return l.size, nil
}

func (l *KeyValueLinklist) KeyValueEmpty() (bool, error) {
	return l.size == 0, nil
}

func (l *KeyValueLinklist) ListValueAt(index int) (interface{}, error) {
	if index >= l.size {
		return nil, fmt.Errorf("index out of bound")
	}
	tmp := l.Head
	for i := 0; i < index; i++ {
		tmp = tmp.next
	}
	return tmp.val, nil
}
func (l *KeyValueLinklist) KeyValueListInsert(index int, key int, val interface{}) error {
	if index >= l.size {
		return fmt.Errorf("index out of bound")
	}
	tmp := l.Head
	for i := 1; i < index; i++ {
		tmp = tmp.next
	}
	node := keyValueNode{val: KeyValue{Key: key, Value: val}, next: tmp.next}
	tmp.next = &node
	l.size++
	return nil
}

func (l *KeyValueLinklist) KeyValueListReverse() error {
	if l.size < 2 {
		return nil
	}
	var s, f *keyValueNode
	for s, f = l.Head, l.Head.next; f != nil; {
		tmp := f.next
		f.next = s
		s = f
		f = tmp
	}
	l.curr = l.Head
	l.Head = s
	return nil
}

func (l *KeyValueLinklist) KeyExists(key int) (bool, error ){
	tmp := l.Head
	for i:= 0; i < l.size; i++ {
		if tmp.val.Key == key { 
			return true, nil
		}
		tmp = tmp.next
	}
	return false, nil
}
