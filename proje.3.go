package main

import "fmt"

func main() {
	var x, y = 16, 15
	topla(x, y)
	çikar(x, y)
	çarp(x, y)
	böl(x, y)
	modal(x, y)
	arttir(x, y)
	azalt(x, y)
}
func topla(x int, y int) {
	fmt.Printf("x + y  = %d\n", x+y)
}
func çikar(x int, y int) {
	fmt.Printf("x - y = %d\n", x-y)
}
func çarp(x int, y int) {
	fmt.Printf("x * y = %d\n", x*y)
}
func böl(x int, y int) {
	fmt.Printf("x / y = %d\n", x/y)
}
func modal(x int, y int) {
	fmt.Printf("x mod y = %d\n", x%y)
}
func arttir(x int, y int) {
	x++
	fmt.Printf("x++ = %d\n", x)
}
func azalt(x int, y int) {
	y--
	fmt.Printf("y-- = %d\n", y)
}
