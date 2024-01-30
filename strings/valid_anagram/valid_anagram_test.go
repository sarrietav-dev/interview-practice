package validanagram_test

import (
	"testing"

	validanagram "github.com/sarrietav-dev/interview-practice/strings/valid_anagram"
)

func TestValidAnagramOne(t *testing.T) {
  s := "anagram"
  q := "nagaram"

  if !validanagram.IsAnagram(s, q) {
    t.Errorf("IsAnagram(%s, %s) = false; want true", s, q)
  }
}

func TestValidAnagramTwo(t *testing.T) {
  s := "rat"
  q := "car"

  if validanagram.IsAnagram(s, q) {
    t.Errorf("IsAnagram(%s, %s) = true; want false", s, q)
  }
}
