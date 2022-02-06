package util

func Range(i, j int, k ...int) []int {
	step := 1
	if i > j {
		step = -1
	}
	if len(k) > 0 {
		step = k[0]
	}
	if j > i {
		s := make([]int, 0, j-i)
		for i < j {
			s = append(s, i)
			i += step
		}
		return s
	} else {
		s := make([]int, 0, i-j)
		for i > j {
			s = append(s, i)
			i += step
		}
		return s
	}
}

type StrIntMap map[string]int

func (m StrIntMap) Add(key string, val int) int {
	m[key] += val
	return m[key]
}

func (m StrIntMap) Contains(key string) bool {
	_, ok := m[key]
	return ok
}

func (m StrIntMap) Equals(with StrIntMap) bool {
	return m.ContainsAll(with) && with.ContainsAll(m)
}

func (m StrIntMap) ContainsAll(with StrIntMap) bool {
	for k, v := range with {
		if x, ok := m[k]; !ok || v != x {
			return false
		}
	}
	return true
}
