package sorts

func SortSelection2(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		minIndex := i
		maxIndex := 0
		for j := i + 1; j < len(ar)-i; j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
			if ar[j] > ar[maxIndex] {
				maxIndex = j
			}
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
		ar[len(ar)-i-1], ar[maxIndex] = ar[maxIndex], ar[len(ar)-i-1]
	}
}

func SortSelectionByMax(ar []int) {
	for i := len(ar) - 1; i >= 0; i-- {
		maxIndex := i
		for j := i - 1; j >= 0; j-- {
			if ar[j] > ar[maxIndex] {
				maxIndex = j
			}
		}
		ar[i], ar[maxIndex] = ar[maxIndex], ar[i]
	}
}

func SortSelection(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}

func BubbleSortRecursive(ar []int) {
	if len(ar) <= 1 {
		return
	}
	for i := 0; i < len(ar)-1; i++ {
		if ar[i] > ar[i+1] {
			ar[i], ar[i+1] = ar[i+1], ar[i]
		}
	}
	BubbleSortRecursive(ar[:len(ar)-1])
}

func BubbleSortReversed(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		mark := false
		for j := 0; j < len(ar)-1; j++ {
			if ar[j] < ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
				mark = true
			}
		}
		if !mark {
			break
		}
	}
}

func BubbleSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		mark := false
		for j := 0; j < len(ar)-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
				mark = true
			}
		}
		if !mark {
			break
		}
	}
}

func InsertionSort(ar []int) {
	i := 1
	for i < len(ar) {
		j := i
		for j >= 1 && ar[j] < ar[j-1] {
			ar[j], ar[j-1] = ar[j-1], ar[j]
			j--
		}
		i++
	}
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	resIn, leftIn, rightIn := 0, 0, 0

	for leftIn < len(left) && rightIn < len(right) {
		if left[leftIn] < right[rightIn] {
			result[resIn] = left[leftIn]
			resIn++
			leftIn++
		} else {
			result[resIn] = right[rightIn]
			resIn++
			rightIn++
		}
	}

	for resIn < len(result) {
		if leftIn != len(left) {
			result[resIn] = left[leftIn]
			resIn++
			leftIn++
		} else {
			result[resIn] = right[rightIn]
			resIn++
			rightIn++
		}
	}

	return result
}

func MergeSort(ar []int) []int {
	if len(ar) == 1 {
		return ar
	}

	left := MergeSort(ar[:len(ar)/2])
	right := MergeSort(ar[len(left):])

	return merge(left, right)
}

func QuickSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	pivot := ar[0]
	var left, right []int
	for _, num := range ar[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	sorted := append(QuickSort(left), pivot)
	sorted = append(sorted, QuickSort(right)...)
	return sorted
}
