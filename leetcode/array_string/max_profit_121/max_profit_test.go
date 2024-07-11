package max_profit_121

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		nums1  []int
		expect int
	}{
		{
			nums1:  []int{7, 1, 5, 3, 6, 4},
			expect: 5,
		},
		{
			nums1:  []int{7, 6, 4, 3, 1},
			expect: 0,
		},
	}

	for _, tt := range tests {
		t.Run("Testing MaxProfit", func(t *testing.T) {
			assert.Equal(t, tt.expect, MaxProfit(tt.nums1))
		})
	}
}
