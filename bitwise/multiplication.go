package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(100)
	var res []int
	res = append(res, n)
	res = append(res, n<<1)      // multiplication by 2
	res = append(res, n<<1+n)    // multiplication by 3
	res = append(res, n<<2)      // multiplication by 4
	res = append(res, n<<2+n)    // multiplication by 5
	res = append(res, n<<2+n<<1) // multiplication by 6
	res = append(res, n<<3-n)    // multiplication by 7
	res = append(res, n<<3)      // multiplication by 8

	for i, v := range res {
		fmt.Println("number =", n, "multiplication by", i+1, "equals to", v)
	}
}
