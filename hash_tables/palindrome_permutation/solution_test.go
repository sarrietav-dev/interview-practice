package palindromepermutation_test

import (
	"testing"

	pp "github.com/sarrietav-dev/interview-practice/hash_tables/palindrome_permutation"
)

func TestMain(t *testing.T) {
	if !pp.PalindomePermutation("Tact Coa") {
		t.Errorf("Expected to be true for 'Tact Coa' but got false")
	}

	if pp.PalindomePermutation("Tact Coaa") {
		t.Errorf("Expected to be false for 'Tact Coaa' but got true")
	}

	if !pp.PalindomePermutation("Tact Coa ") {
		t.Errorf("Expected to be true for 'Tact Coa ' but got false")
	}

	if !pp.PalindomePermutation("Race Car") {
		t.Errorf("Expected to be true for 'Race Car' but got false")
	}

	if !pp.PalindomePermutation("A man, a plan, a canal - Panama") {
		t.Errorf("Expected to be true for 'A man, a plan, a canal - Panama' but got false")
	}

	if pp.PalindomePermutation("") {
		t.Errorf("Expected to be false for '' but got true")
	}
}
