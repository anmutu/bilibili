/*
  author='du'
  date='2021/1/11 23:29'
*/
package main

func main() {
	//var peo People = Stduent{}
	//think := "love"
	//fmt.Println(peo.Speak(think))
}

//People为interface类型，就是指针类型.
type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "love" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

//继承与多态的特点
//在golang中对多态的特点体现从语法上并不是很明显。
//我们知道发生多态的几个要素：
//1、有interface接口，并且有接口定义的方法。
//2、有子类去重写interface的接口。
//3、有父类指针指向子类的具体对象
//那么，满足上述3个条件，就可以产生多态效果，就是，父类指针可以调用子类的具体方法。
//所以上述代码报错的地方在var peo People = Stduent{}这条语句，
//Student{}已经重写了父类People{}中的Speak(string) string方法，那么只需要用父类指针指向子类对象即可。
//所以应该改成var peo People = &Student{} 即可编译通过。（People为interface类型，就是指针类型）
//func main() {
//	var peo People = Stduent{}
//	think := "love"
//	fmt.Println(peo.Speak(think))
//}
//
//type People interface {
//	Speak(string) string
//}
//
//type Stduent struct{}
//
//func (stu *Stduent) Speak(think string) (talk string) {
//	if think == "love" {
//		talk = "You are a good boy"
//	} else {
//		talk = "hi"
//	}
//	return
//}
