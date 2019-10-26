# Merge sort

[![Build Status](https://travis-ci.org/Vehsamrak/merge-sort.svg?branch=master)](https://travis-ci.org/Vehsamrak/merge-sort) [![codecov](https://codecov.io/gh/Vehsamrak/merge-sort/branch/master/graph/badge.svg)](https://codecov.io/gh/Vehsamrak/merge-sort) [![Go Report Card](https://goreportcard.com/badge/github.com/Vehsamrak/merge-sort)](https://goreportcard.com/report/github.com/Vehsamrak/merge-sort)

### merge sort algorithm implementation

## Usage

~~~
unsortedList := []int{3,2,1}
sorter := &MergeSorter{}
sortedList := sorter.Sort(unsortedList)

// sortedList would be: [1, 2, 3]
~~~

## Tests

~~~
go test ./...
~~~