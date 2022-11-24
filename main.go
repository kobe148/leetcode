package main

import (
	"fmt"
	_01 "github.com/raolf/leetcode/1"
	"github.com/raolf/leetcode/_02"
)

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	sum := _01.ProductExceptSelf(nums)
	res := _02.SummaryRanges(nums)
	fmt.Println(sum)
	fmt.Println("Hello World!", res)
}
