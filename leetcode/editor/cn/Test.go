package main

import "fmt"

type inner struct {
	val int
}

type test struct {
	inn *inner
	val int
}

func (t test) changeVal(newVal int) {
	fmt.Printf("【方法内对象】改变前	内存地址：%p，对应的值为：%v，【内部对象】内存地址：%p，对应的值为：%v\n", &t, t.getVal(), t.inn, t.inn.val)
	t.val = newVal
	t.inn.val = newVal
	fmt.Printf("【方法内对象】改变后	内存地址：%p，对应的值为：%v，【内部对象】内存地址：%p，对应的值为：%v\n", &t, t.getVal(), t.inn, t.inn.val)
}

func (t test) getVal() int {
	return t.val
}

//需要在这里赋值str，但是又不能直接引用 str
//func setVal(val []string) {
//	fmt.Printf("【方法内对象】改变前	内存地址：%p，对应的值为：%v\n", &val, val)
//	val = []string{"a", "b"}
//	fmt.Printf("【方法内对象】改变后	内存地址：%p，对应的值为：%v\n", &val, val)
//}

func main() {
	t := &test{&inner{1}, 1}
	fmt.Printf("【原生对象】改变前	内存地址：%p，对应的值为：%v，【内部对象】内存地址：%p，对应的值为：%v\n", t, t.getVal(), t.inn, t.inn.val)
	t.changeVal(100)
	fmt.Printf("【原生对象】改变后	内存地址：%p，对应的值为：%v，【内部对象】内存地址：%p，对应的值为：%v\n", t, t.getVal(), t.inn, t.inn.val)

	//var arr []string
	//arr = []string{"1", "2"}
	//fmt.Printf("【原生对象】改变前	内存地址：%p，对应的值为：%v\n", &arr, arr)
	//setVal(arr)
	//fmt.Printf("【原生对象】改变后	内存地址：%p，对应的值为：%v\n", &arr, arr)
}
