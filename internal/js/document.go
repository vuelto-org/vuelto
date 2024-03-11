package js

import "syscall/js"

type Document struct {
	Body            *Body
	DocumentElement *DocumentElement
}

type Body struct {
	Style StyleBody
}

type Element struct {
	JSElement js.Value
	Style     js.Value
}

type DocumentElement struct {
	Style StyleDocument
}

type StyleDocument struct{}
type StyleBody struct{}

func CreateElement(newElement string) *Element {
	docElement := js.Global().Get("document").Call("createElement", newElement)
	style := docElement.Get("style")

	return &Element{
		JSElement: docElement,
		Style:     style,
	}
}

func CreateCanvasElement() *Canvas {
	newElement := js.Global().Get("document").Call("createElement", "canvas")

	return &Canvas{
		JSCanvas: newElement,
	}
}

func AddEventListener(inputEvent string, inputFunc func(Value, []Value)) {
	js.Global().Get("document").Call("addEventListener", inputEvent, js.FuncOf(inputFunc))
}

func (e *Element) Set(keyValue string, inputValue any) {
	e.JSElement.Set(keyValue, inputValue)
}

func (e *Element) AddEventListener(inputEvent string, inputFunc func(Value, []Value)) {
	e.JSElement.Call("addEventListener", inputEvent, js.FuncOf(inputFunc))
}

func (b *Body) AppendChild(inputElement *Element) {
	appendChild := inputElement.JSElement
	js.Global().Get("document").Get("body").Call("appendChild", appendChild)
}

func (b *Body) AppendCanvasChild(inputCanvas *Canvas) {
	appendChild := inputCanvas.JSCanvas
	js.Global().Get("document").Get("body").Call("appendChild", appendChild)
}

func GetElementById(inputID string) *Canvas {
	return &Canvas{
		JSCanvas: js.Global().Get("document").Call("getElementById", inputID),
	}
}

func (d *DocumentElement) ClientWidth() int {
	return js.Global().Get("document").Get("documentElement").Get("clientWidth").Int()
}

func (d *DocumentElement) ClientHeight() int {
	return js.Global().Get("document").Get("documentElement").Get("clientHeight").Int()
}

func (d *StyleDocument) Set(inputKey string, inputValue any) {
	js.Global().Get("document").Get("documentElement").Get("style").Set(inputKey, inputValue)
}

func (b *StyleBody) Set(inputKey string, inputValue any) {
	js.Global().Get("document").Get("body").Get("style").Set(inputKey, inputValue)
}
