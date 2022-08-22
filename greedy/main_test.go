package main

import (
	"reflect"
	"testing"
)

type greedyStruct struct {
  attributes_needed []string
  expected []string
}

func TestGreedy(t *testing.T) {
	var candidates = make(map[string][]string)

	candidate_key := []string{"kimani", "kabiru", "kumbaya", "gatutura", "kimutai"}
	candidates["kimani"] = []string{"funny", "problem-solver", "passionate"}
	candidates["kabiru"] = []string{"critical", "funny", "sharp"}
	candidates["kumbaya"] = []string{"educated", "problem-solver", "caring"}
	candidates["gatutura"] = []string{"problem-solver", "passionate", "leader"}
	candidates["kimutai"] = []string{"caring", "good attitude"}


	tests := []greedyStruct{
	  {[]string{"funny", "problem-solver", "passionate"}, []string{"kimani"}},
	  {[]string{"caring", "good attitude", "sharp"}, []string{"kimutai", "kabiru"}},
	  {[]string{"passionate", "sharp", "educated"}, []string{"kimani", "kabiru", "kumbaya"}},
	  {[]string{"sharp", "critical", "educated", "funny", "problem-solver", "passionate", "caring", "good attitude"}, []string{"kimani", "kabiru", "kumbaya", "kimutai"}},
	  {[]string{"problem-solver", "passionate", "leader"}, []string{"gatutura"}},
	}

	for _, test := range tests {
	  output := Greedy(test.attributes_needed, candidate_key, candidates)
	  if !reflect.DeepEqual(test.expected, output) {
	    t.Errorf("Want %v, Got %v", test.expected, output)
	  } 
	}
}
