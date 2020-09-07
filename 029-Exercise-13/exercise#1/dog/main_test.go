package dog

import (
	"fmt"
	"testing"
)

func exampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func TestYears(t *testing.T) {
	n := Years(10)
	if n != 70 {
		t.Error("Expected", 70, "got", n)
	}
}

func TestYearsTwo(t *testing.T) {
	n := YearsTwo(10)
	if n != 70 {
		t.Error("Expected", 70, "got", n)
	}
}

func exampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
