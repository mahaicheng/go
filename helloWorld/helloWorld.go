package main

import (
	"fmt"
)

func main() {
	// // fmt.Println("Hello World")

	// var x, y, z []int
	// x = append(x, 1)
	// y = append(x, 2)
	// z = append(x, 3)
	// fmt.Println(x)
	// fmt.Println(y, z)
	// i, _ := strconv.Atoi(os.Args[1])
	// fmt.Println(fabonaqi(uint64(i)))

	strings := []string{"ma", "mahaicheng", "", "huanghui"}
	fmt.Println(noEmpty(strings))
}

func fabonaqi(i uint64) uint64 {
	if i < 1 {
		return 0
	}
	if i <= 2 {
		return 1
	}
	return fabonaqi(i-2) + fabonaqi(i-1)
}

func nami(i uint16) int {
	return int(i)
}

func noEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings
}
