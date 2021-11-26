//为了不在赎金信中暴露字迹，从杂志上搜索各个需要的字母，组成单词来表达意思。 
//
// 给你一个赎金信 (ransomNote) 字符串和一个杂志(magazine)字符串，判断 ransomNote 能不能由 magazines 里面的字符
//构成。 
//
// 如果可以构成，返回 true ；否则返回 false 。 
//
// magazine 中的每个字符只能在 ransomNote 中使用一次。 
//
// 
//
// 示例 1： 
//
// 
//输入：ransomNote = "a", magazine = "b"
//输出：false
// 
//
// 示例 2： 
//
// 
//输入：ransomNote = "aa", magazine = "ab"
//输出：false
// 
//
// 示例 3： 
//
// 
//输入：ransomNote = "aa", magazine = "aab"
//输出：true
// 
//
// 
//
// 提示： 
//
// 
// 1 <= ransomNote.length, magazine.length <= 10⁵ 
// ransomNote 和 magazine 由小写英文字母组成 
// 
// Related Topics 哈希表 字符串 计数 👍 196 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	wordCount := make(map[int32]int)
	count := 0
	flag := false
	for _, w := range magazine {
		count, flag = wordCount[w]
		if flag {
			count++
		} else {
			count = 1
		}
		wordCount[w] = count
	}
	for _, w := range ransomNote {
		count, flag = wordCount[w]
		if flag {
			count--
			if count == 0 {
				delete(wordCount, w)
			} else {
				wordCount[w] = count
			}
		} else {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	println(canConstruct("fihjjjjei", "hjibagacbhadfaefdjaeaebgi"))
}
