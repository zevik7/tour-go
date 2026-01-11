package main

/// split demonstrates named return values
// The return statement without arguments returns the named return values (naked return)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // Returns x and y automatically
}
