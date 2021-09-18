package gomodule

import "fmt"


// WriteHello được dùng để in chữ Hello ra Console
func WriteHello() {
	fmt.Println("Hello")
}

// foo là private (chỉ có trong package, không xuất ra ngoài)
func foo() {
	fmt.Println("foo")
}