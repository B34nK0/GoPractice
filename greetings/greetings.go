package greetings

import "fmt"

//对外暴露function时，采用大写开头，即可被其他文件所访问
//输入参数name，类型为string， 该方法返回string
func Hello(name string) string {
	//采用 := 时，go自动声明及赋值变量，等价于 var message string = ""
	//这里采用格式化将name嵌入到字符串中并赋值给message进行返回
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}