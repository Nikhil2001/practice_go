package max

import "testing"

func TestMax(t *testing.T) {

	testInputs := []struct {
		name           string
		inputSlice     []int
		expectedOutput int
	}{
		{"test1",
			[]int{1, 2, 3},
			3,
		},
		{"test2",
			[]int{-1, -2, -3},
			-3,
		},
		{"test3",
			[]int{-1},
			-1,
		},
		{"test4",
			[]int{},
			0,
		},
	
	}

	for _, tt := range testInputs {
		f := func(t *testing.T) {
			t.Logf("testing Max() with input %#v",tt.inputSlice)
			if v, _ := Max(tt.inputSlice); tt.expectedOutput != v {
				t.Errorf("max returns %d not %d", v,tt.expectedOutput)
			}
		}
		t.Run(tt.name, f)
	}
}
