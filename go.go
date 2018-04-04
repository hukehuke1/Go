// 第一个简单的go程序
// 当前程序的包名
package main
// 导入其他包
import "fmt"
// 由main函数作为程序入口点启动
func main1(){
	// 一般类型声明
	var age int
	age = 10
	// 常量定义
	const PI = 3.14
	fmt.Println("Hello,World!")
	fmt.Println(age)


	var a = 1.5
	var b = 2.5
	fmt.Println(a,b)
}

func main2(){
	var isActive bool  // 全局变量声明
	var enabled, disabled = true, false  //忽略类型的声明
	var available bool //一般声明
	valid  := false  //简短声明
	available = true //赋值操作
	a, b, c := 5, 7, "abc"
	fmt.Print(isActive,enabled,disabled,available,valid)
	fmt.Print(a,b,c)
}

func main(){
	const identifier string = "abc"
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c =1, false ,"str"//多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为:%d",area)
	println()
	println(a,b,c)
}