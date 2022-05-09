package main

//可执行的独立程序都包含一个main包

//导入包
import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
