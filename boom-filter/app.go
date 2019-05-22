package main

import (
	"fmt"
)

func main()  {
	Check3200()
}



// insert 2013, 2010, 2007, 2004, 2001, 1998
//check 0
func Check0() {
	boom := make([]bool, 32)
	fmt.Println(boom)
	hashSet := &HashSet{
		num: 3,
	}

	hashSet.setUp()
	hashSet.hashing(boom, 2013)
	hashSet.hashing(boom, 2010)
	hashSet.hashing(boom, 2007)
	hashSet.hashing(boom, 2004)
	hashSet.hashing(boom, 2001)
	hashSet.hashing(boom, 1998)
	fmt.Println(hashSet.checking(boom, 0))
}


// insert 2013 2010 2007 2004
//check 3200
func Check3200() {
	boom := make([]bool, 32)
	fmt.Println(boom)
	hashSet := &HashSet{
		num: 3,
	}

	hashSet.setUp()
	hashSet.hashing(boom, 2013)
	hashSet.hashing(boom, 2010)
	hashSet.hashing(boom, 2007)
	hashSet.hashing(boom, 2004)
	fmt.Println(hashSet.checking(boom, 3200))
}
// insert  2013, 2010, 2007, 2004, 2001, 1998.
//show true
func ShowTrue1() {
	boom := make([]bool, 32)
	fmt.Println(boom)
	hashSet := &HashSet{
		num: 3,
	}

	hashSet.setUp()
	hashSet.hashing(boom, 2013)
	hashSet.hashing(boom, 2010)
	hashSet.hashing(boom, 2007)
	hashSet.hashing(boom, 2004)
	hashSet.hashing(boom, 2001)
	hashSet.hashing(boom, 1998)
	printTrue(boom)
}

//insert 2013, 2010, 2007
//show true

func ShowTrue2()  {
	boom := make([]bool, 32)
	fmt.Println(boom)
	hashSet := &HashSet{
		num: 3,
	}

	hashSet.setUp()
	hashSet.hashing(boom, 2013)
	hashSet.hashing(boom, 2010)
	hashSet.hashing(boom, 2007)
	printTrue(boom)
}


// show true point
func ShowTrue3() {
	boom := make([]bool, 32)
	fmt.Println(boom)
	hashSet := &HashSet{
		num: 3,
	}
	hashSet.setUp()
	hashSet.hashing(boom, 2010)
	hashSet.hashing(boom, 2013)
	printTrue(boom)
}

// insert 2013, 2010, 2007
// show true
// then insert 2004
// show true
func ShowTrue4()  {
	boom := make([]bool, 32)
	hashSet := &HashSet{
		num: 3,
	}

	hashSet.setUp()
	hashSet.hashing(boom, 2013)
	hashSet.hashing(boom, 2010)
	hashSet.hashing(boom, 2007)
	printTrue(boom)
	fmt.Println("=========== next ============================")
	hashSet.hashing(boom, 2004)
	printTrue(boom)
}

type Hash struct {
	num int
}

func (h *Hash) hashing(v int) int {
	return (h.num*(v*v + v*v*v))%32
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

