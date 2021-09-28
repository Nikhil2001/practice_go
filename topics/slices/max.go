package max

import "fmt"

func Max(numbers []int) (int, error) {

	if len(numbers) == 0 {
		return 0, fmt.Errorf("%#v has no numbers", numbers)
	}

	if len(numbers) == 1 {
		return numbers[0], nil
	}

	max := numbers[0]

	for _, v := range numbers[1:] {

		if v > max {
			max = v
		}
	}

	return max, nil
}
