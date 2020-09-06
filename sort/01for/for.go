/*
  author='du'
  date='2020/8/29 20:10'
*/
package main

func main() {
	parseStudent()
}

type student struct {
	Name string
	Age  int
}

func parseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//错误的写法
	//for _, stu := range stus {
	//	println(&stu)
	//	m[stu.Name] = &stu
	//}

	//正确的写法
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}

	for k, v := range m {
		println(k, "=>", v.Name)
	}
}
