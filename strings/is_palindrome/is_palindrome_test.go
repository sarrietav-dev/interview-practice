package ispalindrome_test

import (
	"testing"

	ispalindrome "github.com/sarrietav-dev/interview-practice/strings/is_palindrome"
)

func TestIsPalindromeOne(t *testing.T) {
	if !ispalindrome.IsPalindrome("A man, a plan, a canal: Panama") {
		t.Error("Expected true, got false")
	}
}

func TestIsPalindromeTwo(t *testing.T) {
  if ispalindrome.IsPalindrome("race a car") {
    t.Error("Expected false, got true")
  }
}

func TestIsPalindromeThree(t *testing.T) {
  if !ispalindrome.IsPalindrome(" ") {
    t.Error("Expected true, got false")
  }
}
