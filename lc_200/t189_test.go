package lc_200

import (
	"fmt"
	"testing"
)

// 189. 轮转数组

// 方法一：额外数组，golang切片截取拼接
func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	// k 可能大于 nums 长度
	k %= len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}

// 方法二：额外数组，算出位置
func rotate2(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

// 方法三：数组翻转，空间复杂的 O(1)
// 参考：https://leetcode-cn.com/problems/rotate-array/solution/xuan-zhuan-shu-zu-by-leetcode-solution-nipk/
// 该方法基于如下的事实：当我们将数组的元素向右移动 k 次后，尾部 k mod n 个元素会移动至数组头部，其余元素向后移动 k mod n 个位置。
// 该方法为数组的翻转：我们可以先将所有元素翻转，这样尾部的 k mod n 个元素就被移至数组头部，然后我们再翻转 [0, k mod n-1]区间的元素和 [k mod n, n-1] 区间的元素即能得到最后的答案。
// 例如： nums = [1,2,3,4,5] k=2
// 翻转所有：nums = [5,4,3,2,1] ，
// 再分别翻转 5,4 和 3,2,1 两个子串，nums = [4,5,1,2,3]
func rotate3(nums []int, k int) {

	reverse := func(a []int) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}

	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

}

func Test_rotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 30)
	fmt.Println(nums)
}
