package middleware

import (
	"fmt"
	"github.com/test1280/demo-go-middleware/test1280"
)

func Demo(next test1280.HandlerFunc) test1280.HandlerFunc {
	return func(c *test1280.Context) (err error) {
		fmt.Println("demo>")
		err = next(c)
		fmt.Println("demo<")
		return
	}
}
