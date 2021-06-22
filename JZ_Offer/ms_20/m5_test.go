package ms_20

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// 面试题05. 替换空格
// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

// 1. 暴力法？
// 2. 用api嘛？这没意思吧，还能偷偷看api怎么做的，哈哈
// 3. 用数组做？在同一个数据数组上，不用额外空间？
// 4. 用切片做？不同语言有各自特色的数据结构，貌似实现起来也有差异
// 5. 现在内存便宜，应该用更节省cpu的做法？
// 6. copy数据，用string相加来做，和用api有啥区别嘛？感觉这题不复合时代背景了，应该加更多的限制条件才行
func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func Test_replaceSpace(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(replaceSpace("We are happy."), "We%20are%20happy.")
}
