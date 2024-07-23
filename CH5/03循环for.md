# 03 循环 for

#### Markdown Notes 创建于 2024-07-23T09:47:53.770Z

For 循环有 3 种形式，只有其中的一种使用分号

```go
for init; condition; post { }
for condition { }
for { }
```

-   init： 一般为赋值表达式，给控制变量赋初值；
-   condition： 关系表达式或逻辑表达式，循环控制条件；
-   post： 一般为赋值表达式，给控制变量增量或减量。
-   `for { }` 等价 `for ; ; {}` 无限循环。通常需要配合 break 语句使用
    for 语句执行过程如下：

1. 先对表达式 init 赋初值；
2. 判别赋值表达式 init 是否满足给定 condition 条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止 for 循环，执行循环体外语句。

```go
for i := 1; i <= 10; i++ {
    fmt.Println("你好！", i)
}
```

```go
s := "abc"

for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
    Println(s[i])
}

n := len(s)
for n > 0 {                // 替代 while (n > 0) {}
    Println(s[n])        // 替代 for (; n > 0;) {}
    n--
}
```

## 循环嵌套

使用循环嵌套来输出 2 到 100 间的素数

```go
func main() {
   /* 定义局部变量 */
   var i, j int

   for i=2; i < 100; i++ {
      for j=2; j <= (i/j); j++ {
         if(i%j==0) {
            break // 如果发现因子，则不是素数
         }
      }
      if(j > (i/j)) {
         fmt.Printf("%d  是素数\n", i)
      }
   }
}
```

## 实现 while 和 do while

-   **while**

```go
循环变量初始化
for {
    if 循环条件表达式 {
        break // 跳出循环
    }
    循环操作语句
    循环变量迭代
}
```

-   **do while**

```go
循环变量初始化
for {
    循环操作语句
    循环变量迭代
    if 循环条件表达式 {
        break // 跳出循环
    }
}
```
