package maxprofit_test

import (
	"testing"

	maxprofit "github.com/sarrietav-dev/interview-practice/arrays/max_profit"
)

func TestMaxProfit1(t *testing.T) {
	result := maxprofit.MaxProfit([]int{
		7, 1, 5, 3, 6, 4,
	})

	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestMaxProfit2(t *testing.T) {
	result := maxprofit.MaxProfit([]int{
		7, 6, 4, 3, 1,
	})

	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
