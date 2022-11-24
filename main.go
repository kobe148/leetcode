package main

import (
	"fmt"
	_01 "github.com/raolf/leetcode/1"
	"github.com/raolf/leetcode/_02"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sum := _01.ProductExceptSelf(nums)
	_02.Rotate(nums, 3)
	fmt.Println("Hello World!", sum)
}
