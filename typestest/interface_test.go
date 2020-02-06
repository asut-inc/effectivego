package typestest

import (
	"fmt"
	"testing"
)

type A struct{}

type Fooer interface{}

func testInterface(a Fooer) string {
	// a Fooer это интерфейc, он состоит T(type) and (V) value
	// когда объявил через var x *A у тебя есть T, но нет V
	// nil ИСКЛЮЧИТЕЛЬНО тогда, КОГДА T и V НЕ УСТАНОВЛЕНЫ
	// https://golang.org/doc/faq#nil_error разбор тут (c) @helgix
	if a != nil {
		return fmt.Sprintf("not a nil")
	}

	return fmt.Sprintf("is nil")
}

func TestInterface(t *testing.T) {
	var x *A
	res := testInterface(x)
	fmt.Println(res)
	res = fmt.Sprintf("%v", x)
	fmt.Println(res)
}
