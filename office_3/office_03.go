package office_3

func findRepeatNumber(nums []int) int {
	reMap := make(map[int]int)
	num := 0
	for _, v := range nums {
		if _, isOk := reMap[v]; isOk {
			num = v
			break
		} else {
			reMap[v] = 1
		}
	}
	return num
}
