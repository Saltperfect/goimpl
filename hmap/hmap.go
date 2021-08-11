package hmap

import (
	"fmt"

	"github.com/saltperfect/goimpl/list"
)

const (
	m = 1000   // size of hmap
	p = 101429 // prime greater than universe of key 100000
	a = 3456   // random a
	b = 8934   // random b
)

type Hmap [1000]*list.KeyValueLinklist

/*
	universal hashing algo: h(k) = [ [(a*k+b) mod p] mod m ]
*/
func hash(k int) int {
	return ((a*k + b) % p) % m
}

func NewHmap() *Hmap {
	return &Hmap{}
}

func (h *Hmap) Put(key int, value interface{}) error {
	index := hash(key)
	if h[index] == nil {
		l := list.NewKeyValuelist()
		h[index] = l
	}
	h[index].PutKeyValue(key, value)
	return nil
}

func (h *Hmap) Exists(key int) (bool, error){
	index := hash(key)
	if h[index] == nil {
		return false, nil
	}
	return h[index].KeyExists(key)
}

func (h *Hmap) Print() error {
	for i:= 0; i < 1000 ; i++ {
		if h[i] != nil {
			fmt.Printf("index %d ", i)
			h[i].PrintKeyValueList()
			fmt.Println("=")
		}
	}
	return nil
}