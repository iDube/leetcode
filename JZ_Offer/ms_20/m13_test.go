package ms_20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 剑指 Offer 13. 机器人的运动范围

// 不是以小机器人的运动轨迹为主，而是判断最大方格中能满足条件的坐标以及它们之间的接触点。

// 方式一：递推
// 先不管m和n的值为多大，在最大方格上进行模拟，可以看出： 当k逐渐增大时，可到达的格子范围是以（0，0）为顶点的等腰直角三角形，例如：
//
// 当k = 8;
// 0  1  2  3  4  5  6 7  8  9 10
// 0 可 可 可 可 可 可 可 可 可 不 可
// 1 可 可 可 可 可 可 可 可 不 不 可
// 2 可 可 可 可 可 可 可 不 不 不 可
// 3 可 可 可 可 可 可 不 不 不 不 可
// 4 可 可 可 可 可 不 不 不 不 不 可
// 5 可 可 可 可 不 不 不 不 不 不 可
// 6 可 可 可 不 不 不 不 不 不 不 可                      （可为可到达的，不为不可到达的）
// 7 可 可 不 不 不 不 不 不 不 不 可
// 8 可 不 不 不 不 不 不 不 不 不 不
// 9 不 不 不 不 不 不 不 不 不 不 不
// 10可 可 可 可 可 可 可 可 可 可 不

type Point struct {
	X, Y int
}

// 方式二：广度优先搜索
func movingCount(m int, n int, k int) int {

	// 待处理节点队列，半生set保证节点不会重复放入队列. set 中放着已经放如过队列中的节点
	queue := make([]Point, 0, 10)
	set := make(map[Point]bool)

	count := 0

	zorePoint := Point{0, 0}
	queue = append(queue, zorePoint)
	set[zorePoint] = true

	for len(queue) > 0 {

		p := queue[0]
		queue = queue[1:]

		if csum(p.X, p.Y) <= k {
			count++
		} else {
			// 剪枝，不用继续了
			continue
		}

		// 向下和右眼神 （根据题目特点，不需要向左和上延伸）
		// 下
		if p.X+1 < m {
			pn := Point{p.X + 1, p.Y}
			if _, ok := set[pn]; !ok {
				queue = append(queue, pn)
				set[pn] = true
			}
		}
		// 右
		if p.Y+1 < n {
			pn := Point{p.X, p.Y + 1}
			if _, ok := set[pn]; !ok {
				queue = append(queue, pn)
				set[pn] = true
			}
		}

	}

	return count

}

// 按位求和
func csum(a, b int) int {
	sum := 0
	for ; a > 0; a /= 10 {
		sum += a % 10
	}
	for ; b > 0; b /= 10 {
		sum += b % 10
	}
	return sum
}

func TestCsum(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(csum(0, 0), 0)
	assertions.Equal(csum(1, 2), 3)
	assertions.Equal(csum(11, 2), 4)
	assertions.Equal(csum(11, 23), 7)
	assertions.Equal(csum(0, 23), 5)
	assertions.Equal(csum(100, 23), 6)
}

func TestMovingCount(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(movingCount(5, 5, 4), 15)
	assertions.Equal(movingCount(16, 8, 4), 15)
	assertions.Equal(movingCount(100, 100, 20), 6117)
	assertions.Equal(movingCount(2, 3, 1), 3)
	assertions.Equal(movingCount(18, 21, 13), 317)
}
