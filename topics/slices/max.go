package slice

import "fmt"

func max(nums []int) (int,error) {

	if len(nums) == 0 {
		return 0, fmt.Errorf("%#v slice has no elements",nums)
	}

	if len(nums) == 1 {
		return nums[0], nil
	}

	max := num[0]


	for _,v := range nums[1:] {

		if v > max {
			max = v
		}
	}

    return max,nil

}