package main

import (
	"sort"
	"fmt"
)

func main() {

	A := []int{1, 3, 6, 4, 1, 2}

	fmt.Println(Solution(A))

}

func Solution(A []int) int {

	sort.Ints(A)

	if len(A) <= 2  {
		if A[len(A)-1] > 0{
			return getResult(A[len(A)-1] +1)
		}else {
			return 1
		}
	}

	for i := 1; i< len(A); i++ {
		if  A[i-1] > 0 && A[i] - A[i-1] >1{

			if A[i-1] <1  {
				return 1
			}
			return getResult(A[i-1]+1)
		}
	}

	if A[len(A)-1] >0 {
		return getResult(A[len(A)-1] +1)
	}

	return 1
}

func getResult(r int) int  {
	if r >1000000   {
		return 1
	}

	return r
}