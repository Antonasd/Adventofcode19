package main

import (
	"adventofcode19/day3/wire"
	"adventofcode19/util"
	"fmt"
	"math"
	"strings"
)

func main() {
	input, inputError := util.GetInput()
	if inputError != nil {
		fmt.Println(inputError.Error())
		return
	}

	var wires []wire.Wire

	for _, wireDef := range strings.Split(input, "\n") {
		newWire, wireErr := wire.NewWire(strings.Split(wireDef, ","))
		if wireErr != nil {
			fmt.Println(wireErr.Error())
			return
		}
		wires = append(wires, *newWire)
	}

	var intrsctPoints []wire.Point
	for i := 0; i < len(wires); i++ {
		for j := 0; j < len(wires); j++ {
			if j != i {
				intrsctPoints = append(intrsctPoints, wires[i].GetIntersectionsPoints(wires[j])...)
			}
		}
	}

	closestDistance := math.MaxInt64
	for _, point := range intrsctPoints {
		distance := point.ManhattanDistance(wire.Point{0, 0})
		if distance < closestDistance && distance != 0 {
			closestDistance = distance
		}
	}

	fmt.Println("Distance to intersection closest to origo: ", closestDistance)

	intersectionDistances := make(map[int]map[int]int)
	for i := 0; i < len(wires); i++ {
		for j := 0; j < len(wires); j++ {
			if j != i {
				intersections := wires[i].GetIntersections(wires[j])
				for _, intersc := range intersections {
					if !(intersc.Intersection.X == 0 && intersc.Intersection.Y == 0) {
						_, hasMap := intersectionDistances[intersc.Intersection.X]
						if !hasMap {
							intersectionDistances[intersc.Intersection.X] = make(map[int]int)
						}
						intersectionDistances[intersc.Intersection.X][intersc.Intersection.Y] += intersc.Distance
					}
				}
			}
		}
	}

	minDistance := math.MaxInt64
	for _, submap := range intersectionDistances {
		for _, distance := range submap {
			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	fmt.Println("Shortest combined distance to intersection: ", minDistance)
}
