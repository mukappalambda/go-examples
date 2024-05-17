package helloworld

import (
	"fmt"
	"testing"
	"time"
)

func TestBugSum(t *testing.T) {
	a := 1
	b := -1
	wrong := a + b
	got := BugSum(a, b)
	if got == wrong {
		t.Errorf("Got %d", got)
	}
}

func TestSum(t *testing.T) {
	out := Sum(1, -1)
	if out != 0 {
		t.Errorf("Got %d; want 0", out)
	}
}

var run bool

func TestSumSkip(t *testing.T) {
	if !run {
		t.Skip("This test is skipped.")
	}
}

func ExampleSum() {
	// This is a testable example in Go
	fmt.Println(Sum(1, 1), Sum(1, -2))
	fmt.Println(Sum(5, 5))
	// Output:
	// 2 -1
	// 10
}

func TestSumTableDriven(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{1, -1, 0},
		{1, 2, 3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d, %d", tt.a, tt.b)

		t.Run(testname, func(t *testing.T) {
			ans := Sum(tt.a, tt.b)

			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(5, 5)
	}
}

func BenchmarkSumAgain(b *testing.B) {
	// time.Sleep simulates a costly computation
	time.Sleep(time.Second)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(5, 5)
	}
}
