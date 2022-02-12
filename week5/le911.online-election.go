package week5

import "sort"

// https://leetcode-cn.com/problems/online-election/
// 911. 在线选举
type TopVotedCandidate struct {
	tops  []int
	times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	tops := make([]int, len(persons))
	top := -1
	voteCounts := map[int]int{}
	voteCounts[-1] = -1
	for i, p := range persons {
		voteCounts[p]++
		if voteCounts[p] >= voteCounts[top] {
			top = p
		}
		tops[i] = top
	}
	return TopVotedCandidate{
		tops:  tops,
		times: times,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	id := sort.SearchInts(this.times, t+1) - 1
	return this.tops[id]
}
