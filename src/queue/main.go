package main

import (
	"fmt"
)

//定义队列中的数据类型
type Element interface{}

//定义队列
type Queue struct {
	element []Element
}

// 创建一个新的队列
func NewQueue() *Queue {
	return &Queue{}
}

//返回当前队列内的元素的个数
func (q *Queue) Size() int {
	return len(q.element)
}

//队列为空则返回True，否则返回False
func (q *Queue) IsEmpty() bool {
	if len(q.element) == 0 {
		return true
	}
	return false
}

//向队列中增加元素
func (q *Queue) Put(item Element) {
	q.element = append(q.element, item)
}

//拿出队列中的元素
func (q *Queue) Get() Element {
	if q.IsEmpty() {
		fmt.Println("queue is empty!")
		return nil
	}

	firstEle := q.element[0]
	q.element = q.element[1:]
	return firstEle
}

//清空队列
func (q *Queue) Clear() bool {
	if q.IsEmpty() {
		fmt.Println("queque is empty!")
		return false
	}
	size := q.Size()
	for i := 0; i < size; i++ {
		q.element[i] = nil
	}
	q.element = nil
	return true
}

func main() {
	queue := NewQueue()
	for i := 0; i < 50; i++ {
		queue.Put(i)
	}
	fmt.Println("size=", queue.Size())
	fmt.Println("最前面的元素=", queue.Get())
	fmt.Println("最前面的元素=", queue.Get())
	//fmt.Println("放进一个元素xixi=", queue.Put(10))
	//往进放一个值
	queue.Put("xixixi")
	size := queue.Size()
	for i := 0; i < size; i++ {
		fmt.Println(queue.Get())
	}
	fmt.Println("这里都弹出了,应该返回false ,看下是不是", queue.Clear())
	fmt.Println("青空之后获取数据,看下获取的是什么=", queue.Get())

}
