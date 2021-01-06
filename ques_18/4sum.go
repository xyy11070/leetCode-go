package ques_18

import "sort"

func fourSum(nums []int, target int) [][]int {
	// 先排序
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i := 0; i < (n-3) && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		// 跳过重复、 nums[i]+nums[n-3]+nums[n-2]+nums[n-1] 是当前最大值
		if i > 0 && (nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target) {
			continue
		}
		for j := i + 1; j < (n-2) && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			// 跳过重复、nums[i]+nums[j]+nums[n-2]+nums[n-1] 是当前最大值
			if j > i+1 && (nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target) {
				continue
			}
			for l, r := j+1, n-1; l < r; {
				if sums := nums[i] + nums[j] + nums[l] + nums[r]; sums == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l++; l < r && nums[l] == nums[l-1]; l++ {
						// 跳过重复
					}
					for r--; l < r && nums[r] == nums[r+1]; r-- {
						// 跳过重复
					}
				} else if sums < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}
