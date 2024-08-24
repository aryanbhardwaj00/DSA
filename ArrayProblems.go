package dsa

import "fmt"

func largestElement(arr []int) int {
	greatest := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= greatest {
			greatest = arr[i]
		}
	}
	return greatest
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func uniqueElements(nums []int) []int {
	var uniqueElement []int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			uniqueElement = append(uniqueElement, nums[i])
		}
	}
	return uniqueElement
}

func leftRotate(nums []int, k int) {
	if k == 0 {
		fmt.Println(nums)
		return
	}
	firstElement := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
	nums[len(nums)-1] = firstElement
	leftRotate(nums, k-1)
}

func linearSearch(nums []int, element int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == element {
			return true
		}
	}
	return false
}

func findUnion(nums1 []int, nums2 []int) []int {
	unionArray := nums1
	for i := 0; i < len(nums1); i++ {
		count := 0
		for j := 0; j < len(nums2); j++ {
			if nums2[i] == nums1[j] {
				count++
			}
		}
		if count == 0 {
			unionArray = append(unionArray, nums2[i])
		}
	}
	return unionArray
}

func missingNum(nums []int) int {
	missingNums := 0
	length := len(nums)
	for i := 0; i <= length; i++ {
		count := 0
		for j := 0; j < length; j++ {
			if nums[j] == i {
				count++
			}
		}
		if count == 0 {
			missingNums = i
		}
	}
	return missingNums
}

func zeroesToEnd(arr []int) []int {
	// Use two Loops
	// First One to Control the inner loop and Move zero to end If found
	// Second loop to compare and swap zero
	// Return the Resulting Array
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] == 0 {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func bubbleSort(nums []int) []int {
	// Use two Loops
	// Outer one to Control the Indexing
	// Inner Loop to Compare and Swap Elements
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func selectionSort(arr []int) []int {
	// Use two Loops
	// Inner Loop to compare and Swap
	// Outer Loop to Control Indexing
	// We have to put the smallest element at First
	// Let Smallest element be first element
	// Compare it with next, if next is smaller , put next in smallest
	// After one iteration , smallest elements Index is found and we swap it with 1st Index
	for i := 0; i < len(arr)-1; i++ {
		smallest := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallest] {
				smallest = j
			}
		}
		arr[i], arr[smallest] = arr[smallest], arr[i]
	}
	return arr
}

func insertionSort(nums []int) []int {
	// Array is Divided in two parts
	// Sorted and Unsorted Array
	// Compare the element of Unsorted Array to Sorted Array
	// If smaller , make space and move it to its place
	for i := 0; i < len(nums)-1; i++ {
		j := i + 1
		for j > 0 {
			if nums[j] < nums[j-1] {
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			}
			j--
		}
	}
	return nums
}

func findUniqueElement(arr []int) int {
	// Use two pointer approach
	// One to store Index of Unique Element
	// Second to find the unique element and iterate over the Array
	// Whenever two elements are not equal , move the index of Unique Element+1
	indexOfUnique := 0
	for i := 1; i < len(arr); i++ {
		if arr[indexOfUnique] != arr[i] {
			i++
			arr[indexOfUnique] = arr[i]
		}
	}
	return indexOfUnique + 1
}

func maxConsecOnes(arr []int) int {
	// One Variable to store the max count till now
	// One variable to count 1's after a 0 was found
	// Compare both and if currentcount is greater than max count then put it in max count
	if len(arr) == 1 {
		return arr[0]
	}
	maxCount := 0
	currentCount := 0
	for _, num := range arr {
		if num == 1 {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 0
		}

	}
	return maxCount
}

func longestSubArray(arr []int, k int) int {
	// Use Nested Loops to iterate through each Sub Array
	// If sum of elements=K , store it in current length and compare current length with max length
	// If greater update Max length
	maxLength := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum == k {
				currentLength := j + 1 - i
				if currentLength > maxLength {
					maxLength = currentLength
				}
			}
		}
	}
	return maxLength
}

func twoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
			}
		}
	}
	return res
}

func twoSumUsingMap(nums []int, target int) []int {
	// We will use Map's found syntax to locate the Index
	// AS 7+2=9 , if we found 7 we have to just look for 2 , if present we will return index of these
	// To find what we need lets substract the element from sum and check if the other element is present
	// 9-7=2 , or if element at index is 4 we will look for 9-4=5
	var res []int
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num

		if index, found := numMap[complement]; found {
			res = append(res, i, index)
		}
		numMap[num] = i
	}
	return res
}

func majorElement(nums []int) int {
	// Use Nested Loops
	// Track Element using Outer Loop
	// Count using Inner Loop
	// If Count>n/2 , return that element
	major := 0
	for i := 0; i < len(nums)-1; i++ {
		count := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
			if count > len(nums)/2 {
				major = nums[i]
			}
		}
	}
	return major
}

func majElement(nums []int) int {
	// Select the element at Index 'i' as candidate
	// Set count to 0
	// Check if count is 0 , if true set it as candidate and count to 1
	// If candidate is already selected , increment its Count
	// Otherwise Decrement Count , Logic here is if we found other elements we will decrement Count to 0
	// After 0 new Candidate will be selected but after all decrements , the element with major element will
	// Still be candidate
	// Return Candidate
	count := 0
	candidate := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
			count = 1
		} else if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func sortColors(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		smallest := i
		for j := i + 0; j < len(nums); j++ {
			if nums[j] < nums[smallest] {
				smallest = j
			}
		}
		nums[i], nums[smallest] = nums[smallest], nums[i]
	}
	return nums
}

func testArr() {
	var arr = make([]int, 5, 5)
	arr[1] = 10
	arr[2] = 20
	fmt.Println(arr)
}

func search(nums []int, target int) int {
	for index, num := range nums {
		if num == target {
			return index
		}
	}
	return -1
}

// Subarray with Max sum
func maxSubArray(nums []int) int {
	maxSoFar := nums[0]
	maxEndHere := nums[0]
	for i := 1; i < len(nums); i++ {
		maxEndHere = max(nums[i], nums[i]+maxEndHere)
		maxSoFar = max(maxSoFar, maxEndHere)
	}
	return maxSoFar
}

// Buy a stock on ith day and sell on other day
// Return the max profit
// Return 0 if no profit
func stockBuySell(prices []int) int {
	// Set a minimum price variable
	// Set a current price variable
	// If current price is smaller than minimun price , set current price to minimum
	// Else compare current profit with max profit and swap
	if len(prices) == 1 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			currentProfit := prices[i] - minPrice
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}
	if minPrice == prices[len(prices)-1] {
		if maxProfit > 0 {
			return maxProfit
		}
		return 0
	}
	return maxProfit
}

func update(elements []int, element int, updElement int) []int {
	// Validation
	validation := validate(elements, element)
	// If true , then updation of Element
	if validation == true {
		for i := 0; i < len(elements); i++ {
			if elements[i] == element {
				elements[i] = updElement
			}
		}
		fmt.Print("Updated Array is:")
		// If False, then returning same array
	} else {
		fmt.Print("ELement not found in Array . Array remains same")
	}
	return elements
}

func validate(elements []int, element int) bool {
	x := 0
	for i := 0; i < len(elements); i++ {
		if elements[i] == element {
			x++
		}
	}
	if x > 0 {
		return true
	}
	return false
}

func searchInsert(arr []int, target int) int {
	if target <= 0 {
		return 0
	}

	if target > arr[len(arr)-1] {
		return len(arr)
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
		if target > arr[i] && target < arr[i+1] {
			return i + 1
		}
	}
	return -1
}

func thirdHighest(arr []int) {
	largest := arr[0]
	secLargest := 0
	thirdHighest := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			thirdHighest = secLargest
			secLargest = largest
			largest = arr[i]
		} else {
			if arr[i] > secLargest && arr[i] != largest {
				secLargest = arr[i]
			} else {
				if arr[i] > thirdHighest && arr[i] != secLargest && arr[i] != largest {
					thirdHighest = arr[i]
				}
			}
		}
	}
	fmt.Println(largest, secLargest, thirdHighest)
}
