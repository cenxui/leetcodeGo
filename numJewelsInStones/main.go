package main

func main() {
	
}

func numJewelsInStones(J string, S string) int{
	m := make(map[string] int)

	for _, r := range J {
		m[string(r)] = 0
	}

	sum := 0
	for _, s := range S {
		key := string(s)
		if _, ok := m[key];ok {
			sum++
		}
	}

	return sum
}