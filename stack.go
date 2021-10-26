package stack

type Obj interface{}
type Stack []Obj
// NewStack -
func NewStack() Stack {
	return Stack{}
}

func (s *Stack) reduce() Stack {
	tmp := make(Stack, s.Length()-1)
	copy(tmp,*s)
	return tmp
}

func (s Stack) IsEmpty() bool {
	if s.Length() == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(o Obj) {
	*s = append(*s,o)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	
	str := (*s)[s.Length()-1]
	*s = s.reduce()
	return str
}

func (s Stack) Length() int {
	return len(s)
}
