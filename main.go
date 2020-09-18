package main

import (

	"log"
	"math"
)

func main() {
	s1 := 100
	s2 := 100
	s3 := 100
	var count int
	for count < 10000001 {
		rng1 := func() {
			s1 = (171 * s1) % 30269
			s2 = (172 * s2) % 30307
			s3 = (170 * s3) % 30323

		}

		rng1()
		u:=math.Mod(float64(s1)/30269.0 + float64(s2)/30307.0 + float64(s3)/30323.0, 1.0)
		//fmt.Println(count, u)
		log.Println(u)
		count += 1
	}

}