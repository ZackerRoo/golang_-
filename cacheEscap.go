package main

import "unsafe"

type A struct {
	S *string
}

func (a *A) String() string {
	return *a.S
}

type ATrick struct {
	S unsafe.Pointer
}

func (a *ATrick) String() string {
	return *(*string)(a.S)
}

func NewA(s string) *A {
	return &A{S: &s}
}

func NewATrick(s string) *ATrick {
	return &ATrick{S: noescape(unsafe.Pointer(&s))}
}

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func main() {
	s := "hello"
	f1 := NewA(s)
	f2 := NewATrick(s)
	s1 := f1.String()
	s2 := f2.String()
	_ = s1 + s2
}
