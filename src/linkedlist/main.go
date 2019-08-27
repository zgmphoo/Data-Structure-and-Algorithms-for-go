package main

import (
	"fmt"
)

//节点值的类型
type ElementType interface{}

//定义节点
type Node struct {
	value interface{}
	next  *Node
}

//获取节点的值
func (node *Node) GetValue() ElementType {
	return node.value
}

//获取节点的下一个节点
func (node *Node) GetNext() *Node {
	return node.next
}

//生成Node
func NewNode(x ElementType) *Node {
	return &Node{x, nil}
}

//定义链表
type LinkedList struct {
	head     *Node
	tailnode *Node
	length   int
}

//生成链表
func NewLinkedList() *LinkedList {
	var tailnode *Node = nil
	head := &Node{nil, tailnode} //默认 node 都是 nil
	return &LinkedList{
		head:     head,
		tailnode: tailnode,
		length:   0,
	}
}

//获取链表的长度
func (list *LinkedList) Length() int {
	return list.length
}

//增加,时间复杂度O(1)
func (list *LinkedList) Append(val ElementType) {
	node := NewNode(val)
	if list.tailnode == nil {
		list.head.next = node
	} else {
		list.tailnode.next = node
	}
	list.tailnode = node
	list.length += 1
}

//左侧增加,时间复杂度O(1)
func (list *LinkedList) AppendLeft(val ElementType) {
	node := NewNode(val)
	if list.tailnode == nil {
		list.head.next = node
		list.tailnode = node
	} else {
		curnode := list.head.next
		node.next = curnode
		list.head.next = node
	}
	list.length += 1
}

//遍历链表,时间复杂度O(n)
func (list *LinkedList) PrintLink() {
	if list.length == 0 {
		fmt.Println("LinkedList is empty!")
		return
	}
	le := list.length
	node := list.head.next
	for i := 0; i < le; i++ {
		fmt.Printf("链表第%d个值为=", i)
		fmt.Println(node.value)
		node = node.next
	}
}

//删除节点,时间复杂度O(n)
func (list *LinkedList) Remove(val ElementType) bool {
	if list.length == 0 {
		fmt.Println("LinkedList is empty!")
		return false
	}
	privnode := list.head
	for {
		if privnode.next.next == nil {
			if privnode.next.value == val {
				//删除节点
				privnode.next = nil
				list.length -= 1
				return true
			}
			return false
		} else {
			if privnode.next.value == val {
				privnode.next = privnode.next.next
				list.length -= 1
				return true
			} else {
				privnode = privnode.next
			}
		}
	}

}

//查找元素,时间复杂度O(n)
func (list *LinkedList) Find(val ElementType) int {
	if list.length == 0 {
		fmt.Println("LinkedList is empty!")
		return -1
	}
	le := list.length
	node := list.head.next
	for i := 0; i < le; i++ {
		if node.value == val {
			return i
		}
		node = node.next
	}
	return -1
}

//清空链表
func (list *LinkedList) Clear() {
	if list.length == 0 {
		fmt.Println("链表本来就是空的!")
	}
	var tailnode *Node = nil
	head := &Node{nil, tailnode} //默认 node 都是 nil
	list.head = head
	list.length = 0
	list.tailnode = tailnode
}

//翻转链表
func (list *LinkedList) Reverse() {
	var privnode *Node = nil
	var nextnode *Node
	if list.length == 0 {
		fmt.Println("链表中没有节点,无法反转!")
	}
	curnode := list.head.next
	list.tailnode = curnode
	for {
		nextnode = curnode.next
		curnode.next = privnode
		if nextnode == nil {
			list.head.next = curnode
			break
		}
		privnode = curnode
		curnode = nextnode
	}

}

func main() {
	//test
	link := NewLinkedList()
	link.Append("one")
	link.Append("two")
	link.Append("three")
	fmt.Println("链表长度＝", link.Length())
	link.PrintLink()
	link.AppendLeft(3)
	link.AppendLeft(2)
	link.AppendLeft(1)
	fmt.Println("链表长度＝", link.Length())
	link.PrintLink()
	fmt.Println("链表删除three，是否成功：", link.Remove("three"))
	fmt.Println("链表删除２，是否成功：", link.Remove(2))
	fmt.Println("链表删除xixi，是否成功：", link.Remove("xixi"))
	fmt.Println("链表长度＝", link.Length())
	link.PrintLink()
	fmt.Println("链表查找＝", link.Find("one"))
	fmt.Println("链表查找＝", link.Find(8))
	fmt.Println("翻转链表")
	link.Reverse()
	link.PrintLink()
	fmt.Println("清空链表")
	link.Clear()
	fmt.Println("清空之后的链表长度=", link.Length())
}
