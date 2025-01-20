package main

func SumAll(numbersToSum ... []int) (sumCollection []int) {
	lenOfNums := len(numbersToSum)
	sumCollection = make([]int, lenOfNums)

	for _, numbers := range numbersToSum {
		sumCollection = append(sumCollection, Sum(numbers))
	}

	return
}