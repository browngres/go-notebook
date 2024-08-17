package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type Person struct {
	Id   int `json:"id" db:"pk"`
	Name string
	Age  int
}

// 匿名字段
type Student struct {
	Person
	Grade string
	Class string
}

// 绑方法
func (p Person) Say() {
	fmt.Println("Hello, I'm ", p.Name)
}

// any关键字 在任何时候都等同于 interface{}
func Struct_reflect1(o interface{}) {
	fmt.Println("==========字段和值==========")
	t := reflect.TypeOf(o)
	fmt.Println("类型 Type：", t)
	fmt.Println("类型名 Name：", t.Name())
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println("值 Value：", v)
	// 获取所有属性
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("字段%s : %v   ", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		// 此时已经是 any 类型，如果相当于变量可以使用。
		// 如需将 any 转换为具体类型如 int，应该使用类型断言：val.(int)。
		// 这里可以省略该步骤，因为Println可以接受 any 类型
		fmt.Println("字段值:", val)
	}
	// 获取字段id的 tag
	f_id := t.Field(0)
	fmt.Println("id Tag json: ", f_id.Tag.Get("json"))
	fmt.Println("id Tag db", f_id.Tag.Get("db"))

	fmt.Println("============方法============")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println("方法名Name：", m.Name)
		fmt.Println("方法类型Type:", m.Type)
	}
}

func Struct_reflect2(o any) {
	t := reflect.TypeOf(o)
	fmt.Println("类型Type：", t)
	fmt.Printf("字段Field：%#v\n", t.Field(0)) // Anonymous
	fmt.Printf("字段值：%#v\n", reflect.ValueOf(o).Field(0))
}

func Struct_reflect3(o any) {
	v := reflect.ValueOf(o)
	// 获取指针指向的元素
	v = v.Elem()
	// 取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("person123")
	}
}

func Struct_reflect4(o any) {
	v := reflect.ValueOf(o)
	// 获取方法
	m := v.MethodByName("Say")
	// 构建参数
	var args []reflect.Value // 没有参数
	// args := []reflect.Value{reflect.ValueOf("6666")} //带参数的方法
	m.Call(args)
}

func main() {
	p1 := Person{1, "person1", 20}
	s1 := Student{p1, "grade1", "class1"}
	Struct_reflect1(p1)

	fmt.Println("============匿名字段============")
	Struct_reflect2(s1)

	fmt.Println("============修改结构体============")
	Struct_reflect3(&s1)
	fmt.Println(s1)

	fmt.Println("============调用方法============")
	Struct_reflect4(s1)

}
