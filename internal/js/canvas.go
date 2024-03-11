package js

import "syscall/js"

type Canvas struct {
	JSCanvas Value
}

type Context struct {
	JSContext Value
}

func (c *canvas) GetContext(context_str string) context {
	return context{c.canvas.Call("getContext", context_str)}
}

func (c *canvas) Get(key string) js.Value {
	return c.canvas.Get(key)
}

func (c *canvas) Set(key string, value interface{}) {
	c.canvas.Set(key, value)
}

func (c *canvas) IsNull() bool {
	return c.canvas.IsNull()
}

func (c *canvas) AddEventListener(ev string, fn js.Func) {
	js.Global().Get("window").Call("addEventListener", ev, fn)
}

func (c *context) Set(key string, value interface{}) {
	c.context.Set(key, value)
}

func (c *context) FillText(text string, x, y float64) {
	c.context.Call("fillText", text, x, y)
}

func (c *context) FillRect(x, y, width, height float64) {
	c.context.Call("fillRect", x, y, width, height)
}

func (c *context) ClearRect(x, y, width, height float64) {
	c.context.Call("clearRect", x, y, width, height)
}

func (c *context) DrawImage(img js.Value, x, y, width, height float64) {
	c.context.Call("drawImage", img, x, y, width, height)
}
