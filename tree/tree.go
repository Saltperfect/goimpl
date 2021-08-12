package tree

import "fmt"

type treenode struct {
	val int
	left *treenode
	right *treenode
}

func (t treenode) String() string{
	return fmt.Sprintf("%d,", t.val)
}

type Tree struct {
	Head *treenode
}

func NewTree() *Tree{
	return new(Tree)
}

func ( t *Tree) Put(val int) error {
	if t.Head == nil {
		t.Head = &treenode{val: val}
		return nil
	}
	tmp := t.Head

	for ; tmp != nil; {
		if tmp.val < val {
			if tmp.right == nil {
				tmp.right = &treenode{val: val}
				return nil
			}
			tmp = tmp.right
		}
		if tmp.val > val {
			if tmp.left == nil {
				tmp.left = &treenode{val: val}
				return nil
			}
			tmp = tmp.left
		}
	}
	return nil
}

func (t * Tree) Print() error {
	if t.Head == nil {
		return fmt.Errorf("nothing to print")
	}
	q := []*treenode{t.Head}

	for ; len(q) != 0; {
		count := 0
		for _, v := range q {
			if v == nil {
				count++
			}
		}
		if count == len(q) {
			break
		}
		for k := len(q); k > 0 ; k-- {
			if q[0] == nil {
				fmt.Printf("_, ")
				q = append(q, nil, nil)
				q = q[1:]
				continue
			}
			fmt.Printf("%+v , ", q[0].val)
			q = append(q, q[0].left,q[0].right )
			q = q[1:]
		}
		fmt.Println("\n ============")
	}
	return nil
}
