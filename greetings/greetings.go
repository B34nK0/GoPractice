package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//对外暴露function时，采用大写开头，即可被其他文件所访问
//输入参数name，类型为string， 该方法返回string
//返回多参数时，需要使用括号
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	//采用 := 时，go自动声明及赋值变量，等价于 var message string = ""
	//这里采用格式化将name嵌入到字符串中并赋值给message进行返回
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//接收数组参数，并返回map用于映射输入的name数组元素对应的信息
func Hellos(names []string) (map[string]string, error) {
	//创建map存储name与message的键值对
	messages := make(map[string]string)
	//遍历names数组
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

//go在初始化完全局变量后会执行init函数
//为rand以当前事件设置一个随机种子
func init() {
	rand.Seed(time.Now().UnixNano())
}

//函数为小写，相当于private，只能在当前文件调用
//使用slice存储多个消息格式，通过随机值获取不同的消息进行展示
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
