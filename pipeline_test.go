package funcpipe

import "testing"

func TestFilter(t *testing.T) {
	t.Run("FilterEvenNumbers", func(t *testing.T) {
		// given
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		// when
		res := NewPipe(numbers).
			Filter(func(a int) bool { return a%2 == 0 }).
			Collect()

		// then
		if len(res) != 5 {
			t.Errorf("Should only have 5 elements left")
		}
		for _, e := range res {
			if e%2 != 0 {
				t.Errorf("Should only have 5 elements left, have %d elements left", len(res))
			}
		}
	})

	t.Run("FilterEvenNumbersChainedCalls", func(t *testing.T) {
		// given
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		// when
		res := NewPipe(numbers).
			Filter(func(a int) bool { return a != 1 }).
			Filter(func(a int) bool { return a != 3 }).
			Filter(func(a int) bool { return a != 5 }).
			Filter(func(a int) bool { return a != 7 }).
			Filter(func(a int) bool { return a != 9 }).
			Collect()

		// then
		if len(res) != 5 {
			t.Errorf("Should only have 5 elements left, have %d elements left", len(res))
		}
		for _, e := range res {
			if e%2 != 0 {
				t.Errorf("Should only have even numbers!")
			}
		}
	})
}
