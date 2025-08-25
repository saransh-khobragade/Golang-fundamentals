package main

import (
	"fmt"
	"sort"
)

// Arrays
func main() {
	arr := []string{"apple", "banana", "cherry", "berry", "carry"}

	fmt.Println(arr[0]) // apple

	// Length
	fmt.Println(len(arr)) // 5

	// Add item
	arr = append(arr, "pineapple")
	fmt.Println(arr) // [apple banana cherry berry carry pineapple]

	// Insert at index
	arr = append(arr[:2], append([]string{"potatoe"}, arr[2:]...)...)
	fmt.Println(arr) // [apple banana potatoe cherry berry carry pineapple]

	// Extend/merge
	arr2 := []string{"kiwi", "orange"}
	arr = append(arr, arr2...)
	fmt.Println(arr)
	// [apple banana potatoe cherry berry carry pineapple kiwi orange]

	// Remove item (find and remove)
	for i, item := range arr {
		if item == "apple" {
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	fmt.Println(arr)
	// [banana potatoe cherry berry carry pineapple kiwi orange]

	// Remove at index
	arr = append(arr[:0], arr[1:]...)
	fmt.Println(arr) // [potatoe cherry berry carry pineapple kiwi orange]

	// Sort slice
	sort.Strings(arr)
	fmt.Println(arr) // [berry carry cherry kiwi orange pineapple potatoe]
	
	// Sort reverse
	sort.Sort(sort.Reverse(sort.StringSlice(arr)))
	fmt.Println(arr) // [potatoe pineapple orange kiwi cherry carry berry]

	// Reverse
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(arr) // [berry carry cherry kiwi orange pineapple potatoe]

	// Dynamic Array
	dynamicArr := make([]int, 5)
	for i := range dynamicArr {
		dynamicArr[i] = 1
	}
	fmt.Println(dynamicArr)

	// Map a slice
	numArr := []int{1, 2, 3, 4, 5}
	squaredArr := make([]int, len(numArr))
	for i, val := range numArr {
		squaredArr[i] = val * val
	}
	fmt.Println(squaredArr) // [1 4 9 16 25]

	// Filter a slice
	var filteredArr []int
	for _, val := range numArr {
		if val%2 == 0 {
			filteredArr = append(filteredArr, val)
		}
	}
	fmt.Println(filteredArr) // [2 4]

	// Reduce a slice
	reducedValue := 0
	for _, val := range numArr {
		reducedValue += val
	}
	fmt.Println(reducedValue) // 15

	// Destructure slice
	destructureArr := []string{"apple", "banana"}
	a := destructureArr[0]
	b := destructureArr[1]
	fmt.Println(a, b) // apple banana

	// Swap
	a, b = b, a
	fmt.Println(a, b) // banana apple

	// Slices
	sliceArr := []string{"apple", "banana", "cherry", "berry", "carry"}

	// Slicing (start-included : end-excluded)
	fmt.Println(sliceArr[0:1]) // [apple]
	fmt.Println(sliceArr[1:2]) // [banana]

	// slicing (start-included : n)
	fmt.Println(sliceArr[1:]) // [banana cherry berry carry]

	// slicing (0: end-excluded)
	fmt.Println(sliceArr[:2]) // [apple banana]

	// slicing-reverse (0: end-included)
	fmt.Println(sliceArr[:len(sliceArr)-1]) // [apple banana cherry berry]

	// slicing-reverse (start-excluded : end-included)
	fmt.Println(sliceArr[len(sliceArr)-3 : len(sliceArr)-1]) // [cherry berry]
} 