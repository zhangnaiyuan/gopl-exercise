// 修改echo程序,使其能够打印os.Args[0],即命令本身
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}
