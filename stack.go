package xmltool

type stack []interface{}

func (stk *stack) push(items ...interface{}) (n int) {
	for i, s := range items {
		*stk = append(*stk, s)
		n = i
	}
	return n + 1
}

func (stk *stack) len() int {
	return len(*stk)
}

func (stk *stack) pop() (interface{}, bool) {
	if stk.len() > 0 {
		last := (*stk)[stk.len()-1]
		*stk = (*stk)[:stk.len()-1]
		return last, true
	}
	return nil, false
}

func (stk *stack) peek() (interface{}, bool) {
	if stk.len() > 0 {
		return (*stk)[stk.len()-1], true
	}
	return nil, false
}

func (stk *stack) clear() (n int) {
	n = stk.len()
	stk = &stack{}
	return n
}
