package validparenthesis_test

import (
	"testing"

	validparenthesis "github.com/sarrietav-dev/interview-practice/stacks/valid_parenthesis"
)

func TestValidParenthesis1(t *testing.T) {
  result := validparenthesis.IsValid("()")

  if !result {
    t.Errorf("Expected to be %v but got %v", true, result)
  }
}

func TestValidParenthesis2(t *testing.T) {
  result := validparenthesis.IsValid("()[]{}")

  if !result {
    t.Errorf("Expected to be %v but got %v", true, result)
  }
}

func TestValidParenthesis3(t *testing.T) {
  result := validparenthesis.IsValid("(]")

  if result {
    t.Errorf("Expected to be %v but got %v", false, result)
  }
}
