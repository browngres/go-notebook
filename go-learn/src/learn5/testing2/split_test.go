package split

import (
	"reflect"
	"testing"
)

func TestSplit1(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split1("a:b:c", ":")        // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
} // 通过案例

func TestMoreSplit1(t *testing.T) {
	got := Split1("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
} //  不通过案例

// 修改后的split，起名为split2
func TestMoreSplit2(t *testing.T) {
	got := Split2("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

func BenchmarkSplit2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split2("枯藤老树昏鸦", "老")
	}
}

// BenchmarkSplit-24    	 4105478	       336.3 ns/op	      48 B/op	       2 allocs/op

func BenchmarkSplit3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split3("枯藤老树昏鸦", "老")
	}
}

// BenchmarkSplit3-24    	 7299222	       255.4 ns/op	      32 B/op	       1 allocs/op
