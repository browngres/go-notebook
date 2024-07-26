# 08 闭包 Closure

#### Markdown Notes 创建于 2024-07-26T02:29:31.509Z

## 引入

闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)。
一个拥有许多变量和绑定了这些变量的环境的表达式（通常是一个函数），因而这些变量也是该表达式的一部分。
维基百科讲，闭包（Closure），是引用了自由变量的函数。这个被引用的自由变量将和这个函数一同存在，即使已经离开了创造它的环境也不例外。闭包在运行时可以有多个实例，不同的引用环境和相同的函数组合可以产生不同的实例。

跳过闭包的创建过程直接理解闭包的定义是非常困难的。目前在 JavaScript、Go、PHP、Scala、Scheme、Common Lisp、Smalltalk、Groovy、Ruby、 Python、Lua、objective C、Swift 以及 Java8 以上等语言中都能找到对闭包不同程度的支持。通过支持闭包的语法可以发现一个特点，他们都有垃圾回收(GC)机制。

## 使用 JS 演示

Javascript 应该是普及度比较高的编程语言了，通过这个来举例应该好理解写。看下面的代码，只要关注 script 里方法的定义和调用就可以了。

```html
<script
    src="http://ajax.googleapis.com/ajax/libs/jquery/1.2.6/jquery.min.js"
    type="text/javascript"></script>
<script>
    function a() {
        var i = 0
        function b() {
            console.log(++i)
            document.write("<h1>" + i + "</h1>")
        }
        return b
    }

    $(function () {
        var c = a()
        c()
        c()
        c()
        //a(); //不会有信息输出
        document.write("<h1>=============</h1>")
        var c2 = a()
        c2()
        c2()
    })
</script>
```

这段代码有两个特点：

`函数b`嵌套在`函数a`内部，`函数a`返回`函数b`。这样在执行完 `var c=a()` 后，`变量c`实际上是指向了 `函数b`，再执行`函数c`后就会显示 i 的值，第一次为 1，第二次为 2，第三次为 3，以此类推。 这段代码就创建了一个闭包。当`函数a`的内部`函数b`被`函数a`外的一个变量引用的时候，就创建了一个闭包。
在上面的例子中，由于闭包的存在使得`函数a`返回后，a 中的 i 始终存在，这样每次执行 `c()`，i 都是自加 1 后的值。 从上面可以看出闭包的作用就是在 `a()` 执行完并返回后，使得垃圾回收机制不会收回 a 所占用的资源，因为内部`函数b`的执行需要依赖 a 中的变量 i。

在给定函数被多次调用的过程中，这些私有变量能够保持其持久性。变量的作用域仅限于包含它们的函数，无法从其它程序代码部分进行访问。不过，变量的生存期是可以很长，在一次函数调用期间所创建所生成的值在下次函数调用时仍然存在。正因为这一特点，闭包可以用来完成信息隐藏，并进而应用于需要状态表达的某些编程范型中。

下面来想象另一种情况，如果 a 返回的不是`函数b`，情况就完全不同了。因为 `a()` 执行完后，b 没有被返回给 a 的外界，只是被 a 所引用，而此时 a 也只会被 b 引用，因此函数 a 和 b 互相引用但又不被外界打扰（被外界引用），函数 a 和 b 就会被 GC 回收。所以直接调用`a()`并没有信息输出。

下面来说闭包的另一要素——引用环境。c()跟 c2()引用的是不同的环境，在调用 i++时修改的不是同一个 i，因此两次的输出都是 1。`函数a`每进入一次，就形成了一个新的环境，对应的闭包中，函数都是同一个函数，环境却是引用不同的环境。这和 c()和 c2()的调用顺序都是无关的。

## 用 Go 实现闭包

下面是用 GO 实现刚才的代码。

```go
func a() func() int {
    i := 0
    b := func() int {
        i++
        fmt.Println(i)
        return i
    }
    return b
}

func main() {
    c := a()
    c()
    c()
    c()

    a() //不会输出i
}
```

闭包复制的是原对象指针，这就很容易解释延迟引用现象。

```go

func test() func() {
    x := 100
    fmt.Printf("x (%p) = %d\n", &x, x)

    return func() {
        fmt.Printf("x (%p) = %d\n", &x, x)
    }
}

func main() {
    f := test()
    f()
}
// 输出：
// x (0xc42007c008) = 100
// x (0xc42007c008) = 100
```

test 实际返回的是 FuncVal 对象，其中包含了匿名函数地址、闭包对象指针。当调用匿名函数时，只需以某个寄存器传递该对象即可。

可以这样理解: 闭包是类, 函数是操作，里面的变量 n 是字段。函数和它使用到 n 构成闭包。当我们反复的调用时，因为 n 是初始化一次，因此每调用一次就进行累计。

### 外部引用函数参数局部变量

```go
func add(base int) func(int) int {
    return func(i int) int {
        base += i
        return base
    }
}

func main() {
    tmp1 := add(10)
    fmt.Println(tmp1(1), tmp1(2))
    // 此时tmp1和tmp2不是一个实体了
    tmp2 := add(100)
    fmt.Println(tmp2(1), tmp2(2))
}
```

## 返回两个函数

```go
func test01(base int) (func(int) int, func(int) int) {
    // 定义2个函数，并返回
    // 相加
    add := func(i int) int {
        base += i
        return base
    }
    // 相减
    sub := func(i int) int {
        base -= i
        return base
    }
    // 返回
    return add, sub
}

func main() {
    f1, f2 := test01(10)
    // base一直是没有消
    fmt.Println(f1(1), f2(2))
    // 此时base是9
    fmt.Println(f1(3), f2(4))
}
```
