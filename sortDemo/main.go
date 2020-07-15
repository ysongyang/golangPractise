package main

//排序案例 选择排序、插入排序、快速排序
import (
	"fmt"
	"math/rand"
	"time"
)

//选择排序案例
func selectSort(arr *[800000]int) {

	//循环 len(*arr)-1 次，进行排序
	for index := 0; index < len(*arr)-1; index++ {
		//先定义第一个元素为最大值
		max := (*arr)[index]
		//先定义最大元素的坐标为0
		maxIndex := index
		for i := index + 1; i < len(*arr); i++ {
			if max < (*arr)[i] {
				//如果找到最大值
				max = (*arr)[i]
				//最大值元素的坐标
				maxIndex = i
			}
		}
		if maxIndex != index {
			//元素交换
			(*arr)[index], (*arr)[maxIndex] = (*arr)[maxIndex], (*arr)[index]
		}
		//fmt.Printf("第%d次交换 %v\n", index+1, *arr)
	}

}

//插入排序
func insertSort(arr *[800000]int) {
	//循环 len(*arr)-1 次，进行排序
	for index := 1; index < len(*arr); index++ {
		//第一次先给第二个元素找到合适的位置并插入
		insertVal := (*arr)[index]
		insertIndex := index - 1
		for insertIndex >= 0 && (*arr)[insertIndex] < insertVal {
			(*arr)[insertIndex+1] = (*arr)[insertIndex] //数据后移
			insertIndex--
		}
		//插入数据
		if insertIndex+1 != index {
			(*arr)[insertIndex+1] = insertVal //数据前移
		}
		//fmt.Printf("第%d次插入 %v\n", index, *arr)
	}

}

//快速排序法
/**
* @param left 数组左边的下标
* @param right 数组右边的下标
* @param arr 数组
 */
func quickSort(left int, right int, arr *[800000]int) {
	l := left
	r := right
	//获取数组中轴的下标
	pivot := (*arr)[(left+right)/2]
	temp := 0
	//比pivot 小的数放左边
	//比pivot 打的数放右边
	for ; l < r; {
		//先从pivot 左边找到大于 pivot的值
		for ; (*arr)[l] < pivot; { //<升序
			l++
		}
		//从pivot 右边找到小于 pivot的值
		for ; (*arr)[r] > pivot; { // > 升序
			r--
		}
		if l >= r {
			break
		}
		temp = (*arr)[l]
		(*arr)[l] = (*arr)[r]
		(*arr)[r] = temp

		if (*arr)[l] == pivot {
			r--
		}
		if (*arr)[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		go quickSort(left, r, arr)
	}
	//向右递归
	if right > l {
		go quickSort(l, right, arr)
	}
}

func main() {
	//执行时间
	t := time.Now()
	var arr [800000]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 800000; i++ {
		arr[i] = rand.Intn(900000)
	}
	//fmt.Println("原始数组：", arr)
	//selectSort(&arr)  //app elapsed: 3.1187062s
	//fmt.Println("选择排序：", arr)
	//insertSort(&arr)  //app elapsed: 860.2854ms
	//fmt.Println("插入排序：", arr)
	quickSort(0, len(arr)-1, &arr) //app elapsed: 7.5459ms

	//快速排序速度最快

	//fmt.Println("快速排序：", arr)
	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)
}
