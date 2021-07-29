package test

import (
	"fmt"
	"testing"
)

/*
切片的长度是切片中元素的数量。切片的容量是从创建切片的索引开始的底层数组中元素的数量。
数组中cap和len相同
*/
func TestSliceLenCap(t *testing.T) {
	slice := []int{1, 2, 4, 6, 8, 10, 12, 14, 16}
	fmt.Printf("切片slice， len:%d cap:%d \n", len(slice), cap(slice))

	//清空切片slice赋值给slice1
	slice1 := slice[0:0]
	fmt.Printf("切片slice1， len:%d cap:%d \n", len(slice1), cap(slice1))

	//使用make指定容量创建切片
	var slice2 = make([]int, 4, 8)
	slice2[1] = 1
	slice2[2] = 3
	fmt.Printf("切片slice2， len:%d cap:%d \n", len(slice2), cap(slice2))

	// 数组计算cap()结果与len()相同
	arr := [...]int{1, 2, 4, 6, 8, 10, 12, 14, 16}
	fmt.Printf("数组arr， len:%d cap:%d \n", len(arr), cap(arr))
}

/*
	切片没有自己的任何数据。它只是底层数组的一个引用。对切片所做的任何修改都将反映在底层数组中。
	数组是值类型，而切片是引用类型。
*/
func TestSliceParam(t *testing.T) {
	var slice = make([]int,4,4)
	slice = []int{1,2,3,4}
	fmt.Printf("变量slice  类型: %T 内存地址: %p 值:%v \n",slice,&slice,slice)

	//调用函数修改
	testSlice(slice)
	fmt.Printf("调用函数后,变量slice  类型: %T 内存地址: %p 值:%v \n",slice,&slice,slice)
}

func testSlice(slice []int) {
	//修改切片的值
	slice[0] = 100
}


/*
多切片共享
修改切片数值，当多个切片共享相同的底层数值时，对每个元素所做的更改将在数值中反映出来
*/
func TestSliceMulti(t *testing.T) {
	var slice = make([]int, 4, 4)
	slice = []int{1, 2, 3, 4}
	fmt.Printf("变量slice  类型: %T 内存地址: %p 值:%v \n", slice, &slice, slice)
	// 复制给另外一个切片
	sliceCopy := slice
	fmt.Printf("变量sliceCopy  类型: %T 内存地址: %p 值:%v \n", sliceCopy, &sliceCopy, sliceCopy)
	// 从切片中截取
	slice2 := slice[1:3]
	fmt.Printf("变量slice2  类型: %T 内存地址: %p 值:%v \n", slice2, &slice2, slice2)
	// 调用函数修改
	testSlice(slice)
	fmt.Printf("调用函数后---变量slice  类型: %T 内存地址: %p 值:%v \n", slice, &slice, slice)
	fmt.Printf("调用函数后---变量sliceCopy  类型: %T 内存地址: %p 值:%v \n", sliceCopy, &sliceCopy, sliceCopy)
	fmt.Printf("调用函数后---变量slice2  类型: %T 内存地址: %p 值:%v \n", slice2, &slice2, slice2)
}