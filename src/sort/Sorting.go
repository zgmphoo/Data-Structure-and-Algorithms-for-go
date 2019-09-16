package main

import "fmt"

// 冒泡排序
func BubbleSort(arr []int) []int {
	l := len(arr) - 1
	for i := 0; i < l; i++ { // 外循环为排序趟数，len个数进行len-1趟
		for j := 0; j < l-i; j++ { //内循环为每趟比较的次数，第i趟比较len-i次
			if arr[j] > arr[j+1] { //相邻元素比较，若逆序则交换（升序为左大于右，降序反之）
				arr[j], arr[j+1] = arr[j+1], arr[j] //快速交换
			}
		}
	}
	return arr
}

//选择排序
func SelectSort(arr []int) []int {
	// 循环的次数
	l := len(arr)
	for i := 0; i < l-1; i++ {
		min := arr[i]	//初始化最小值
		minIndex := i 	//记录最小值的下标
		for j := i+1; j < l; j++ {
			if min > arr[j] {
				min = arr[j]
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = min, arr[i]
	}
	return arr
}


// 

func main() {
	var slice1 = []int{44, 80, 78, 77, 42}
	fmt.Println(BubbleSort(slice1))
	var slice2 = []int{44, 80, 78, 77, 42}
	fmt.Println(SelectSort(slice2))
}
