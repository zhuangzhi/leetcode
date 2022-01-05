package leetcode

type anystack []interface{}

func newanystack(cap int) anystack {
	return anystack(make([]interface{}, 0, cap))
}

func (s anystack) empty() bool {
	return len(s) == 0
}

func (s anystack) top() interface{} {
	return s[len(s)-1]
}

func (s *anystack) pop() {
	x := *s
	*s = x[:len(x)-1]
}

func (s *anystack) push(c interface{}) {
	*s = append(*s, c)
}

type intstack []int

func (s intstack) empty() bool {
	return len(s) == 0
}

func (s intstack) top() int {
	return s[len(s)-1]
}

func (s *intstack) pop() {
	x := *s
	*s = x[:len(x)-1]
}

func (s *intstack) push(c int) {
	*s = append(*s, c)
}

type charstack []rune

func (s charstack) empty() bool {
	return len(s) == 0
}

func (s charstack) top() rune {
	return s[len(s)-1]
}

func (s *charstack) pop() {
	x := *s
	*s = x[:len(x)-1]
}

func (s *charstack) push(c rune) {
	*s = append(*s, c)
}

type strstack []string

func (s strstack) empty() bool {
	return len(s) == 0
}

func (s strstack) top() string {
	return s[len(s)-1]
}

func (s *strstack) pop() {
	x := *s
	*s = x[:len(x)-1]
}

func (s *strstack) push(str string) {
	*s = append(*s, str)
}
