# 05map

#### Markdown Notes 创建于 2024-08-02T06:49:28.196Z

map 是一种无序的基于 key-value 的数据结构。又称为字段或者关联数组，类似其它编程语言的集合。

-   map 是引用类型，必须初始化才能使用。
-   定义语法：`map[KeyType]ValueType`。KeyType 表示键的类型。ValueType 表示键对应的值的类型。
-   map 类型的变量默认初始值为 nil，需要使用 make 函数来分配内存：
    -   `make(map[KeyType]ValueType, [cap])`
    -   cap 表示 map 的容量。不是必须的，但是我们应该在初始化时指定合适的容量。
    -   如果不指定 initialCapacity，Go 语言会根据实际情况选择一个合适的值。
    -   当 Map 中的键值对数量达到容量时，Map 会自动扩容。
-   value 仍为 map 的 map： `map[string]map[string]string`。value 是 struct 类型的 map，更适合管理复杂的数据(比 value 是 map 更好)。
-   map 是引用类型，遵守引用类型传递的机制，在一个函数接收 map，修改后，会直接修改原来的 map

## 基本使用

```go
scoreMap := make(map[string]int, 8)
scoreMap["张三"] = 90
scoreMap["小明"] = 100
fmt.Println(scoreMap)
fmt.Println(scoreMap["小明"])
fmt.Printf("type:%T\n", scoreMap)

// 声明时填写数据
userInfo := map[string]string{
    "username": "123",
    "password": "123456",
}
```

## key 的类型

key 可以是很多种类型，比如 bool, 数字，string, 指针, channel , 还可以是只包含前面几个类型的 接口, 结构体, 数组。
通常 key 为 int 、string
注意: slice， map 还有 function 不可以，因为这几个没法用 == 来判断

## 判断某个键是否存在

如果 key 存在， ok 为 true，v 为对应的值；不存在， ok 为 false，v 为值类型的零值

```go
v, ok := scoreMap["张三"]
if ok {
    fmt.Println(v)
} else {
    fmt.Println("查无此人")
}
```

map 的 key 不能重复，如果重复了，则以最后的 key-value 为准

## for-range 遍历

map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，遍历时返回的键值对的顺序是不确定的。

```go
for k, v := range scoreMap {
    fmt.Println(k, v)
}
// 可以只遍历key
for k := range scoreMap {
    fmt.Println(k)
}
```

## 按顺序遍历

先用切片接受所有 key，排序切片。再用排序后的 key 去遍历 map

```go
for key := range scoreMap {
    keys = append(keys, key)
}
//对切片进行排序
sort.Strings(keys)
//按照排序后的key遍历map
for _, key := range keys {
    fmt.Println(key, scoreMap[key])
}
```

## 删除键值对

使用内建函数 delete 从 map 中删除一组键值对：`delete(map, key)`
`delete(scoreMap, "小明")//将小明:100从map中删除`
如果 key 不存在，不操作也不报错。
删除所有 key：遍历使用 delete。或者直接 make 一个新的，让原来的成为垃圾，被 gc 回收

## 元素为 map 类型的切片

map 切片（slice of map），这样使 map 个数就可以动态变化了。

```go
var mapSlice = make([]map[string]string, 3)
for index, value := range mapSlice {
    fmt.Printf("index:%d value:%v\n", index, value)
}
fmt.Println("after init")
// 对切片中的map元素进行初始化
mapSlice[0] = make(map[string]string, 10)
mapSlice[0]["name"] = "123"
mapSlice[0]["password"] = "123456"
mapSlice[0]["address"] = "ABC"
for index, value := range mapSlice {
    fmt.Printf("index:%d value:%v\n", index, value)
}
```

## 值为切片类型的 map

```go
var sliceMap = make(map[string][]string, 3)
fmt.Println(sliceMap)
fmt.Println("after init")
key := "中国"
value, ok := sliceMap[key]
if !ok {
    value = make([]string, 0, 2)
}
value = append(value, "北京", "上海")
sliceMap[key] = value
fmt.Println(sliceMap)
```
