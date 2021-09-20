// Thử nghiệm publish một module lên pkg.go.dev
//
// Cái này sẽ xuất hiện trong v0.3.0
//
// Ok
package gomodule

import "fmt"


// WriteHello được dùng để in chữ Hello ra Console
/*
To go back to where you were, just check out the branch you were on again. (If you've made changes, as always when switching branches, you'll have to deal with them as appropriate. You could reset to throw them away; you could stash, checkout, stash pop to take them with you; you could commit them to a branch there if you want a branch there.)
*/
func WriteHello() {
	foo()
	fmt.Println("Hello")
	fmt.Println("Tính năng mới: Viết doc nhiều dòng")
}

// foo là private (chỉ có trong package, không xuất ra ngoài (exported))
func foo() {
	fmt.Println("foo")
}