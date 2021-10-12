package LeetCode_热题_HOT_100

/*
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
输入: 2
输出: [0,1,1]
进阶:
	给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
	要求算法的空间复杂度为O(n)。
	你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的__builtin_popcount）来执行此操作。
x=x & (x−1)，该运算将 x 的二进制表示的最后一个 1 变成 0
对 x 重复该操作，直到 x 变成 0，则操作次数即为 x 的「一比特数」。
对于给定的 n，计算从 0 到 n 的每个整数的「一比特数」的时间都不会超过 O(logn)，因此总时间复杂度为 O(nlogn)。

*/
func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}

/*
对于正整数 xx，如果可以知道最大的正整数 yy，使得 y \le xy≤x 且 yy 是 22 的整数次幂，则 yy 的二进制表示中只有最高位是 11，其余都是 00，
此时称 yy 为 xx 的「最高有效位」
如果 i&(i−1)=0，则令 highBit=i，更新当前的最高有效位。
i 比 i−highBit 的「一比特数」多 1，由于是从小到大遍历每个整数，因此遍历到 ii 时，i−highBit 的「一比特数」已知，令 bits[i]=bits[i−highBit]+1。

*/
func countBits2(n int) []int {
	bits := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}
