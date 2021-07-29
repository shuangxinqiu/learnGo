package main
import "fmt"
func testSlice(slice []int) {
	// 修改切片的值
	slice[0] = 100
}
func main() {
	// 定义一个切片类型，长度为0，容量为4
	var slice = make([]int, 0, 4)

	//向切片中添加1个元素 (不超过容量)
	slice1 := append(slice,1)
	fmt.Printf("变量slice1 --- 值: %v 长度(len):%d 容量(cap): %d  地址: %p \n",
		slice1,len(slice1),cap(slice1),&slice1 )

	// 向切片中添加多个元素(不超过容量)
	slice2 := append(slice,3,4,5)
	fmt.Printf("变量slice2 --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n",
		slice2,len(slice2),cap(slice2),&slice2 )

	slice4 := append(slice,3,4,5,1)
	fmt.Printf("变量slice2 --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n",
		slice4,len(slice4),cap(slice4),&slice4 )

	fmt.Printf("变量slice2 --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n",
		slice,len(slice),cap(slice),&slice )

	// 向切片中添加一个切片(超过容量)
	newSlice := []int{1,2,3,4,5,6,7}
	slice3 := append(slice,newSlice[:]...)
	fmt.Printf("变量slice3 --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n",
		slice3,len(slice3),cap(slice3),&slice3 )

	// 调用函数修改
	testSlice(slice1)
	fmt.Printf("调用函数后-变量slice1 --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n",
		slice1,len(slice1),cap(slice1),&slice1 )
	fmt.Printf("调用函数后-变量slice2 --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n",
		slice2,len(slice2),cap(slice2),&slice2 )
	fmt.Printf("调用函数后-变量slice3 --- 值: %v 长度(len):%d 容量(cap): %d 地址: %p \n",
		slice3,len(slice3),cap(slice3),&slice3 )
}