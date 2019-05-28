package main

import "log"

type MyCircularQueue struct {
	Size  int
	Head  int
	Tail  int
	Queue []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	q := MyCircularQueue{
		Size:  k,
		Head:  0,
		Tail:  0,
		Queue: make([]int, k),
	}
	return q
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		log.Println("is fulled")
		return false
	}
	this.Queue[this.Tail] = value

	this.Tail = (this.Tail + 1) % this.Size
	log.Printf("con co be be :%+v\n", this)
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		log.Println("is empty")
		return false
	}
	this.Head = (this.Head + 1) % this.Size
	log.Printf("con heo :%+v\n", this)
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	return this.Queue[this.Head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	return this.Queue[this.Tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.Head == this.Tail
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {

	return this.Head == (this.Tail+1)%this.Size
}

func init() {
	obj := Constructor(3)
	log.Println("init: ", obj)
	param_1 := obj.EnQueue(1)
	param_2 := obj.EnQueue(2)
	param_3 := obj.EnQueue(3)
	param_4 := obj.EnQueue(4)
	param_5 := obj.Rear()
	param_6 := obj.IsFull()
	param_7 := obj.DeQueue()
	param_8 := obj.EnQueue(4)
	param_9 := obj.Rear()

	log.Println(obj, param_1, param_2, param_3, param_4, param_5, param_6, param_7, param_8, param_9)
}

// ["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","Rear","isFull","deQueue","enQueue","Rear"]
// [[3],[1],[2],[3],[4],[],[],[],[4],[]]
// [true,true,true,false,3,true,true,true,4]
