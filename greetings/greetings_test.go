package greetings

import (
	"regexp"
	"testing"
)

//指针指向testing包的testing.T类型，用于记录和报告错误信息
func TestHelloName(t *testing.T) {
	name := "Bean"
	//正则表达式
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello("Bean")
	//正则匹配检测调用Hello后输出的消息是否成功，或者返回错误信息时
	if !want.MatchString(message) || err != nil {
		//反馈给测试系统
		t.Fatalf(`Hello("Bean") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestEmptyName(t *testing.T) {
	message, err := Hello("")
	if err == nil || message != "" {
		t.Fatalf(`Hello("") = %q, %v, want ""`, message, err)
	}
}
