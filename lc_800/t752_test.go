package lc_800

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 752. 打开转盘锁

// BFS 求 0000 -> target 的最短路径，避开 deadends

// 密码盘
type LockNum struct {
	A, B, C, D int
}

type Job struct {
	lockNum LockNum
	deep    int
}

// 从字符串获取密码盘对象
func NewNum(s string) LockNum {
	// 有题目的前提，省去s的合规性验证了

	return LockNum{
		A: int(s[0]) - 48,
		B: int(s[1]) - 48,
		C: int(s[2]) - 48,
		D: int(s[3]) - 48,
	}
}

// 获取下一步可以到达密码盘 如 0000 —> 0001,0010,0100,1000,9000,0900,0090,0009
func (this *LockNum) getNext() []LockNum {
	return []LockNum{
		{
			A: LSpin(this.A),
			B: this.B,
			C: this.C,
			D: this.D,
		},
		{
			A: this.A,
			B: LSpin(this.B),
			C: this.C,
			D: this.D,
		},
		{
			A: this.A,
			B: this.B,
			C: LSpin(this.C),
			D: this.D,
		},
		{
			A: this.A,
			B: this.B,
			C: this.C,
			D: LSpin(this.D),
		},
		{
			A: RSpin(this.A),
			B: this.B,
			C: this.C,
			D: this.D,
		},
		{
			A: this.A,
			B: RSpin(this.B),
			C: this.C,
			D: this.D,
		},
		{
			A: this.A,
			B: this.B,
			C: RSpin(this.C),
			D: this.D,
		},
		{
			A: this.A,
			B: this.B,
			C: this.C,
			D: RSpin(this.D),
		},
	}
}

// LSpin 左旋
func LSpin(n int) int {
	if n == 0 {
		return 9
	} else {
		return n - 1
	}
}

// RSpin 右旋转
func RSpin(n int) int {
	if n == 9 {
		return 0
	} else {
		return n + 1
	}
}

func openLock(deadends []string, target string) int {

	// 数据格式转换，
	dead := make(map[LockNum]bool, len(deadends))
	for i, _ := range deadends {
		dead[NewNum(deadends[i])] = true
	}

	// 任务队列
	queue := make([]Job, 0, 20)
	queue = append(queue, Job{NewNum("0000"), 0})
	// 已经出现过的数字
	exit := make(map[LockNum]bool)
	exit[NewNum("0000")] = true
	// 目标
	tag := NewNum(target)

	// 边界值处理
	if dead[NewNum("0000")] == true {
		return -1
	}
	if target == "0000" {
		return 0
	}

	for len(queue) > 0 {

		// 从队列中取一个
		job := queue[0]
		queue = queue[1:]

		for _, ln := range job.lockNum.getNext() {
			if dead[ln] != true && exit[ln] != true {
				if ln == tag {
					return job.deep + 1
				} else {
					queue = append(queue, Job{lockNum: ln, deep: job.deep + 1})
					exit[ln] = true
				}
			}
		}

	}

	return -1

}

func Test_openLock(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"), 6)
	assertions.Equal(openLock([]string{"8888"}, "0009"), 1)
	assertions.Equal(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"), -1)
	assertions.Equal(openLock([]string{"0000"}, "8888"), -1)
}
