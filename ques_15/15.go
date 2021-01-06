package ques_15

import "sort"

// 三数之和
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i <= len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		flag, arr := FindNum(nums[i], nums[i+1:])
		if flag {
			res = append(res, arr...)
		}

	}
	return res
}

func FindNum(num int, nums []int) (bool, [][]int) {
	flag := false
	data := [][]int{}
	l, r := 0, len(nums)-1
	for l < r {
		if l > 0 && nums[l] == nums[l-1] {
			l++
			continue
		}
		if r < len(nums)-1 && nums[r] == nums[r+1] {
			r--
			continue
		}
		if num+nums[l]+nums[r] == 0 {
			flag = true
			data = append(data, []int{num, nums[l], nums[r]})
			l++
		} else if num+nums[l]+nums[r] > 0 {
			r--
		} else {
			l++
		}

	}
	return flag, data
}
