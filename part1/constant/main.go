package main

import "fmt"

func main() {

	const length int = 10
	const width = 20
	fmt.Println(length)
	fmt.Println(width)
	//多重赋值
	const A, B, C = 1, false, "string"
	fmt.Println(A, B, C)

	//枚举
	const (
		UNKNOWN = 0
		FEMALE  = 1
		MALE    = 2
	)

	fmt.Println(UNKNOWN, FEMALE, MALE)

	//iota,一个特殊的常量.iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
	// const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

}
