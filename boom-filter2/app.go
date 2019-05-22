package main

import (
	"fmt"
)

func main()  {
	Check2015()
}

// insert 1975, 1985, 1995, 2005
//check 0
func Check2015() {
	boom := make([]bool, 64)
	fmt.Println(boom)
	hashSet := &HashSet{
		num: 2,
	}

	hashSet.setUp()
	hashSet.hashing(boom, 1975)
	hashSet.hashing(boom, 1985)
	hashSet.hashing(boom, 1995)
	hashSet.hashing(boom, 2005)
	fmt.Println(hashSet.checking(boom, 2015))
}

type Hash struct {
	num int
	filter func(int) int
}

func (h *Hash) hashing(v int) int {

	return h.filter(v)
}

type HashSet struct {
	num int
	hash []*Hash
}

func (hs *HashSet)setUp()  {
	hs.hash = make([]*Hash,hs.num)
	for i:= 0 ;i< hs.num; i++ {
		hs.hash[i] = &Hash{
			num:i+1,
			filter:func(v int) int{
				return i*v%64
			},
		}
	}
}

func (hs *HashSet) hashing(boom []bool, v int)  {

	for _,h := range hs.hash {
		i := h.hashing(v)
		boom[i] = true
	}
}

func (hs *HashSet) checking(boom []bool, v int) bool {

	for _, h := range hs.hash{
		if boom[h.hashing(v)] {
			return true
		}
	}
	return false
}


func printTrue(boom []bool)  {
	for i, v := range boom {
		if v {
			fmt.Print(i)
			fmt.Print(",")
		}

	}

}

