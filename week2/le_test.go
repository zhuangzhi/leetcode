package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLe697(t *testing.T) {
	units := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 2, 3, 1}, 2},
		{[]int{1, 2, 2, 3, 1, 4, 2}, 6},
	}

	for _, u := range units {
		assert.Equal(t, u.expected, findShortestSubArray(u.nums))
	}
}

func TestLe811(t *testing.T) {
	units := []struct {
		cpdomains []string
		expected  []string
	}{
		{[]string{"9001 discuss.leetcode.com"}, []string{"9001 leetcode.com", "9001 discuss.leetcode.com", "9001 com"}},
		{
			[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			[]string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
		{[]string{}, []string{}},
	}

	for _, u := range units {
		assert.ElementsMatch(t, u.expected, subdomainVisits(u.cpdomains))
	}
}
