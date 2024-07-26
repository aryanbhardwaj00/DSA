package main

import (
	"fmt"
	"math"
	"strconv"
)

func revNum(num int) {
	// Check for single digit num
	if num%10 == num {
		fmt.Println(num)
	}

	numCopy := num
	var result string
	for numCopy > 0 {
		x := strconv.Itoa(numCopy % 10)
		result += x
		numCopy /= 10
	}
	fmt.Println("Reverse of", num, "is", result)
}

func isPalindrome(num int) bool {
	// Check for single digit num
	if num%10 == num {
		return false
	}
	// Check Palindrome
	numCopy := num
	var result string
	for numCopy > 0 {
		x := strconv.Itoa(numCopy % 10)
		result += x
		numCopy /= 10
	}
	fResult, err := strconv.Atoi(result)
	if err != nil {
		fmt.Print("Error", err)
	} else if num == fResult {
		return true
	}
	return false
}

func revTable(n int) {
	mul := 10
	for mul > 0 {
		fmt.Println(n, "*", mul, "=", n*mul)
		mul--
	}
}

func revNum2(num int) {
	result := 0
	for num > 0 {
		num1 := num % 10
		result = result*10 + num1
		num /= 10
	}
	fmt.Println(result)
}

func greaterNum(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func GCD(num1 int, num2 int) int {
	greater := greaterNum(num1, num2)
	greatest := 0
	for i := 1; i < greater; i++ {
		if num1%i == 0 && num2%i == 0 {
			greatest = i
		}
	}
	if greatest == 0 {
		return 1
	}
	return greatest
}

func lenInt(num int) int {
	length := 0
	for num > 0 {
		num = num / 10
		length++
	}
	return length
}

func armStrong(num int) bool {
	// Lenght to know Power
	length := lenInt(num)
	// Create a Copy
	cNum := num
	// Sum Variable to store value
	sum := 0
	// Iterate through each number and raise it to power= Length
	for num > 0 {
		tempNum := num % 10
		result := math.Pow(float64(tempNum), float64(length))
		// Add it to sum
		sum += int(result)
		num /= 10
	}
	// return sum
	return sum == cNum
}

func sumOfDivisors(num int) int {
	// Find the divisors og given num
	// Add them
	// Give recursive call for Num-1
	// Add them to sum of Divisors of Num
	if num == 0 {
		return 0
	}
	sum := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	fResult := sum + sumOfDivisors(num-1)
	return fResult
}

func LCM(num1 int, num2 int) int {
	// Find the greater number out of two
	// Loop until greater number
	// If "i" divides either number,store it in "Multiplier"
	// Return Multiplier
	muliplier := 1
	greater := greaterNum(num1, num2)
	for i := 2; i < greater; i++ {
		if num1%i == 0 || num2%i == 0 {
			muliplier *= i
		}
	}
	return muliplier
}

func recursion(num int) {

	if num == 0 {
		return
	}
	fmt.Println("Hello")
	recursion(num - 1)
}

func recursion1(num int) {
	if num == 0 {
		return
	}
	fmt.Println(num)
	recursion1(num - 1)
}

func recursion2(num int) {
	if num == 1 {
		fmt.Println(1)
		return
	}
	recursion2(num - 1)
	fmt.Println(num)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	result := num * factorial(num-1)
	return result
}

func fibonacci(init int, sec int, num int, sum int) {
	// Add initial and Second Value to get next value
	// Pass Second as Initial Values and Sum as Second Value
	// Decrement num in each recursive call
	if num == 0 {
		fmt.Println(sum)

	}
	sum = init + sec
	fibonacci(sec, sum, num-1, sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CountChocolates(n int) int {
	// For every 3 he will get extra 1
	// 15 / 3 = 5 + 0
	// 5 / 3 = 1 + 2
	// 3 / 3 = 1
	// Divide N by 3 and add remainder
	// Divide N until it becomes N<3

	count := 0
	remanider := 0

	for n >= 3 {
		count += n / 3
		remanider = n % 3
		n = n/3 + remanider
		// if n < 3 {
		// 	break
		// }
	}
	return count
}

func winnerFriend(n, k int) {
	friend := make([]int, n)
	for i := 0; i < n; i++ {
		friend[i] = i + 1
	}

	current := 0
	for len(friend) > 1 {
		current = (current + k - 1) % len(friend)
		friend = append(friend[:current], friend[current+1:]...)
	}
	fmt.Println(friend[0])
}

func distanceToOffice(homeToOffice, homeToConstruction, homeToEnginner int) int {
	if homeToConstruction < 0 || homeToConstruction > homeToOffice {
		return homeToOffice
	}
	if homeToEnginner < 0 {
		return homeToConstruction + (homeToConstruction - homeToEnginner) + (-homeToEnginner + homeToOffice)
	}
	if homeToEnginner > homeToConstruction {
		fmt.Println("Not Possible to go to office")
		return 0
	}
	return homeToConstruction + (homeToConstruction - homeToEnginner) + (homeToOffice - homeToEnginner)
}
