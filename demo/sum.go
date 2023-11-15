package demo

func Sum(nums []uint) uint {
	var result uint

	for _, num := range nums {
		result += num
	}

	return result
}
