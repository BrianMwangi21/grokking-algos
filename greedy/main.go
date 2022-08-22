package main

import "fmt"

func equalData(attributes_needed []string, attributes []string) []string {
  var covered []string 
  for _, attribute_needed := range attributes_needed {
    for _, attribute := range attributes {
      if attribute_needed == attribute {
        covered = append(covered, attribute_needed)
      }
    }
  }
  return covered
}

func removeData(attributes_needed []string, attributes_covered []string) []string {
  for _, attribute_covered := range attributes_covered {
    attributes_needed = remove(attributes_needed, attribute_covered)
  } 
  return attributes_needed
}

func remove(attributes_needed []string, attribute_covered string) []string {
  for i, attribute_needed := range attributes_needed {
    if attribute_covered == attribute_needed {
      return append(attributes_needed[:i], attributes_needed[i+1:]...)
    }
  }
  return attributes_needed
}

func Greedy(attributes_needed, candidate_key []string, candidates map[string][]string) []string {
	var final_candidates []string

  for len(attributes_needed) > 0 {
    var (
      best_candidate string 
      attributes_covered []string
    )

    for _, candidate := range candidate_key {
      attributes := candidates[candidate]
      var covered = equalData(attributes_needed, attributes)
      if len(covered) > len(attributes_covered) {
        best_candidate = candidate 
        attributes_covered = covered 
      }
    }

    attributes_needed = removeData(attributes_needed, attributes_covered)
    final_candidates = append(final_candidates, best_candidate)
  }

  return final_candidates
}

func main() {
	fmt.Println("Hello, Greedy Algorithms and I run at O(n^2)")
}
