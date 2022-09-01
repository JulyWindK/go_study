package test

import "fmt"

// 接口定义
type Flyer interface { // 定义飞行动物接口
	Fly()
}

type Walker interface { // 定义行走动物接口
	Walk()
}

// 接口实现
type bird struct{} // 定义飞行类型

func (b *bird) Fly() { // 实现飞行动物飞行接口
	fmt.Println("bird: fly")
}

func (b *bird) Walk() { // 实现飞行动物行走接口
	fmt.Println("bird: walk")
}

// 接口实现
type pig struct{} // 定义行走类型

func (p pig) Walk() { // 实现行走动物行走接口
	fmt.Println("pig: walk")
}

// InterfaceTest 自定义接口类型的切片初始化及访问
func interfaceTest1() {

	stus := []people{NewStu("ls", 23), NewStu("ww", 43)}

	for i, s := range stus {
		s.say("hello s say ")
		stus[i].say("hello stus say")
	}
	stus[0].say("hello world")
	fmt.Printf("len: %d\n", len(stus))
	fmt.Printf("stus:%T\n", stus)
	fmt.Printf("stus:%+v\n", stus)
}

// 断言
func interfaceTest2() {
	// 创建动物名到结构体实例的字典
	animals := map[string]interface{}{
		"bird": new(bird), // 创建出的结构体实例
		"pig":  new(pig),
	}

	// 遍历字典
	for name, obj := range animals { // obj为字典的值，是interface{}类型
		f, isFlyer := obj.(Flyer)   // 使用类型断言获得变量f，转换后的类型是Flyer；isFlyer是接口类型转换是否成功的结果
		w, isWalker := obj.(Walker) // 使用类型断言获得变量w，转换后的类型是Walker；isWalker是接口类型转换是否成功的结果

		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)

		if isFlyer {
			f.Fly() // 调用接口方法
		}

		if isWalker {
			w.Walk() // 调用接口方法
		}
	}
}

// 把接口转换为其他类型
func interfaceTest3() {
	//p1 := new(pig)
	//
	//var w Walker = p1
	//p2 := w.(*pig)
	//
	//fmt.Printf("p1=%p p2=%p", p1, p2)

	p1 := pig{}

	var w Walker = p1

	fmt.Printf("p1 type:%T %p,  w type:%T %p\n", p1, &p1, w, &w)

}

func InterfaceTest() {
	// interfaceTest1()
	// interfaceTest2()
	interfaceTest3()
}
