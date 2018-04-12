// 第一个简单的go程序
// 当前程序的包名
package main
// 导入其他包
import (
	"strings"
	"flag"
	"bufio"
	"os"
	"fmt"
)
// import "unsafe"
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

func main3(){
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


func main4(){
	/* const(
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	println(a,b,c) */
    const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}

func main5(){
	const(
		i = 1 <<iota
		j = 3 <<iota
		k
		l
	)
    fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
}

func main6(){
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("1 - C = %d\n",c)
	c = a - b
	fmt.Printf("2 - C = %d\n",c)
	c = a * b
	fmt.Printf("3 - C = %d\n",c)
	c = a / b
	fmt.Printf("4 - C = %d\n",c)
	c = a % b
	fmt.Printf("5 - C = %d\n",c)
	a++
	fmt.Printf("6 - a = %d\n",a)
	a = 21
	a--
	fmt.Printf("7 - a = %d\n",a)

	if( a == b ) {
		fmt.Printf("第一行 - a 等于 b\n" )
	 } else {
		fmt.Printf("第一行 - a 不等于 b\n" )
	 }
	 if ( a < b ) {
		fmt.Printf("第二行 - a 小于 b\n" )
	 } else {
		fmt.Printf("第二行 - a 不小于 b\n" )
	 } 
	 
	 if ( a > b ) {
		fmt.Printf("第三行 - a 大于 b\n" )
	 } else {
		fmt.Printf("第三行 - a 不大于 b\n" )
	 }
	 /* Lets change value of a and b */
	 a = 5
	 b = 20
	 if ( a <= b ) {
		fmt.Printf("第四行 - a 小于等于 b\n" )
	 }
	 if ( b >= a ) {
		fmt.Printf("第五行 - b 大于等于 a\n" )
	 }
}

func main7(){
	//求100以内的素数
	var C ,c int//声明变量
	C = 1 
	A: for C <100{
		C++ //C=1不能写入for这里就不能写入
		for c = 2; c<C ; c ++ {
			if C%c == 0{
				goto A //若发现因子则不是素数
			}
		}
		fmt.Println(C,"是素数")
	}
}


func main8(){
	for i := 1; i <= 9; i++{ //i控制行
		for j :=1; j <= i ; j++{
			fmt.Printf("%d*%d=%d ",j,i,j*i)
		}
		fmt.Println(" ")
	}
}


func max(num1,num2 int) int{
	/* 声明局部变量 */
	var result int

	if (num1 > num2){
		result = num1
	}else {
		result = num2
	}
	return result
}

func main9(){
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 80

	switch marks{
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70: grade = "C"
		default: grade = "D"
	}

	switch {
		case grade =="A":
			fmt.Printf("优秀！\n")
		case grade =="B":
			fmt.Printf("良好!\n")
		case grade =="D":
			fmt.Printf("及格\n")
		case grade =="F":
			fmt.Printf("不及格\n")
		default:
			fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是%s \n",grade)
}


/* 定义相互交换值的函数 */
func swap(x,y int) int {
	var temp int
	temp = x /* 保存X的值 */
	x = y 
	y = temp

	return temp;
}


func main10(){
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a )
	fmt.Printf("交换前 b 的值为 : %d\n", b )

   /* 调用 swap() 函数
   * &a 指向 a 指针，a 变量的地址
   * &b 指向 b 指针，b 变量的地址
   */
   swap2(&a, &b)

	fmt.Printf("交换后 a 的值为 : %d\n", a )
	fmt.Printf("交换后 b 的值为 : %d\n", b )
}


func swap2(x *int,y *int){
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}



//形式参数
var abc int = 20;
func main11(){
	var abc ,b, c int

	/* 初始化参数 */
	abc = 10
	b = 20
	c = 0

	fmt.Printf("main()函数中 a = %d\n",  abc);
	c = sum(abc,b)
	fmt.Printf("main()函数中 c = %d\n",  c);
}


/* 函数定义-两数相加 */
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n",  a);
	fmt.Printf("sum() 函数中 b = %d\n",  b);
 
	return a + b;
 }


 func main12(){
	//  var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	 var n [10] int /* n 是一个长度为 10 的数组 */
	 var i ,j int
	 /* 为数组 n 初始化元素 */   
	 for i = 0; i < 10; i ++{
		 n[i] = i + 100 /* 设置元素为 i + 100 */
	 }

	 /* 输出每个数组元素的值 */
	 for j = 0; j <10; j++{
		 fmt.Printf("Element[%d]=%d\n",j,n[j])
	 }
 }


 func main13(){
	 /* 数组长度为 5 */
	 var balance = []int{1000,2,3,17,50}
	 var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage( balance, 5 ) ;

	/* 输出返回的平均值 */
	fmt.Printf("平均值为:%f",avg)
 }


 func getAverage(arr []int, size int) float32{
	 var i ,sum int
	 var avg float32

	 for i = 0 ; i <size; i ++{
		 sum += arr[i]
	 }

	 avg = float32(sum)/float32(size)

	 return avg;
 }


func main14(){
	a := 1690
	b := 1700
	c := a*b
	fmt.Println(c)
	fmt.Println(float64(c)/ 1000000)
}



func main15(){
	var a int = 20
	var ip *int   /* 声明指针变量 */
	ip = &a  /* 指针变量的存储地址 */
	fmt.Printf("a 变量的地址是:%x\n",&a)
   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )
}

func main16(){
	var s ,sep string
	for i := 1;i <len(os.Args); i ++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(os.Args[0])
}

func main17(){
	fmt.Printf("第0个参数是%s",os.Args[0])
	for i :=1; i <len(os.Args);i ++{
		fmt.Printf("第%d个参数是%s",i,os.Args[i])
	}
}

func main18(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counts[input.Text()]++
	}
	for line, n:= range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}


// go圣经 ftoc
func main19(){
	const freezingF, boilingF= 32.0,212.0
	fmt.Printf("%g°F= %g°C\n",freezingF,fToC(freezingF))
	fmt.Printf("%g°F= %g°C\n",boilingF,fToC(boilingF))
}

func fToC(f float64) float64{
	return (f-32)*5/9
}

func main20(){
	x := 1
	p := &x  
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
}

func main21(){
	var x,y int
	fmt.Println(&x == &x,&x == &y,&x == nil)
	//true false false
}

func incr(p *int) int{
	*p++//只增加p指向变量的值，并不改变p指针
	return *p
}

func main22(){
	v := 1
	incr(&v)   //v is now 2
	fmt.Println(incr(&v))   //v is now 3
}

func main(){
	echo4()
}


var n = flag.Bool("n",false,"omit trailing newline")
var sep = flag.String("s"," ","separator")
func echo4(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n{
		fmt.Println()
	}
}


var global *int
func f(){
	var x int
	x = 1
	global = &x
}

func g(){
	y := new(int)
	*y = 1
}