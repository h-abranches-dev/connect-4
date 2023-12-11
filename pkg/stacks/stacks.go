package stacks

type Stack []rune

func New() *Stack {
	return &Stack{}
}

func (st *Stack) Push(r rune) {
	*st = append(*st, r)
}

func (st *Stack) Pop() rune {
	popped := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return popped
}
