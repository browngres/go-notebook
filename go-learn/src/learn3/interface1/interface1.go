package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{

}

func (stu *Student) Speak(think string) (talk string) {
	if think == "hello" {
		talk = "你好"
	} else {
		talk = "good bye"
	}
	return
}

func main() {
	// var peo People = Student{}
	// cannot use Student{} (value of type Student) as People value in variable declaration:
	// Student does not implement People (method Speak has pointer receiver)
	var peo People = &Student{}   // 这样可以
	think := "123"
	fmt.Println(peo.Speak(think))
}