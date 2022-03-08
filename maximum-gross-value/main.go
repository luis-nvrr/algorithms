package main

import "fmt"

func GrossValue(in, jn, kn int, arr []int) (grval int) {
	n := len(arr)
	if !(1 <= in && in <= jn && jn <= kn && kn <= n+1) {
		panic("gross value error")
	}

	grval = Sum(1, in, arr) - Sum(in, jn, arr) + Sum(jn, kn, arr) - Sum(kn, n+1, arr)
	return grval
}

func Sum(l, r int, arr []int) (sum int) {
	if !(1 <= l && l <= r && r <= len(arr)+1) {
		panic("sum error")
	}

	// use 0 indexed array
	l--
	r--

	for i := l; i < r; i++ {
		sum += arr[i]
	}

	return sum
}

func GenerateValidTriplets(arr []int) (triplets [][]int) {
	var trplt []int
	for i := range arr {
		for j := i; j < len(arr); j++ {
			for k := j; k <= len(arr); k++ {
				trplt = []int{i, j, k}
				triplets = append(triplets, trplt)
			}
		}
	}

	return triplets
}

func MaxGrossValue(arr []int) (max int) {
	triplets := GenerateValidTriplets(arr)
	for i, trplt := range triplets {
		// use 1 based indexing
		in := trplt[0] + 1
		jn := trplt[1] + 1
		kn := trplt[2] + 1

		grossvalue := GrossValue(in, jn, kn, arr)

		if i == 0 {
			max = grossvalue
		}

		if grossvalue > max {
			max = grossvalue
		}

		fmt.Printf("i:%d, j:%d, k:%d ", trplt[0]+1, trplt[1]+1, trplt[2]+1)
		fmt.Printf("Grossvalue: %d \n", grossvalue)
	}
	fmt.Printf("Max gross value is %d", max)
	return max
}

func main() {
	arr := []int{-5, 3, 9, 4}
	MaxGrossValue(arr)
}
