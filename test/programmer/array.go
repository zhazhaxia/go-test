package main

import "fmt"

func main() {
   var n [10]int /* n 是一个长度为 10 的数组 */
   var i,j int

   /* 为数组 n 初始化元素 */        
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* 设置元素为 i + 100 */
   }

   /* 输出每个数组元素的值 */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }

   fmt.Println("长度：", len(n) ,cap(n))


   // 切片 声明一个未指定大小的数组来定义切片：
   var arr []int
   arr = append(arr,1)
   arr = append(arr,2)

   fmt.Println("arr", arr )

   var arr1 = []int{1}
   arr1 = append(arr1,1)

   fmt.Println("arr", arr1 )

   var arr2 = make([]int,1,2)
   arr2[0] = 1
   arr2 = append(arr2,2)
   arr2 = append(arr2,3,4,5)

   fmt.Println("arr", arr2 ,arr2[1:],arr2[:2],arr2[2:3])
   // 复制，必须切片是原来的翻倍
   var arr3  = make([]int,len(arr2),cap(arr2)*2);
   copy(arr3,arr2)

   fmt.Println("arr3", arr3 ,len(arr3),cap(arr3))
}