package whatDistance

import "fmt"

func WhatDistanceFunc (dist int, lapDist int) {
	switch {
	case dist < 0 || lapDist < 0:
		fmt.Println("ERROR: Enter non-negative values")
	case dist == 0:
		fmt.Println("No need to run")
	case lapDist == 0:
		fmt.Println("Ya gonna run forever")
	default:
		curDist := 0
		for i := 0; curDist < dist; i++ {
			curDist += lapDist
			if curDist >= dist {
				fmt.Println(i + 1, "full laps")
			}
		}
	}
}
