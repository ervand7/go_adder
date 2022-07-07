package go_adder

func Add(nums ...int) (result int) {
	for _, value := range nums {
		result += value
	}

	return
}
