package main

import "fmt"

func main() {
   sum := 0
   for i := 0; i <= 10; i++ {
            sum += i
   }
   fmt.Println(sum)
   // 类似for each
   strings := []string{"google", "runoob"}
   for i, s := range strings {
      fmt.Println(i, s)
   }


}