package main

import "log"

func main() {
	test := []int{-2, 11, 8, -4, -1, 16, 5, 0}
	log.Println("-2, 11, 8, -4, -1, 16, 5, 0的最大连续子序列和: \n", maxSubsequenceSumGeneral(test))
	log.Println("-2, 11, 8, -4, -1, 16, 5, 0的最大连续子序列和: \n", maxSubsequenceSumBetter(test,0,len(test)-1))
	log.Println("-2, 11, 8, -4, -1, 16, 5, 0的最大连续子序列和: \n", maxSubsequenceSumBest(test))
}

// O(n^2)
func maxSubsequenceSumGeneral(a []int) int {
	maxSum := 0
	for i := 0; i < len(a); i++ {
		temp := 0
		for j := i; j < len(a); j++ {
			temp += a[j]
			if temp > maxSum {
				maxSum = temp
			}
		}
	}
	return maxSum
}

// O(nLogn)
func maxSubsequenceSumBetter(a []int, left int, right int) int {
	if left == right {
		if a[left] > 0 {
			return a[left]
		}else {
			return 0
		}
	}

	center := (left + right) / 2
	maxSumLeft := maxSubsequenceSumBetter(a, left, center)
	maxSumRight := maxSubsequenceSumBetter(a, center + 1, right)

	var (
		tempSumLeftBorder 	int	= 0
		maxSumLeftBorder 	int	= 0
		tempSumRightBorder 	int = 0
		maxSumRightBorder 	int = 0
	)

	for i := center; i >= left; i-- {
		tempSumLeftBorder += a[i]
		if tempSumLeftBorder > maxSumLeftBorder {
			maxSumLeftBorder = tempSumLeftBorder
		}
	}

	for j := center + 1; j <= right; j++ {
		tempSumRightBorder += a[j]
		if tempSumRightBorder > maxSumRightBorder {
			maxSumRightBorder = tempSumRightBorder
		}
	}

	return Max(maxSumLeft, maxSumRight, maxSumLeftBorder + maxSumRightBorder)
}

// O(n) a pretty clever solution
func maxSubsequenceSumBest(a []int) int {
	var (
		maxSum int = 0
		tempSum int = 0
	)

	for i := 0; i < len(a); i++ {
		tempSum += a[i]
		if tempSum > maxSum {
			maxSum = tempSum
		}
		if tempSum < 0 {
			tempSum = 0
		}
	}

	return maxSum
}

func Max(first int, args... int) int {
	for _ , v := range args{
		if first < v {
			first = v
		}
	}
	return first
}
