package merge_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestMergeSorter(test *testing.T) {
	suite.Run(test, new(MergeSorterTestSuite))
}

type MergeSorterTestSuite struct {
	suite.Suite
}

func (suite *MergeSorterTestSuite) Test_Sort_unsortedIntegers_sortedIntegers() {
	for id, dataset := range suite.provideIntegerList() {
		sorter := &MergeSorter{}

		sorted := sorter.Sort(dataset.unsorted)

		assert.Equal(suite.T(), dataset.sorted, sorted, fmt.Sprintf("Dataset #%v %#v", id, dataset))
	}
}

func (suite *MergeSorterTestSuite) provideIntegerList() []struct {
	unsorted []int
	sorted   []int
} {
	return []struct {
		unsorted []int
		sorted   []int
	}{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{4, 5, 6}, []int{4, 5, 6}},
		{[]int{3, 5, 2, 1, 7, 8, 9, 4, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{3, 5, 2, 1, 7, 8, 10, 9, 4, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
}
