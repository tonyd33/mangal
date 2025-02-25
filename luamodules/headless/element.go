package headless

import (
	"github.com/go-rod/rod"
	"github.com/metafates/mangal/log"
	lua "github.com/yuin/gopher-lua"
)

var elementMethods = map[string]lua.LGFunction{
	"input":     elementInput,
	"click":     elementClick,
	"text":      elementText,
	"attribute": elementAttribute,
	"html":      elementHtml,
	"property":  elementProperty,
}

func registerElementType(L *lua.LState) {
	mt := L.NewTypeMetatable("pageElement")
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), elementMethods))
}

func checkElement(L *lua.LState) *rod.Element {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*rod.Element); ok {
		return v
	}
	log.Error("element expected")
	L.ArgError(1, "element expected")
	return nil
}

func elementInput(L *lua.LState) int {
	el := checkElement(L)
	value := L.ToString(2)
	log.Info("inputting " + value)
	el.MustInput(value)

	return 0
}

func elementClick(L *lua.LState) int {
	el := checkElement(L)
	log.Info("clicking")
	el.MustClick()

	return 0
}

func elementText(L *lua.LState) int {
	el := checkElement(L)
	log.Info("getting text")
	text := el.MustText()

	L.Push(lua.LString(text))
	return 1
}

func elementAttribute(L *lua.LState) int {
	el := checkElement(L)
	name := L.ToString(2)
	log.Info("getting attribute " + name)
	value := el.MustAttribute(name)

	L.Push(lua.LString(*value))
	return 1
}

func elementHtml(L *lua.LState) int {
	el := checkElement(L)
	log.Info("getting html")
	html := el.MustHTML()

	L.Push(lua.LString(html))
	return 1
}

func elementProperty(L *lua.LState) int {
	el := checkElement(L)
	name := L.ToString(2)
	log.Info("getting property " + name)
	value := el.MustProperty(name)

	L.Push(lua.LString(value.Str()))
	return 1
}
