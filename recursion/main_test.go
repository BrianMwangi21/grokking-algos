package main 

import "testing"

type recursionTest struct {
  number int
  expected int 
}

func TestFactorial(t *testing.T) {
  tests := []recursionTest{
    {1, 1},
    {3, 6},
    {5, 120},
    {7, 5040},
    {10, 3628800},
  }

  for _, test := range tests {
    output := Factorial(test.number)
    if output != test.expected {
      t.Errorf("Want %q, Got %q", test.expected, output)
    }
  }
}
