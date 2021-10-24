package main

/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
换句话说，第一个字符串的排列之一是第二个字符串的 子串 。

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

滑动窗口
题意理解：由于变位词不会改变字符串中每个字符的个数，所以只有当两个字符串每个字符的个数均相等时，一个字符串才是另一个字符串的变位词
窗口大小为s1的长度
如何比较：
	利用数组比较，来避免s1与滑动窗口的每个字符数量相等性比较的麻烦(也可以优化为cnt数组与空数组的比较)
	不适合位运算，因为单词可能重复多次
优化：
	注意到每次窗口滑动时，只统计了一进一出两个字符，却比较了整个cnt1和cnt2数组。
	可以用一个变量 diff 来记录cnt1和cnt2的不同值的个数，这样判断cnt1和cnt2是否相等就转换成了判断 diff 是否为 0.
	每次窗口滑动，记一进一出两个字符为 x 和 y.
	若 x=y 则对 \cnt2无影响，可以直接跳过。
	若 x!=y，对于字符 x，在修改 cnt2之前若有cnt2[x]=cnt1[x]，则将 diff 加一；
	在修改cnt2之后若有 cnt2[x]=cnt1[x]，则将 diff 减一。字符 y 同理。

	此外，为简化上述逻辑，我们可以只用一个数组 cnt，其中 cnt[x]=cnt2[x]−cnt1[x]，将cnt1[x] 与 cnt2[x] 的比较替换成cnt[x] 与 0 的比较。

方法二：双指针
	for right, ch := range s2 {
        x := ch - 'a'
        cnt[x]++
        for cnt[x] > 0 {			// 此时进入了一个新字符，要left移动，直到=0
            cnt[s2[left]-'a']--
            left++
        }
        if right-left+1 == n {
            return true
        }
    }
*/

func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	var cnt1, cnt2 [26]int // 标记每个字符的数量
	for i, char := range s1 {
		cnt1[char-'a']++
		cnt2[s2[i]-'a']++ // s2初始化前n个
	}
	for cnt1 == cnt2 {
		return true
	}
	for i := m; i < n; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-m]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

func checkInclusion2(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	var cnt [26]int // 标记每个字符的数量
	for i, char := range s1 {
		cnt[char-'a']--
		cnt[s2[i]-'a']++ // s2初始化前n个
	}
	diff := 0
	for _, count := range cnt {
		if count != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := m; i < n; i++ {
		x, y := s2[i]-'a', s2[i-m]-'a'
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}
