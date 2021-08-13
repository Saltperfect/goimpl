package tree

import (
	"fmt"
)

type treenode struct {
	val   int
	left  *treenode
	right *treenode
}

func (t treenode) String() string {
	return fmt.Sprintf("%d,", t.val)
}

type Tree struct {
	Head *treenode
}

func NewTree() *Tree {
	return new(Tree)
}

func (t *Tree) Put(val int) error {
	if t.Head == nil {
		t.Head = &treenode{val: val}
		return nil
	}
	tmp := t.Head

	for tmp != nil {
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
		if tmp.val == val {
			return fmt.Errorf("value already exists")
		}
	}
	return nil
}

func (t *Tree) Search(val int) (bool, error) {
	if t.Head == nil {
		return false, fmt.Errorf("empty tree recieved")
	}
	tmp := t.Head
	res := false

	var searchdfs func(t *treenode, val int)

	searchdfs = func(t *treenode, val int) {
		if t == nil {
			res = false
			return
		}
		switch {
		case val < t.val:
			searchdfs(t.left, val)
		case val > t.val:
			searchdfs(t.right, val)
		case val == t.val:
			res = true
		}
	}
	searchdfs(tmp, val)
	return res, nil
}

func (t *Tree) Print() error {
	if t.Head == nil {
		return fmt.Errorf("nothing to print")
	}
	q := []*treenode{t.Head}

	for len(q) != 0 {
		count := 0
		for _, v := range q {
			if v == nil {
				count++
			}
		}
		if count == len(q) {
			break
		}
		for k := len(q); k > 0; k-- {
			if q[0] == nil {
				fmt.Printf("_, ")
				q = append(q, nil, nil)
				q = q[1:]
				continue
			}
			fmt.Printf("%+v , ", q[0].val)
			q = append(q, q[0].left, q[0].right)
			q = q[1:]
		}
		fmt.Println("\n ============")
	}
	fmt.Print("\n")
	return nil
}

func (t *Tree) Printdfs() error {
	dfs(t.Head)
	return nil
}

func dfs(t *treenode) {
	if t == nil {
		return
	}
	dfs(t.left)
	fmt.Printf("%d, ", t.val)
	dfs(t.right)
}

func (t *Tree) Height() int {
	var heightdfs func(t *treenode) int

	heightdfs = func(t *treenode) int {
		if t == nil {
			return 0
		}
		leftheight := heightdfs(t.left)
		rightheight := heightdfs(t.right)
		if leftheight < rightheight {
			return rightheight + 1
		}
		return leftheight + 1

	}

	return heightdfs(t.Head)
}

