package collatzconjecture

import "errors"

// CollatzConjecture finds the number of steps required to reach 1 using collatz conjecture
func CollatzConjecture(num int) (int, error) {
	if num < 1 {
		return 0, errors.New("Number cannot be less than one")
	}

	steps := 0
	for {
		if num == 1 {
			return steps, nil
		}

		if num%2 == 0 {
			num = num / 2
		} else {
			num = (num * 3) + 1
		}
		steps++
	}
}
