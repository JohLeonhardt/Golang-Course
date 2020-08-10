package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v gets divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}
}

//Test1
//package main

//import (
//"fmt"
//)

//func main() {
//for i := 10; i <= 100; i++ {
//if i%4 != 0 {
//continue
//}
//fmt.Println(i)
//}
//}

//Test2
//package main

//import (
//"fmt"
//)

//func main() {
//x := 10
//for {
//x++
//if x > 100 {
//break
//}

//if x%4 != 0 {
//continue
//}

//fmt.Println(x)
//}
//}
