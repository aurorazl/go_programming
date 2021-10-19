package main

/*
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

全员进行异或操作即可。考虑异或操作的性质：对于两个操作数的每一位，相同结果为 0，不同结果为 1。

先对所有数字进行一次异或，得到两个出现一次的数字的异或值。
在异或结果中找到任意为 1 的位。
根据这一位对所有的数字进行分组。
在每个组内进行异或操作，得到两个数字。
*/
