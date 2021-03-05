package leetcode

type arrayStack struct {
	stack []int
}

func (as *arrayStack) push(n int) {
	as.stack = append(as.stack, n)
}

func newArrayStack() *arrayStack {
	return &arrayStack{
		stack: make([]int, 0),
	}
}

func (as *arrayStack) pop() int {
	if len(as.stack) == 0 {
		return 0
	}

	top := as.stack[len(as.stack)-1]
	as.stack = as.stack[:len(as.stack)-1]
	return top
}

func (as arrayStack) top() int {
	if len(as.stack) == 0 {
		return 0
	}

	return as.stack[len(as.stack)-1]
}

func (as arrayStack) len() int {
	return len(as.stack)
}

type MyQueue struct {
	// order means pop is queue order
	order *arrayStack
	// reverse means pop is stack order
	reverse *arrayStack
}

/** Initialize your data structure here. */
func Constructor3() MyQueue {
	return MyQueue{
		order:   newArrayStack(),
		reverse: newArrayStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.reverse.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.order.len() > 0 {
		return this.order.pop()
	}

	for this.reverse.len() > 1 {
		this.order.push(this.reverse.pop())
	}

	return this.reverse.pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.order.len() > 0 {
		return this.order.top()
	}

	for this.reverse.len() > 0 {
		this.order.push(this.reverse.pop())
	}

	return this.order.top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.order.len() == 0 && this.reverse.len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
