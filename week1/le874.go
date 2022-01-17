package leetcode

import (
	. "github.com/zhuangzhi/leetcode/util"
)

func robotSim(commands []int, obstacles [][]int) int {
	robot := robot{}
	return robot.sim(commands, obstacles)
}

type obstruct map[int]int

func hash(x, y int) int {
	return x*1e5 + y
}

func buildObstruct(obstacles [][]int) obstruct {
	obs := make(map[int]int)
	for _, pair := range obstacles {
		obs[hash(pair[0], pair[1])] = 1
	}
	return obs
}

func hashObstruct(obs obstruct, x, y int) bool {
	_, ok := obs[hash(x, y)]
	return ok
}

type robot struct {
	x, y int
	dir  direction
	obs  obstruct
	ans  int
}

func (r *robot) sim(commands []int, obstacles [][]int) int {
	r.x, r.y = 0, 0
	r.dir = direction{0, 1}
	r.obs = buildObstruct(obstacles)
	for _, cmd := range commands {
		r.take(cmd)
	}
	return r.ans
}

func (r *robot) take(cmd int) {
	if cmd < 0 {
		r.dir.turn(cmd)
	} else {
		r.run(cmd)
	}
}

func (r *robot) run(step int) {
	for s := 0; s < step; s++ {
		x1 := r.x + r.dir[0]
		y1 := r.y + r.dir[1]
		if hashObstruct(r.obs, x1, y1) {
			return
		}
		r.x, r.y = x1, y1
		r.ans = MaxInt(r.ans, x1*x1+y1*y1)
	}
}

type direction [2]int

func (d *direction) turn(cmd int) {
	switch cmd {
	case -2:
		d[0], d[1] = d[1], d[0]
		d[0] = -d[0]
	case -1:
		d[0], d[1] = d[1], d[0]
		d[1] = -d[1]
	}
}
