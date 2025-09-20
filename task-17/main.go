package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"sort"
)

func randomSlice(n int) []int {
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(10*n))
	}

	return res
}

func binSearch[T cmp.Ordered](nums []T, target T) int {
	l := 0
	r := len(nums)

	for l < r {
		mid := (l + r) / 2
		switch {
		case target == nums[mid]:
			return mid
		case target < nums[mid]:
			r = mid
		case target > nums[mid]:
			l = mid + 1
		}
	}

	return -1
}

func main() {
	arr := randomSlice(10)
	sort.Ints(arr)

	fmt.Println(arr)
	fmt.Println(binSearch(arr, arr[7]))
	fmt.Println(binSearch(arr, -1))
}
