package middleware

import (
	"fmt"
	"github.com/test1280/demo-go-middleware/test1280"
	"strings"
)

// index 标记Wrap生产出的中间件的序号
var index = 0

// Wrap 中间件工厂函数，创建中间件并返回，不可重入（index无锁保护）
func Wrap(name string) test1280.MiddlewareFunc {

	// 自增index标记
	index = index + 1

	// 根据name生成携带有index的新name
	name = fmt.Sprintf("wrap[%02d] %s", index, name)

	// 前导空格数量，注意，不能放置到MiddlewareFunc或者HandlerFunc 中
	spaceN := index

	// 创建新的MiddlewareFunc函数，注意，每次调用Wrap生成的是不同的MiddlewareFunc
	return func(next test1280.HandlerFunc) test1280.HandlerFunc {
		// 创建新的HandlerFunc函数，其是对原有HandlerFunc函数(next)的包装
		return func(c *test1280.Context) (err error) {
			// 调用下个中间件  前输出
			fmt.Println(strings.Repeat(" ", spaceN) + name + ">")
			// 调用下个中间件
			err = next(c)
			// 调用下个中间件  后输出
			fmt.Println(strings.Repeat(" ", spaceN) + name + "<")
			return
		}
	}
}
