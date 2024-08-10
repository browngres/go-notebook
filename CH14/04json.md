# 04 JSON

#### Markdown Notes 创建于 2024-08-10T09:26:02.562Z

[json](https://pkg.go.dev/encoding/json@go1.22.6)

## json 的序列化

json 序列化是指，将有 key-value 结构的数据类型(比如结构体、map、切片)序列化成 json 字符串的操作。

-   结构体
    ```go
    //定义一个结构体
    type Monster struct {
        Name string `json:"monster_name"` //反射机制
        Age int `json:"monster_age"`
        Birthday string //....
        Sal float64
        Skill string
    }
    ```
    `data, err := json.Marshal(&monster)  // 结构体 Monster`
    `fmt.Printf("monster 序列化后=%v\n", string(data))`
    对于结构体的序列化，如果我们希望序列化后的 key 的名字，又我们自己重新制定，那么可以给 struct 指定一个 tag 标签。
    tag 可以写`json:"-"`， "-"是忽略的意思
-   map
    `var a map[string]interface{}`
    `data, err := json.Marshal(a)`
-   map 切片
    ```go
    var slice []map[string]interface{}
    var m1 map[string]interface{}
    // map需要 make，然后赋值。略
    slice = append(slice, m1)
    var m2 map[string]interface{}
    // map需要 make，然后赋值。略
    slice = append(slice, m1)
    //将切片进行序列化操作
    data, err := json.Marshal(slice)
    if err != nil {
    fmt.Printf("序列化错误 err=%v\n", err)
    }
    //输出序列化后的结果
    fmt.Printf("slice 序列化后=%v\n", string(data))
    ```
-   还可以对基本类型序列化，但是意义不大
    `var num1 float64 = 2345.67`
    `data, err := json.Marshal(num1)`

### 反序列化

结构体。其他同理

```go
//演示将 json 字符串，反序列化成 struct
//str 在项目开发中，是通过网络传输获取到.. 或者是读取文件获取到
str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
//定义一个 Monster 实例
var monster Monster
err := json.Unmarshal([]byte(str), &monster)
if err != nil {
fmt.Printf("unmarshal err=%v\n", err)
}
fmt.Printf("反序列化后 monster=%v monster.Name=%v \n", monster, monster.Name)
```

1. 在反序列化一个 json 字符串时，要确保反序列化后的数据类型和原来序列化前的数据类型一致。
2. 如果 json 字符串是通过程序获取到的，则不需要再对 “ 转义处理。
