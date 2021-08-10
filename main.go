package main

import (
	"fmt"
	"log"

	"github.com/saltperfect/goimpl/list"
)

func main() {

	l := list.Newlist()
	err := l.Put(1)
	err = l.Put(`asv`)
	err = l.Put(map[string]string{})
	err = l.Put([]int{1,2,3})
	if err != nil {
		log.Fatal(err)
	}
	err = l.Insert(2, 45)
	v, err := l.ValueAt(2)
	fmt.Printf("value at 2 is %+v \n", v )
	l.Print() 
}