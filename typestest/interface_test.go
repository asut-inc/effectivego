package typestest

import (
	"fmt"
	"testing"
)

type A struct{}

type Fooer interface{}

func testInterface(a Fooer) Fooer {
	// a Fooer это интерфейc, он состоит T(type) and (V) value
	// когда объявил через var x *A у тебя есть T, но нет V
	// nil ИСКЛЮЧИТЕЛЬНО тогда, КОГДА T и V НЕ УСТАНОВЛЕНЫ
	// https://golang.org/doc/faq#nil_error разбор тут (c) @helgix
	if a != nil {
		return a
	}

	return a
}

func TestInterface(t *testing.T) {
	var x *A
	if x != nil {
		t.Errorf("Error not a nil pointer %T : %v", x, x)
	}
	res := testInterface(x)
	if res != nil {
		fmt.Printf("%T not a nil %v \n", res, res)
	}

	if res == nil {
		t.Errorf("Interface type is nil %T : %v", res, res)
	}
}
