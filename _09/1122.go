package _09

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var m = make(map[int]int)
	for i := range arr2 {
		m[arr2[i]] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		x, okx := m[arr1[i]]
		y, oky := m[arr1[j]]
		if okx {
			if oky {
				return x < y
			} else {
				return true
			}
		} else {
			if oky {
				return false
			} else {
				return arr1[i] < arr1[j]
			}
		}
	})

	return arr1
}
