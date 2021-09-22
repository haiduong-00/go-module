package join_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-learn/TestnBench/join"
)

func TestCat(t *testing.T) {
	s := "Hello World of Go"
	xs := strings.Split(s, " ")
	if re := join.Cat(xs); re != s {
		t.Error("expected", s, ".", "got", re, ".")
	}
}

func TestJoin(t *testing.T) {
	s := "Hello World of Go"
	xs := strings.Split(s, " ")
	if re := join.Join(xs); re != s {
		t.Error("expected", s, ".", "got", re, ".")
	}
}

func ExampleCat() {
	s := "Hello World of Go"
	xs := strings.Split(s, " ")
	fmt.Println(join.Cat(xs))
	//Output:
	//Hello World of Go
}

func ExampleJoin() {
	s := "Hello World of Go"
	xs := strings.Split(s, " ")
	fmt.Println(join.Join(xs))
	//Output:
	//Hello World of Go
}

func BenchmarkJoin(b *testing.B) {
	s := "Hello World of Go"
	xs := strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if re := join.Join(xs); re != s {
			b.Error("expected", s, ".", "got", re, ".")
		}
	}
	b.ReportAllocs()
}

func BenchmarkCat(b *testing.B) {
	s := "Hello World of Go"
	xs := strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if re := join.Cat(xs); re != s {
			b.Error("expected", s, ".", "got", re, ".")
		}
	}
	b.ReportAllocs()
}
