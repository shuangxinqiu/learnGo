package test

import (
	"fmt"
	"testing"
)

/*
闭包是由函数和与其相关的引用环境组合而成的实体。在实现深约束时，需要创建一个能显示表示引用环境的东西，并将它与相关的子程序捆绑在一起，这样捆绑起来的整体被称为闭包。函数+引用环境=闭包。
*/

/*
闭包和函数的区别
闭包只是形式和表现上像函数，但实际上不是函数。具体区别如下：
	函数运行时只有一个实例，函数体被定义后就确定了，不会在执行时发生变化
	闭包在运行时可以有多个实例，不同的引用环境和相同的函数组合可以产生不同的实例
	函数本身不存储任何信息，只有与引用环境结合后形成的闭包才具有记忆性
	函数是编译器静态的概念，而闭包是运行期动态的概念
*/
func counterByClosure() func(i int) int {
	sum := 0
	return func(i int) int {
		sum = sum + i
		return sum
	}
}

func TestClosure(t *testing.T) {
	count := counterByClosure()
	for i := 1; i <= 5; i++ {
		fmt.Println(count(i))
	}

	for i := 6; i <= 10; i++ {
		fmt.Println(count(i))
	}
}
