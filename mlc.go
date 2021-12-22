package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now().Local()
	j := julian(n)
	c := longCount(j)
	//f := time.Now().Local()
	for i := range c {
		fmt.Printf("%v", c[i])
		if i != len(c)-1 {
			fmt.Printf(".")
		}
	}
	fmt.Println()
}

func julian(t time.Time) float64 {
	const j = 2453738.4195 // January 1, 1970 (unix epoch)
	u := time.Unix(1136239445, 0)
	const d = float64(86400. * time.Second)
	return j + float64(t.Sub(u))/d
}

func longCount(j float64) [5]int {
	const b = 2456283. // December 21, 2012 (13 b'ak'tun)
	d := int(j - b)
	c := [5]int{13, 0, 0, 0, 0}
	for i := 0; i < d; i++ {
		c[4]++
		if c[4]%20 == 0 {
			e := true
			j := 4
			for e {
				c[j] = 0
				c[j-1]++
				j--
				m := 20
				if j == 3 {
					m = 18
				}
				e = c[j]%m == 0
			}
		}
	}
	return c
}

// http://www.russellcottrell.com/longcount/
