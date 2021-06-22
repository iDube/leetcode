package lc_6000

import (
	"fmt"
	"strings"
	"testing"
)

// 5413. 重新排列句子中的单词

// 变体的排序

func arrangeWords(text string) string {

	text = strings.ToLower(text)
	word := strings.Split(text, " ")

	size := len(word)

	// 冒泡排序，只不过把比较的数据由数组的元素的值变为了数组元素值的长度
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if len(word[j]) > len(word[j+1]) {
				word[j], word[j+1] = word[j+1], word[j]
			}
		}
	}

	result := strings.Join(word, " ")

	result = string(result[0]-32) + result[1:]

	return result

}

func TestArrangeWords(t *testing.T) {
	fmt.Println(arrangeWords("Keep calm and code on"))
	fmt.Println(arrangeWords("Leetcode is cool"))
	fmt.Println(arrangeWords("To be or not to be"))
}
