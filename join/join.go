// Package join thực hiện các chức năng gắn 1 []string vào 1 string to hơn
package join

import "strings"

// Cat viết tắt cho concatenation (nối): nối các strings nhỏ lại với nhau
func Cat(xs []string) string {

	// Tạo ra Cat làm kết quả, gắn chữ đầu tiên trong xs vào Cat
	Cat := xs[0]

	// Vòng lặp thêm giá trị vào Cat
	for i, v := range xs {
		if i == 0 {
			continue
		}
		Cat += " "
		Cat += v
	}

	return Cat
}

// Join thực hiện hành động Join trong package strings, thực hiện chức năng nối các strings nhỏ lại với nhau, và có tốc độ nhanh hơn Cat
func Join(xs []string) string {
	return strings.Join(xs," ")
}
