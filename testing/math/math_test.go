package math

import (
	"fmt"
	"testing"
)

func TestMySum(t *testing.T) {
	got := MySum(3, 4, 6)
	if got != 13 {
		t.Errorf("Expected 13, got %v", got)
	}
}

func ExampleMySum() {
	fmt.Println(MySum(5, 6, 7))
	//Output:
	//18
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MySum(12, 13)
	}
}
