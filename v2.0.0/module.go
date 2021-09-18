// Thử nghiệm publish một module lên pkg.go.dev
//Chào mừng phiên bản v2.0.0 với tính năng check lỗi, đảm bảo code WriteHello được chạy trơn tru mà vẫn có thể dễ dàng gỡ lỗi
//
//Các tính năng mới
/*
	Thêm check
*/
//Sửa đổi
/* 
	Check kỹ WriteHello
*/
package gomodule

import "fmt"


// WriteHello được dùng để in chữ Hello ra Console
/*
To go back to where you were, just check out the branch you were on again. (If you've made changes, as always when switching branches, you'll have to deal with them as appropriate. You could reset to throw them away; you could stash, checkout, stash pop to take them with you; you could commit them to a branch there if you want a branch there.)
*/
func WriteHello() {
	_, err :=fmt.Println("Hello")
	Check(err)
	_, err = fmt.Println("Tính năng mới: Viết doc nhiều dòng")
	Check(err)
}

// Check được sử dụng để check lỗi
func Check(err error)  {
	if err != nil {
		fmt.Println(err)
	}
}
