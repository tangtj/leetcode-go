package leetcode

type MyQueue struct {
	queue []int
}

/** Initialize your data structure here. */
func MyQueueConstructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	x := this.queue[0]
	this.queue = this.queue[1:]
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	x := this.queue[0]
	return x
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.queue) == 0
}
