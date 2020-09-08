package mymath

import (
	"fmt"
	"testing"
)

func TestCenterAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{10, 20, 30, 40, 50}, 30},
		test{[]int{2, 4, 6, 8, 10, 12}, 7},
		test{[]int{1, 2, 3, 4, 5, 6, 7}, 4},
	}

	for _, v := range tests {
		f := CenteredAvg(v.data)

		if f != v.answer {
			t.Error("got", f, "want", v.answer)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 4, 6, 9}))
	// Output:
	// 4
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100000})
	}
}
