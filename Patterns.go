package main

import "fmt"

func appendString(s, st string, count int) string {
	for i := 0; i < count; i++ {
		s += st
	}
	return s
}

func main() {
	eqTriangle()
}

// Right Angle Triangle
// *
// **
// ***
// ****
func rightTriangle() {
	for i := 0; i < 4; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// *    *
// **  **
// ******

func HalfButterfly() {
	n := 7 // Number of Columns we want
	star := "*"
	space := "  "
	for i := 1; i < n; i++ {
		fmt.Print(star)
		res := appendString(space, space, n-1-i)
		fmt.Print(res)
		fmt.Print(star)
		fmt.Println()
		star += "*"
	}
}

// *
// **
// ***
// ****
// ****
// ***
// **
// *

func trianglePattern() {
	n := 5
	star := "*"
	for i := 1; i < n; i++ {
		fmt.Println(star)
		star += "*"
	}

	for j := 1; j < n; j++ {
		star = "*"
		star = appendString(star, star, n-1-j)
		fmt.Println(star)
	}
}

//    *
//   ***
//  *****
// *******

func eqTriangle() {
	n := 5
	star := "*"
	spc := " "
	for i := 1; i < n; i++ {
		res := appendString(spc, spc, n-2-i)
		if i == n-1 {
			res = ""
		}
		fmt.Print(res)
		fmt.Print(star)
		fmt.Print(res)
		fmt.Println()
		star += "**"
	}
}
