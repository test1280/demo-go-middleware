package test1280

type (
	// HandlerFunc 处理函数
	HandlerFunc func(*Context) error
	// MiddlewareFunc 修饰函数，修饰HandlerFunc；包裹函数，包裹HandlerFunc
	MiddlewareFunc func(HandlerFunc) HandlerFunc
)

// Test1280 框架对象
type Test1280 struct {
	// middleware 中间件，由多个MiddlewareFunc构成，或称MiddlewareFuncs更合适
	middleware []MiddlewareFunc
	// handlerFunc 处理函数，例如HTTP请求回调函数
	handlerFunc HandlerFunc
}

// New 创建框架对象
func New() (t *Test1280) {
	t = &Test1280{
		middleware:  make([]MiddlewareFunc, 0),
		handlerFunc: nil,
	}
	return
}

// SetHandlerFunc 设置框架对象处理函数
func (t *Test1280) SetHandlerFunc(h HandlerFunc) {
	t.handlerFunc = h
}

// Use 添加中间件
func (t *Test1280) Use(middleware ...MiddlewareFunc) {
	t.middleware = append(t.middleware, middleware...)
}

func (t *Test1280) Do(c *Context) {

	// 请求回调处理函数；通常，此处应当根据请求的特征，如Method、URL等路由到某个HandlerFunc
	core := t.handlerFunc

	// 最终回调处理函数（集）：中间件(middleware)与请求回调处理函数(handlerFunc)
	h := core
	for i := len(t.middleware) - 1; i >= 0; i-- {
		h = t.middleware[i](h)
	}

	// 执行
	h(c)
}

type Context struct {
	request  interface{}
	response interface{}
	other    interface{}
}

func (c *Context) Request() interface{} {
	return c.request
}

func (c *Context) SetRequest(request interface{}) {
	c.request = request
}

func (c *Context) Response() interface{} {
	return c.response
}

func (c *Context) SetResponse(response interface{}) {
	c.response = response
}

func (c *Context) Other() interface{} {
	return c.other
}
