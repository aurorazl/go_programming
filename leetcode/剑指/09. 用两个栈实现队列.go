package main

/*
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。
(若队列中没有元素，deleteHead操作返回 -1 )

stack_in只负责进入
stack_out只负责取出
只有stack_out为空时才把stack_in的所有元素倾倒进stack_out中，这样顺序就不会乱了

*/
