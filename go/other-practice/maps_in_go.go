package main

/* Go provides us with two methods of initialising maps
 * 
 * One allows you to specify capacity (1), the other (2)
 * allows you to initialize values.
 * With empty maps with zero capacity, it comes down to
 * preference.
 */

// (2) This a literal with a non-empty map
m := map[bool]string{false: "FALSE", true: "TRUE"}

// This is a map literal with no initial values
m := map[T]U{}

// Which is equivalent to
m := make(map[T]U)

// The only difference is that `make` is the only way you can manually
// specify the capacity of the map. In general the language is supposed
// to handle the size of the maps for you, but if you know the map is
// going to hold a maximum of N elements, you can minimize the runtime
// cost by allocating the map with:
// (1)
m := make(map[Y]U, N)
