package test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	//使用make() 函数来创建切片
	var slice1 = make([]int, 3)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 4
	fmt.Printf("通过make关键字[slice1] 类型: %T 值: %v \n", slice1, slice1)

	//直接初始化使用
	slice2 := []int{1, 2, 4}
	fmt.Printf("通过make关键字[slice2] 类型: %T 值: %v \n", slice2, slice2)

	//通过数组截取
	arr := [...]int{1, 2, 4, 6, 8, 10, 12, 14, 16}
	//从索引0开始截取到索引为4（不包括4）
	slice3 := arr[0:4]
	fmt.Printf("从索引0开始截取到索引为4(不包括4)[slice3] 类型:%T 值:%v \n", slice3, slice3)
	//索引从0开始时，0也可以省略不写
	slice33 := arr[:4]
	fmt.Printf("从索引0开始截取到索引为4(不包括4)[slice33] 类型:%T 值:%v \n", slice33, slice33)
	//从索引4开始截取到最后
	slice4 := arr[5:]
	fmt.Printf("从索引5开始截取到最后[slice4] 类型:%T 值:%v \n", slice4, slice4)
}

//删除
func TestSliceDel(t *testing.T) {
	//定义一个切片
	slice := []int{1, 3, 5, 3, 2}
	fmt.Printf("切片slice--> 值:%v len: %d cap: %d\n", slice, len(slice), cap(slice))
	//删除第一个元素
	slice = slice[1:]
	fmt.Printf("删除第一个元素后--> 值:%v len: %d cap: %d\n", slice, len(slice), cap(slice))

	//删除最后一个元素
	slice = slice[:len(slice)-1]
	fmt.Printf("删除最后一个元素后--> 值:%v len: %d cap: %d\n", slice, len(slice), cap(slice))
}

//删除指定位置元素
func TestSliceIndexDel(t *testing.T) {
	slice := []int{9, 8, 7, 5, 2}
	index := 2
	front := slice[:index]
	behind := slice[index+1:]
	slice = append(front, behind...)
	fmt.Printf("删除索引%d后--> 值:%v len: %d cap: %d\n", index, slice, len(slice), cap(slice))

}
