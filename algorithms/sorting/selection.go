/* Selection sort is quadratic, running time is O(n^2)

The sort works by finding the min and placing that at
index 0, it then looks to find the next item to place
at index 1 which should be the 2nd smallest item.
Once it has the index of the 2nd smallest item it swaps
2nd smallest index and index 1.
*/
package main

import "fmt"

func selectionSort(s []int, n int) []int {
	// Index of next minimum
	var min int

	// Loop counters i and j
	for i := 0; i < n; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			// This if is allowing min to hold index of
			// the smallest in range(i, j)
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}
	return s
}

func main() {
	items := []int{5, 2, 4, 7, 6, 1}
	items = selectionSort(items, len(items))
	fmt.Println(items)
	// Prints [1 2 4 5 6 7]
}
