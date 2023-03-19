package main

import "sort"

// Cho một mảng là giá cổ phiếu qua các ngày viết chương trình mua cổ phiếu vào ngày giá thấp và bán ra ngày giá
// cao để có lãi, nếu không xảy ra được trường hợp lãi return 0
func sortArr(arr []int)int{
	sort.Ints(arr)
	return arr[0]
}
func maxProfit(arr []int) {
	left := 0
	right := len(arr) - 1
	mid := len(arr) / 2
	if arr[mid] >
}
func main() {

}
