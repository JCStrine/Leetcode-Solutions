package main

func twoSum(nums []int, target int) []int {

	hash := make(map[int]int)

	for i, s := range nums {
		_, exists := hash[s]

		if exists {
			return []int{i, hash[s]}
		}

		hash[target-s] = i
	}

	return nil
}

//https://leetcode.com/problems/two-sum/
