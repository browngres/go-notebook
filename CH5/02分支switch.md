# 02 分支 switch

#### Markdown Notes 创建于 2024-07-23T08:45:33.355Z

switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。switch 分支表达式可以是任意类型，不限于常量。
默认相当于每个 case 最后带有 break，匹配成功后不会自动向下执行其他 case，而是跳出整个 switch。

```go
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。 （常量、变量、一个有返回值的函数等都可以）
您可以同时测试多个可能符合条件的值，使用逗号分割，例如：case val1, val2, val3。

```go
switch marks {
    case 90: grade = "A"
    case 80: grade = "B"
    case 50,60,70 : grade = "C"
    default: grade = "D"
}

switch {
    case grade == "A" :
        fmt.Printf("优秀!\n" )
    case grade == "B", grade == "C" :
        fmt.Printf("良好\n" )
    case grade == "D" :
        fmt.Printf("及格\n" )
    case grade == "F":
        fmt.Printf("不及格\n" )
    default:
        fmt.Printf("差\n" )
}
```

```go
var k = 0
switch k {
case 0:
    Println("fallthrough")
    fallthrough
// 可以使用fallthrough强制执行后面的case代码。
case 1:
    fmt.Println("1")
case 2:
    fmt.Println("2")
default:
    fmt.Println("def")
}
```

```go
var n = 0
switch { //省略条件表达式，可当 if...else if...else
case n > 0 && n < 10:
    fmt.Println("i > 0 and i < 10")
case n > 10 && n < 20:
    fmt.Println("i > 10 and i < 20")
default:
    fmt.Println("def")
}
```

## 注意

-   case 后面的表达式如果是常量值(字面量)，则要求不能重复。
-   case 后面不需要带 break , 程序匹配到一个 case 后就会执行对应的代码块，然后退出 switch，如果一个都匹配不到，则执行 default。
-   default 语句不是必须的。
-   switch 可以不带表达式，类似 if else 来使用，它会匹配 true。
-   穿透 fallthrough ，如果在 case 语句块后增加 fallthrough ,则会继续执行下一个 case。（不判断，直接执行下一层）。但是这种用法可以直接写逗号两个表达式即可。

## Type Switch

switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。

```go
switch x.(type){
    case type:
       statement(s)
    case type:
       statement(s)
    /* 你可以定义任意个数的case */
    default: /* 可选 */
       statement(s)
}
```

```go
var x interface{}
switch i := x.(type) { // 带初始化语句
case nil:
    fmt.Printf(" x 的类型 :%T\r\n", i)
case int:
    fmt.Printf("x 是 int 型")
case float64:
    fmt.Printf("x 是 float64 型")
case func(int) float64:
    fmt.Printf("x 是 func(int) 型")
case bool, string:
    fmt.Printf("x 是 bool 或 string 型")
default:
    fmt.Printf("未知型")
}
```
