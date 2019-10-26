package merge_sort

type MergeSorter struct {
}

func (sorter *MergeSorter) Sort(unsortedElements []int) (sortedElements []int) {
	unsortedElementsLength := len(unsortedElements)

	if unsortedElementsLength == 1 {
		return unsortedElements
	}

	var firstHalf []int
	var secondHalf []int
	if unsortedElementsLength == 2 {
		firstHalf = unsortedElements[:1]
		secondHalf = unsortedElements[1:]
	} else {
		firstHalf = unsortedElements[:unsortedElementsLength/2]
		secondHalf = unsortedElements[unsortedElementsLength/2:]
	}

	sortedFirstHalf := sorter.Sort(firstHalf)
	sortedSecondHalf := sorter.Sort(secondHalf)

	return sorter.merge(sortedFirstHalf, sortedSecondHalf)
}

func (sorter *MergeSorter) merge(sortedFirstHalf []int, sortedSecondHalf []int) (sortedElements []int) {
	for len(sortedFirstHalf) > 0 && len(sortedSecondHalf) > 0 {
		firstElement := sortedFirstHalf[0]
		secondElement := sortedSecondHalf[0]

		if firstElement < secondElement {
			sortedElements = append(sortedElements, firstElement)
			sortedFirstHalf = sortedFirstHalf[1:]
		} else {
			sortedElements = append(sortedElements, secondElement)
			sortedSecondHalf = sortedSecondHalf[1:]
		}
	}

	sortedElements = append(sortedElements, sortedFirstHalf...)
	sortedElements = append(sortedElements, sortedSecondHalf...)

	return
}
