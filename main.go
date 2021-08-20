package main

import (
	"fmt"
	"github.com/test1280/demo-go-middleware/test1280"
	"github.com/test1280/demo-go-middleware/test1280/middleware"
	"strings"
)

func main() {
	// 1.创建框架对象
	t := test1280.New()

	// 2.设置处理函数
	t.SetHandlerFunc(Hello)

	// 3.设置中间件集
	t.Use(middleware.Wrap("a"))
	t.Use(middleware.Wrap("b"))
	t.Use(middleware.Wrap("c"))
	t.Use(middleware.Wrap("d"))
	t.Use(middleware.Wrap("e"))
	t.Use(middleware.Demo)

	// 4.构建请求对象 or test1280.NewContext
	c := &test1280.Context{}
	c.SetRequest("test,1280")

	// 5.传入请求对象；通常是test1280做服务（如HTTP服务）监听端口收请求，此处简化直接通过函数调用模拟接收请求
	t.Do(c)

	// 6.获取请求结果
	res := c.Response()
	fmt.Println(res)
}

// Hello 请求回调处理函数
func Hello(c *test1280.Context) error {
	fmt.Println("Hello>")
	defer fmt.Println("Hello<")

	// 获取请求
	req := c.Request()
	input, ok := req.(string)
	if !ok {
		return fmt.Errorf("invalid req: %v", c.Request())
	}
	ss := strings.Split(input, ",")
	result := ""
	for _, s := range ss {
		result = result + s
	}
	// 设置应答
	c.SetResponse(result)
	return nil
}
