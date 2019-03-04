/* Since the outer loop goes around n times, insertion sort must
be a quadratic-time algorithm, i.e. O(n^2). */
package main

import "fmt"

// InsertSort is a basic sorting algorithm
func InsertSort(items []int) []int {
	var j int
	n := len(items)

	for i := 1; i < n; i++ {
		j = i
		for (j > 0) && (items[j] < items[j-1]) {
			// Swap items
			items[j], items[j-1] = items[j-1], items[j]
			j--
		}
	}
	return items
}

func _main() {
	items := []int{5, 2, 4}
	items = InsertSort(items)
	fmt.Println(items)
}
