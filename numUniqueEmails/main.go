package main

import "strings"

func main() {
	
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]int)
	sum := 0
	for _, v := range emails {

		vs := strings.Split(v, "@")

		vs[0] = strings.Replace(vs[0],".", "", -1)

		if strings.Contains(vs[0], "+") {
			temp := strings.Split(vs[0], "+")

			vs[0] = temp[0]
		}

		key := vs[0] + vs[1]

		if _, ok := m[key]; !ok {
			m[key]=0
			sum ++
		}
	}
	return sum
}