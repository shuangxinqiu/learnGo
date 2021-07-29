package main

import (
	"fmt"
	"reflect"
	"testing"
)

type i interface {

}

type person struct {
	Name string `json:"name" bson:"b_name"`
	Age int `json:"age" bson:"b_age"`
}

//person实现了fmt.Stringer接口
//func (p person) String() string {
//	return fmt.Sprintf("Name is %s,Age is %d",p.Name,p.Age)
//}

func (p person) Print(prefix string)  {
	fmt.Printf("%s:Name is %s,Age is %d\n",prefix,p.Name,p.Age)
}

func TestPointer(t *testing.T) {
	p := person{
		Name: "feixue",
		Age:  22,
	}
	pv := reflect.ValueOf(p)
	//反射调用person的Print方法
	mPrint := pv.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("登录")}
	mPrint.Call(args)
}

func modifyMap(p map[string]int)  {
	fmt.Printf("modifyMap函数：p的内存地址为%p\n",p)

	p["feixue"] = 22
}