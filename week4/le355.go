package week4

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/design-twitter/
// 355. 设计推特
type Twitter struct {
	users map[int]*TwNode
	time  int
}

type TwNode struct {
	followee map[int]int // Set
	tweet    Queue
}
type Tweet [2]int

const tweetCount = 10

func Constructor() Twitter {
	return Twitter{
		users: map[int]*TwNode{},
	}
}

func newNode() *TwNode {
	return &TwNode{
		followee: map[int]int{},
		tweet:    NewFixedRingQueue(tweetCount),
	}
}

func (tw *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := tw.users[userId]; !ok {
		tw.users[userId] = newNode()
	}
	user := tw.users[userId]
	tw.time++
	user.tweet.PushBack(Tweet{tweetId, tw.time})
}

func (tw *Twitter) GetNewsFeed(userId int) []int {
	user := tw.users[userId]
	if user == nil {
		return nil
	}
	users := make([]*TwNode, 0, 100)
	users = append(users, user)
	for id := range user.followee {
		if u, ok := tw.users[id]; ok {
			users = append(users, u)
		}
	}

	heap := HeapSlice(make([]Tweet, 0, 10), func(i, j interface{}) bool {
		return i.(Tweet)[1] < j.(Tweet)[1]
	})

	for _, u := range users {
		tws := u.tweet.Copy()
		for _, tw := range tws {
			if heap.Len() < tweetCount {
				heap.Push(tw)
			} else {
				if tw.(Tweet)[1] > heap.Top().(Tweet)[1] {
					heap.ReplaceTop(tw)
				}
			}
		}
	}
	l := heap.Len()
	ans := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		tw := heap.Pop().(Tweet)
		ans[i] = tw[0]
	}
	return ans
}

func (tw *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := tw.users[followerId]; !ok {
		tw.users[followerId] = newNode()
	}
	tw.users[followerId].followee[followeeId] = followeeId
}

func (tw *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := tw.users[followerId]; !ok {
		tw.users[followerId] = newNode()
	}
	delete(tw.users[followerId].followee, followeeId)
}
