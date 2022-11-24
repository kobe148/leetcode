package main

import (
	"fmt"
	_01 "github.com/raolf/leetcode/1"
	"github.com/raolf/leetcode/_02"
)

func main() {
	nums := []int{3, 5, 5}
	sum := _01.ProductExceptSelf(nums)
	n := _02.ValidMountainArray(nums)
	fmt.Println("Hello World!", sum, n)
}
