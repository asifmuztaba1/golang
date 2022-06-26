package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	_, err := fmt.Scanln(&t)
	if err != nil {
		return
	}
	i := 0
	for i < t {
		var s, sub string

		_, err := fmt.Scanln(&s)
		if err != nil {
			return
		}
		for j := 0; j < len(s); j++ {
			sub = s[j+1 : len(s)]

			if strings.Contains(sub, string(s[j])) != true {
				s = s[j:len(s)]
				break
			}

		}
		fmt.Println(s)
		i++
	}

}
