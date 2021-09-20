// Thử nghiệm publish một module lên pkg.go.dev
//
// Cái này sẽ xuất hiện trong v0.3.0
//
// Ok
//
// Còn Palind sẽ xuất hiện ở v2.1.0
// 
// Ok, fine
package gomodule

// Check xem số đó có phải palindrome không
func IsPalindrome(n int) string {
	x := []int{}
	for n >= 10 {
		x = append(x, n%10)
		n = n / 10
	}
	x = append(x, n)
	for i := len(x) / 2; i >= 0; i-- {
		if x[i] == x[len(x)-1-i] {
		} else {
			return "No"
		}
	}
	return "Yes"
}
