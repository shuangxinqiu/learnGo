package test

import (
	"fmt"
	"testing"
)

/*
map是一种集合，可以像遍历数组或切片那样去遍历它。因为map是由Hash表实现的，所以对map的读取顺序不固定。
*/
/*
1. map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
2. map的长度是不固定的，和切片一样可以扩展。内置的len()函数同样适用于map，返回map拥有的键值对的数量
3. 同一个map中key必须保证唯一
4. key的数据类型必须是可参与比较运算的类型，也就是支持==或!=操作的类型。
5. 引用类型则不能作为key的数据类型
6. map的value可以是任何数据类型
7. map和切片一样，也是一种引用类型
*/

/*
可以用var定义map，也可以用内建函数make()
var关键字声明map，未初始化的map默认值是nil，不能存放键值对。如果要使用map存储键值对，必须在声明时初始化，或者使用make函数分配到内存空间
*/
func TestMap(t *testing.T) {
	//声明的同时初始化
	var ageMap = map[string]int{
		"zhangsan": 20,
		"lisi":     33,
		"wangwu":   55,
	}
	fmt.Printf("变量ageMap--> 值: %v 类型: %T \n", ageMap, ageMap)

	//短变量声明初始化
	ageMap2 := map[string]int{
		"zhangsan": 20,
		"lisi":     33,
		"wangwu":   55,
	}
	fmt.Printf("变量ageMap2--> 值: %v 类型: %T \n", ageMap2, ageMap2)

	var ageMap3 = map[string]int{}
	ageMap3["a"] = 11
	fmt.Println(ageMap3)
}

func TestMapMake(t *testing.T) {
	//先创建后赋值
	ageMap := make(map[string]int)
	ageMap["zhangsan"] = 22
	ageMap["lisi"] = 33
	ageMap["wangwu"] = 55
	fmt.Printf("变量ageMap-- 值: %v 类型: %T", ageMap, ageMap)

}

//遍历基础map
func TestRangeMap(t *testing.T) {
	ageMap := map[string]int{"zhangsan": 33, "lisi": 44, "wangwu": 55}
	for key, value := range ageMap {
		fmt.Printf("k-->: %v v--> %v \n", key, value)
	}
}

//遍历嵌套map
func TestRangeMoreMap(t *testing.T) {
	//声明嵌套map
	productMap := map[string]map[string]float32{
		"fruit": {"banana": 3.22, "apple": 1.11, "orange": 4.44},
		"phone": {"sungsam": 123.2, "oppo": 1.4, "vivo": 11},
		"sport": {"basketball": 111, "football": 444, "tennis": 123123},
	}

	for class, product := range productMap {
		for name, price := range product {
			fmt.Printf("分类：%s 产品: %s 价格: %0.2f \n", class, name, price)
		}
	}
}

//操作map
/*
判断key是否存在
通过value, ok := map[key] 获知key/value是否存在。
ok是bool型，如果ok是true，则该键值对存在，否则不存在
*/
func TestMapKeyExist(t *testing.T) {
	fruitMap := map[string]float32{"banana": 123, "apple": 23, "juice": 55}
	price, ok := fruitMap["apple"]
	if ok {
		fmt.Printf("apple is exist and the price is %0.2f \n", price)
	}

	price2, ok2 := fruitMap["西瓜"]
	if ok2 {
		fmt.Printf("西瓜 is exist and the price is %0.2f \n", price2)
	} else {
		fmt.Printf("西瓜 is not exist \n")
	}

	//简写
	if price3, ok3 := fruitMap["banana"]; ok3 {
		fmt.Printf("banana is exist and the price is %0.2f \n", price3)
	}
}

/*
删除
delete(map,key) 函数用于删除集合的某个元素，参数为map和其对应的key。删除函数不返回任何值。
*/
func TestDeleteMap(t *testing.T)  {
	fruitMap := map[string]float32{"香蕉": 3.22, "苹果": 1.88, "葡萄": 2.49,"梨":4.13}
	fmt.Printf("删除前-->fruitMap = %v \n",fruitMap)

	delete(fruitMap,"香蕉")
	fmt.Printf("删除后-->fruitMap = %v \n",fruitMap)

	//清空map
	fruitMap = map[string]float32{}
	fmt.Printf("清空后-->fruitMap = %v \n",fruitMap)

}

/*
map是引用类型
map和切片相似，都是引用类型。将一个map赋值给一个新的变量时，它们指向同一块内存（底层数据结构）。因此，修改两个变量的内容都能够引起它们指向的数据发生变化。
*/
func TestMap5(t *testing.T)  {
	//声明map
	fruitMap := map[string]float32{"香蕉": 3.22, "苹果": 1.88, "葡萄": 2.49,"梨":4.13}
	fmt.Printf("调用函数前-->fruitMap = %v \n",fruitMap)

	//调用函数
	testMap(fruitMap)
	fmt.Printf("调用函数后-->fruitMap = %v \n",fruitMap)
}
// map是引用类型，函数改变他的值外层变量也会变
func testMap(fruitMap map[string]float32) {
	fruitMap["香蕉"] = 5.99
}
