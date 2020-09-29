// 数据类型
package main

import "fmt"

func main() {
    const (// 这种因式分解关键字的写法一般用于声明全局变量
        g1 string = "g1"
        g2 string = "g2"
    )
    var subName string= "subName"
    var d1,d2 int = 1,2
    var boolean bool = false;
    shengming := 33;// 这种不带声明格式的只能在函数体中出现

    

    mnInt1,mnInt2,mnString := multiNumber();

    var arrayInt  = [3] int{3,5,6}

    var point *int 
    point = &d2;

    fmt.Println(subName)
    fmt.Println(d1,d2)
    fmt.Println(boolean)
    fmt.Println(shengming)
    fmt.Println(g1,g2)
    fmt.Println(mnInt1,mnInt2,mnString)
    fmt.Println(arrayInt)
    fmt.Println("指针：",point,*point)
}


func multiNumber()(int,int,string){
    a , b , c := 1 , 2 , "str"
    return a,b,c
}