package test

import (
	"fmt"
	"testing"
)

func TestArr(t *testing.T) {
	//先声明后使用
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Printf("arr1 type: %T val: %v ", arr1, arr1)
	fmt.Println()

	//直接初始化使用
	arr2 := [3]int{1, 3, 5}
	fmt.Printf("arr2 type: %T val: %v ", arr2, arr2)
	fmt.Println()

	//声明数组（不指定长度）
	arr3 := [...]int{1,3,5,7,9}
	fmt.Printf("arr3 type: %T val: %v ",arr3,arr3)
}

//遍历二维数组
func TestArr2(t *testing.T) {
	arr := [3][4]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
	}

	for key,val := range arr{
		for k,v := range val {
			fmt.Printf("[%d][%d] is %d \n",key,k,v)
		}
	}
}
