package util

type AnyStack []interface{}

func NewStack(cap int) AnyStack {
	return AnyStack(make([]interface{}, 0, cap))
}

func (s AnyStack) Empty() bool {
	return len(s) == 0
}

func (s AnyStack) Top() interface{} {
	return s[len(s)-1]
}

func (s *AnyStack) Pop() {
	x := *s
	*s = x[:len(x)-1]
}

func (s *AnyStack) Push(c interface{}) {
	*s = append(*s, c)
}

type IntStack []int

func NewIntStack(cap int) IntStack {
	return make([]int, 0, cap)
}

func (s IntStack) Empty() bool {
	return len(s) == 0
}

func (s IntStack) Top() int {
	return s[len(s)-1]
}

func (s *IntStack) Pop() {
	x := *s
	*s = x[:len(x)-1]
}

func (s *IntStack) Push(c int) {
	*s = append(*s, c)
}

type Pair struct {
	First, Second int
}
type PairStack []Pair

func NewPairStack(cap int) PairStack {
	return make([]Pair, 0, cap)
}
func (s PairStack) Empty() bool {
	return len(s) == 0
}

func (s PairStack) Top() *Pair {
	if s.Empty() {
		return nil
	}
	return &s[len(s)-1]
}

func (s *PairStack) Pop() {
	x := *s
	*s = x[:len(x)-1]
}

func (s *PairStack) Push(c Pair) {
	*s = append(*s, c)
}
func (s *PairStack) Append(first, second int) {
	*s = append(*s, Pair{First: first, Second: second})
}

type CharStack []rune

func NewCharStack(cap int) CharStack {
	return make([]rune, 0, cap)
}

func (s CharStack) Empty() bool {
	return len(s) == 0
}

func (s CharStack) Top() rune {
	return s[len(s)-1]
}

func (s *CharStack) Pop() {
	x := *s
	*s = x[:len(x)-1]
}

func (s *CharStack) Push(c rune) {
	*s = append(*s, c)
}

type StrStack []string

func NewStrStack(cap int) StrStack {
	return make([]string, 0, cap)
}

func (s StrStack) Empty() bool {
	return len(s) == 0
}

func (s StrStack) Top() string {
	return s[len(s)-1]
}

func (s *StrStack) Pop() {
	x := *s
	*s = x[:len(x)-1]
}

func (s *StrStack) Push(str string) {
	*s = append(*s, str)
}
