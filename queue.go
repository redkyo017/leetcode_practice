package main

import "log"

type MyCircularQueue struct {
	Size  int
	Len   int
	Head  int
	Tail  int
	Queue []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	q := MyCircularQueue{
		Size:  k,
		Len:   0,
		Head:  -1,
		Tail:  -1,
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
	if this.Tail == -1 {
		this.Tail = 0
		this.Head = 0
	} else {
		this.Tail = (this.Tail + 1) % this.Size
	}

	this.Queue[this.Tail] = value
	this.Len++
	log.Printf("con co be be :%+v\n", this)
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		log.Println("is empty")
		return false
	}
	if this.Tail == this.Head {
		this.Tail = -1
		this.Head = -1
	} else {
		this.Head = (this.Head + 1) % this.Size
	}
	this.Len--
	log.Printf("Dequeue: %+v\n", this)
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	log.Printf("FRONT: %+v --- %t\n", this, this.IsEmpty())
	if this.IsEmpty() {
		return -1
	}
	return this.Queue[this.Head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	log.Printf("REAR: %+v\n", this)
	if this.IsEmpty() {
		return -1
	}
	return this.Queue[this.Tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	log.Printf("IS EMPTY: %+v\n", this)
	// return this.Len == 0
	return this.Head == -1
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.Head == (this.Tail+1)%this.Size
	// return this.Len == this.Size
}

func init() {
	// obj := Constructor(3)
	// param_1 := obj.EnQueue(1)
	// param_2 := obj.EnQueue(2)
	// param_3 := obj.EnQueue(3)
	// param_4 := obj.EnQueue(4)
	// param_5 := obj.Rear()
	// param_6 := obj.IsFull()
	// param_7 := obj.DeQueue()
	// param_8 := obj.EnQueue(4)
	// param_9 := obj.Rear()
	// obj := Constructor(6)
	// param_1 := obj.EnQueue(6)
	// param_2 := obj.Rear()
	// param_3 := obj.Rear()
	// param_4 := obj.DeQueue()
	// param_5 := obj.EnQueue(5)
	// param_6 := obj.Rear()
	// param_7 := obj.DeQueue()
	// param_8 := obj.Front()
	// param_9 := obj.DeQueue()
	// param_10 := obj.DeQueue()
	// param_11 := obj.DeQueue()
	// obj := Constructor(2)
	// param_1 := obj.EnQueue(4)
	// param_2 := obj.Rear()
	// param_3 := obj.EnQueue(9)
	// param_4 := obj.DeQueue()
	// param_5 := obj.Front()
	// param_6 := obj.DeQueue()
	// param_7 := obj.DeQueue()
	// param_8 := obj.IsEmpty()
	// param_9 := obj.DeQueue()
	// param_10 := obj.EnQueue(6)
	// param_11 := obj.EnQueue(4)

	// log.Println(obj, param_1, param_2, param_3, param_4, param_5, param_6, param_7, param_8, param_9, param_10, param_11)
}

// ["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","Rear","isFull","deQueue","enQueue","Rear"]
// [[3],[1],[2],[3],[4],[],[],[],[4],[]]
// [true,true,true,false,3,true,true,true,4]

// ["MyCircularQueue","enQueue","Rear","Rear","deQueue","enQueue","Rear","deQueue","Front","deQueue","deQueue","deQueue"]
// [[6],[6],[],[],[],[5],[],[],[],[],[],[]]
// [true,6,6,true,true,5,true,-1,false,false,false]

// ["MyCircularQueue","enQueue","Rear","enQueue","deQueue","Front","deQueue","deQueue","isEmpty","deQueue","enQueue","enQueue"]
// [[2],[4],[],[9],[],[],[],[],[],[],[6],[4]]
// [true,6,6,true,true,5,true,-1,false,false,false]
// [true,4,true,true,9,true,false,true,false,true,true]
