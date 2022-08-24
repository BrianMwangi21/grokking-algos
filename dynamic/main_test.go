package main

import "testing"

type subTest struct {
	a, b, expected string
}

type seqTest struct {
  a, b string 
  expected int
}

func TestSubstring(t *testing.T) {
  tests := []subTest{
    {"vista", "hish", "is"},
    {"fish", "hish", "ish"},
    {"geeksforgeeks", "geeksquiz", "geeks"},
    {"abcdxyz", "xyzabcd", "abcd"},
    {"bandana", "march band", "band"},
  }

  for _, test := range tests {
    output := Substring(test.a, test.b)
    if output != test.expected {
      t.Errorf("Want %v, Got %v", test.expected, output)
    }
  }
}

func TestSequence(t *testing.T) {
  tests := []seqTest{
    {"fish", "fosh", 3},
    {"fort", "fosh", 2},
    {"abcde", "ace", 3},
    {"abc", "abc", 3},
    {"abc", "def", 0},
  }

  for _, test := range tests {
    output := Subsequence(test.a, test.b)
    if output != test.expected {
      t.Errorf("Want %v, Got %v", test.expected, output)
    }
  }
}
