package main

import (
	"fmt"
)

func main() {

	A := []int {1,2,3,4,4,5,5,0,12}

	fmt.Println(Solution(A))
}



func Solution(A []int) int {
	result := 1

	// start from fist value
	for i, v:= range A {
		m := make(map[int]bool)
		m[v] = true
		for j:= i; j <len(A); j++ {
			// only allow contain 2 values
			if len(m) > 2 || j== len(A)-1{
				//max value with i
				if result < j {
					result = j
				}
				break
			}else if _, ok:= m[A[j]]; !ok {
				m[A[j]] = true
			}
			fmt.Println(m)
		}
	}

	return result
}
