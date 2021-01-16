/*
  author='du'
  date='2021/1/11 23:39'
*/
package main

import "fmt"

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

//非空接口
type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

//interface的内部构造(非空接口iface情况)
//我们需要了解interface的内部结构，才能理解这个题目的含义。
//interface在使用的过程中，共有两种表现形式
//一种为空接口(empty interface)，定义如下：
//var MyInterface interface{}
//另一种为非空接口(non-empty interface), 定义如下：
//type MyInterface interface {
//	function()
//}
//这两种interface类型分别用两种struct表示，空接口为eface, 非空接口为iface.

//空接口eface
//空接口eface结构，由两个属性构成，一个是类型信息_type，一个是数据信息。其数据结构声明如下：
//type eface struct {      //空接口
//	_type *_type         //类型信息
//	data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
//}
//_type属性：是GO语言中所有类型的公共描述，Go语言几乎所有的数据结构都可以抽象成 _type，
//是所有类型的公共描述，**type负责决定data应该如何解释和操作，**type的结构代码如下:
//type _type struct {
//	size       uintptr  //类型大小
//	ptrdata    uintptr  //前缀持有所有指针的内存大小
//	hash       uint32   //数据hash值
//	tflag      tflag
//	align      uint8    //对齐
//	fieldalign uint8    //嵌入结构体时的对齐
//	kind       uint8    //kind 有些枚举值kind等于0是无效的
//	alg        *typeAlg //函数指针数组，类型实现的所有方法
//	gcdata    *byte
//	str       nameOff
//	ptrToThis typeOff
//}
//data属性: 表示指向具体的实例数据的指针，他是一个unsafe.Pointer类型，相当于一个C的万能指针void*。

//非空接口iface
//iface 表示 non-empty interface 的数据结构，非空接口初始化的过程就是初始化一个iface类型的结构，
//其中data的作用同eface的相同，这里不再多加描述。
//
//type iface struct {
//	tab  *itab
//	data unsafe.Pointer
//}
//iface结构中最重要的是itab结构（结构如下），每一个 itab 都占 32 字节的空间。
//itab可以理解为pair<interface type, concrete type> 。itab里面包含了interface的一些关键信息，比如method的具体实现。
//
//type itab struct {
//	inter  *interfacetype   // 接口自身的元信息
//	_type  *_type           // 具体类型的元信息
//	link   *itab
//	bad    int32
//	hash   int32            // _type里也有一个同样的hash，此处多放一个是为了方便运行接口断言
//	fun    [1]uintptr       // 函数指针，指向具体类型所实现的方法
//}
//其中值得注意的字段，个人理解如下：
//interface type包含了一些关于interface本身的信息，比如package path，包含的method。
//这里的interfacetype是定义interface的一种抽象表示。
//type表示具体化的类型，与eface的 type类型相同。
//hash字段其实是对_type.hash的拷贝，它会在interface的实例化时，用于快速判断目标类型和接口中的类型是否一致。
//另，Go的interface的Duck-typing机制也是依赖这个字段来实现。
//fun字段其实是一个动态大小的数组，虽然声明时是固定大小为1，但在使用时会直接通过fun指针获取其中的数据，
//并且不会检查数组的边界，所以该数组中保存的元素数量是不确定的。

//stu是一个指向nil的空指针，但是最后return stu 会触发匿名变量 People = stu值拷贝动作，
//所以最后live()放回给上层的是一个People insterface{}类型，也就是一个iface struct{}类型。
//stu为nil，只是iface中的data 为nil而已。 但是iface struct{}本身并不为nil.

//func main() {
//	if live() == nil {
//		fmt.Println("AAAAAAA")
//	} else {
//		fmt.Println("BBBBBBB")
//	}
//}
//
////非空接口
//type People interface {
//	Show()
//}
//
//type Student struct{}
//
//func (stu *Student) Show() {
//
//}
//
//func live() People {
//	var stu *Student
//	return stu
//}
//
