# 01 反射 reflect

#### Markdown Notes 创建于 2024-08-17T13:44:46.423Z

反射是指在程序运行期对程序本身进行访问和修改的能力。
(以前用到的结构体 json tag 就是反射)

-   变量的内在机制

    变量包含类型信息和值信息 `var arr [10]int` `arr[0] = 10`
    类型信息：是静态的元信息，是预先定义好的。
    值信息：是程序运行过程中动态改变的。

## 反射的使用

`import ("reflect")`

-   反射可以在运行时动态获取变量的各种信息, 比如变量的类型(type)，类别(kind)。
-   如果是结构体变量，还可以获取到结构体本身的信息(包括字段、方法)。
-   通过反射可以修改变量的值，可以调用关联的方法。

获取类型信息：`reflect.TypeOf`，是静态的
获取值信息：`reflect.ValueOf`，是动态的

[`type Type` 全部方法](https://pkg.go.dev/reflect@go1.22.6#Type)
[`type Value` 全部方法](https://pkg.go.dev/reflect@go1.22.6#Value)
也就是说可以通过一个变量拿到很多东西，而这些东西不仅提供了大量方法，还可以互相转化。
比如 Value 的 `func (v Value) NumField() int` 方法就可以拿到结构体有多少个字段。
拿到的东西甚至可以操作原来的变量。
Kind 代表 Type 的具体类型
[`type Kind`](https://pkg.go.dev/reflect@go1.22.6#Kind)

```go
//反射获取interface类型信息

func reflect_type(a interface{}) {
   t := reflect.TypeOf(a)
   fmt.Println("类型是：", t)
   // kind()可以获取具体类型
   k := t.Kind()
   fmt.Println(k)
   switch k {
   case reflect.Float64:
      fmt.Printf("a is float64\n")
   case reflect.String:
      fmt.Println("string")
   }
}
//反射获取interface值信息
func reflect_value(a interface{}) {
    v := reflect.ValueOf(a)
    fmt.Println(v)
    k := v.Kind()
    fmt.Println(k)
    switch k {
    case reflect.Float64:
        fmt.Println("a是：", v.Float())
    }
}

func main() {
   var x float64 = 3.4
   reflect_type(x)
    reflect_value(x)
}
```

```go
//反射修改值
func reflect_set_value(a interface{}) {
    v := reflect.ValueOf(a)
    k := v.Kind()
    switch k {
    case reflect.Float64:
        // 反射修改值
        v.SetFloat(6.9)
        fmt.Println("a is ", v.Float())
    case reflect.Ptr:
        // Elem()获取地址指向的值
        v.Elem().SetFloat(7.9)
        fmt.Println("case:", v.Elem().Float())
        // 地址
        fmt.Println(v.Pointer())
    }
}

func main() {
    var x float64 = 3.4
    // 反射认为下面是指针类型，不是float类型
    reflect_set_value(&x)
    fmt.Println("main:", x)
}
```

## 反射操作结构体

详见代码
可以获取所有属性，i 是顺序。 `f := t.Field(i)`
获取结构体字段个数：`t.NumField()`
获取字段的值信息： `v.Field(i).Interface()`
获取方法：`t.Method(i)` `fmt.Println(m.Name)` `fmt.Println(m.Type)`
可以查看匿名字段
修改结构体值：`f := v.FieldByName("Name")` `f.SetString("123")`
获取方法：`m := v.MethodByName("Hello")`
构建参数并调用方法：`args := []reflect.Value{reflect.ValueOf("6666")}` ` m.Call(args)`
获取字段的 tag：`f.Tag.Get("json")`

### 应用场景：编写程序框架

（类似于 django 的 dispatch 函数）
定义两个匿名函数 test1 和 test2，各自的形参数量都不同。使用一个 bridge 用做统一处理接口。

```go
bridge := func(call interface{}, args ...interface{}) {
    //处理内容
}
bridge(test1,1,2)  // 实现调用test1
bridge(test2,1,2,"abc")  // 实现调用test2
```
