package main

import (
	"fmt"
	"testing"
)

func Sum(a int, b int) int {
  return a + b
}

func TestSum(t *testing.T) {
  out := Sum(1, -1)
  if out != 0 {
    t.Errorf("Got %d; want 0", out)
  }
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

    t.Run(testname, func (t *testing.T) {
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