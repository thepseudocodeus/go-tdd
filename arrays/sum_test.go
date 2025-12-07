package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 30}

		got := Sum(numbers)
		want := 36

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}
