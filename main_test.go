package main

import (
	"fmt"
	"testing"
)

func TestRandomMatrix(t *testing.T) {
	tests := []struct {
		h, w, maxNumber, equalRangeMatrix int
	}{
		{2, 2, 100, 4},
		{4, 4, 150, 16},
		{10, 5, 65, 50},
		{50, 50, 2500, 1600},
	}
	for _, test := range tests {
		name := fmt.Sprintf("Case(%d,%d,%d)", test.w, test.h, test.maxNumber)
		t.Run(name, func(t *testing.T) {
			got, err := RandomMatrix(test.w, test.h, test.maxNumber)
			if err != nil {
				t.Errorf("%v", err)
			}

			for _, k := range got {
				for _, j := range k {
					if j > test.maxNumber {
						t.Errorf("got:%d; want:%d", j, test.maxNumber)
					}
				}
			}
			count := 0
			for _, k := range got {
				count += len(k)
			}
			if count < test.equalRangeMatrix {
				t.Errorf("got:%d; want:%d", len(got), test.equalRangeMatrix)
			}

		})
	}
}

func TestUniqNumber(t *testing.T) {
	tests := []struct {
		w, h, maxNumber, uniqCount int
	}{
		{2, 2, 100, 1},
		{4, 4, 150, 1},
		{10, 5, 65, 1},
		{30, 30, 1000, 1},
		{50, 50, 2500, 1},
	}

	for _, test := range tests {
		name := fmt.Sprintf("Case(%d,%d,%d)", test.w, test.h, test.maxNumber)
		t.Run(name, func(t *testing.T) {
			got := UniqNumber(test.w, test.h, test.maxNumber)
			if test.maxNumber < (test.w * test.h) {
				t.Errorf("got:%d; want:%d", test.maxNumber, test.w*test.h)
			}

			if len(got) > (test.w * test.h) {
				t.Errorf("got %d; want %d", len(got), test.w*test.h)
			}
			muniq := map[int]int{}
			for _, k := range got {
				muniq[k]++
			}
			for _, i := range muniq {
				if i >= 2 {
					t.Errorf("got %d; want %d", i, test.uniqCount)
				}
			}
		})
	}
}
