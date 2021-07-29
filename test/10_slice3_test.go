package test

import (
	"fmt"
	"testing"
)

/*
append():用于往切片中追加新元素，可以向切片里面追加一个或者多个元素，也可以追加一个切片
copy():不会建立源切片与目标切片之间的联系。也就是两个切片不存在联系，其中一个修改不影响另一个
*/
/*
append()会改变切片所引用的数组的内容，从而影响到引用同一数组的其他切片。当使用append()追加元素到切片时，如果容量不够（也就是(cap-len) == 0），Go就会创建一个新的内存地址来储存元素。
*/
func TestSliceAppend(t *testing.T) {
	var slice = make([]int, 1, 4)
	slice[0] = 222
	fmt.Printf("变量slice --- 值: %v 长度(len):%d 容量(cap): %d  地址: %p \n", slice, len(slice), cap(slice), &slice)

	//向切片中添加一个元素，不超过容量
	slice1 := append(slice, 1)
	fmt.Printf("变量slice1 --- 值: %v 长度(len):%d 容量(cap): %d  地址: %p \n", slice1, len(slice1), cap(slice1), &slice1)

	// 向切片中添加多个元素(不超过容量)
	slice2 := append(slice, 3, 4, 5)
	fmt.Printf("变量slice2 --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n", slice2, len(slice2), cap(slice2), &slice2)

	// 向切片中添加一个切片(超过容量)
	newSlice := []int{1, 2, 3, 4, 5, 6, 7}
	slice3 := append(slice, newSlice...)
	fmt.Printf("变量slice3 --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n", slice3, len(slice3), cap(slice3), &slice3)

	// 调用函数修改
	testSlice(slice1)
	fmt.Printf("调用函数后-变量slice --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n", slice, len(slice), cap(slice), &slice)
	fmt.Printf("调用函数后-变量slice1 --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n", slice1, len(slice1), cap(slice1), &slice1)
	fmt.Printf("调用函数后-变量slice2 --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n",
		slice2, len(slice2), cap(slice2), &slice2)
	fmt.Printf("调用函数后-变量slice3 --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n",
		slice3, len(slice3), cap(slice3), &slice3)
}


/*
copy()
*/
func TestSliceCopy(t *testing.T) {
	// 定义一个切片
	slice := []int{1,2,3,4}
	fmt.Printf("变量slice --- 值: %v 长度(len):%d 容量(cap): %d  地址: %p \n",
		slice,len(slice),cap(slice),&slice )

	// 定义一个切片变量用来复制slice
	copySlice := make([]int,8)

	// 使用copy
	count := copy(copySlice,slice)
	fmt.Printf("复制的数量: %d \n",count)
	fmt.Printf("变量copySlice --- 值: %v 长度(len):%d 容量(cap): %d  地址: %p \n",
		copySlice,len(copySlice),cap(copySlice),&copySlice )

	// 修改复制后的切片，源切片不会变
	copySlice[0] = 100
	fmt.Printf("修改后-->变量slice --- 值: %v 长度(len):%d 容量(cap): %d  地址: %p \n",
		slice,len(slice),cap(slice),&slice )
	fmt.Printf("修改后-->变量copySlice --- 值: %v 长度(len):%d 容量(cap): %d  地址: %p \n",
		copySlice,len(copySlice),cap(copySlice),&copySlice )
}