package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	str := "mondays are red, sundays are blues"
	n := Count(str)
	if n != 6 {
		t.Error("Expected", 6, "got", n)
	}
}

func TestUseCount(t *testing.T) {
	str := "abcd abcd efgh abcd"
	m := UseCount(str)

	if m["abcd"] != 3 || m["efgh"] != 1 {
		t.Error("Expected: _", "got", m)
	}

	// switch m {
	// case "abcd":
	// 	t.Error("Expected", 3, "got", m["abcd"])
	// case "efgh":
	// 	t.Error("Expected", 1, "got", m["efgh"])
	// }

}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// k := "ab bcd cfhiuajdn de"
		Count("ab bcd cfhiuajdn de")
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("abcd abcd efgh abcd")
	}
}

func ExampleCount() {
	fmt.Println(Count("a b c d"))
	// Output:
	// 4
}
