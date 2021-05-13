package main

func SameBsts(arrayOne, arrayTwo []int) bool {
	// if not the same length, then they can't be the same tree
	if len(arrayOne) != len(arrayTwo) {
		return false
	}

	if len(arrayOne) == 0 && len(arrayTwo) == 0 {
		return true
	}

	// if roots are not the same, they are not the same tree
	if arrayOne[0] != arrayTwo[0] {
		return false
	}

	leftOne := GetLeftSubtree(arrayOne)
	leftTwo := GetLeftSubtree(arrayTwo)
	rightOne := GetRightSubtree(arrayOne)
	rightTwo := GetRightSubtree(arrayTwo)

	return SameBsts(leftOne, leftTwo) && SameBsts(rightOne, rightTwo)
}

func GetLeftSubtree(arr []int) []int {
	left := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[0] {
			left = append(left, arr[i])
		}
	}
	return left
}

func GetRightSubtree(arr []int) []int {
	right := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] >= arr[0] {
			right = append(right, arr[i])
		}
	}
	return right
}
