package main

import (
	"adventofcode19/util"
	"fmt"
	"strings"
)

func main() {
	objetcts := make(map[string]*object)

	fmt.Println("Please provide the path to your puzzle input: ")
	input, inputError := util.GetFileInput()
	if inputError != nil {
		fmt.Println(inputError.Error())
	}

	orbits := strings.Split(input, "\n")

	for _, orbit := range orbits {
		if orbit != "" {
			orbitComp := strings.Split(orbit, ")")

			orbObject, hasOrbObject := objetcts[orbitComp[1]]
			if !hasOrbObject {
				orbObject = newObject(orbitComp[1])
				objetcts[orbitComp[1]] = orbObject
			}

			if orbitComp[0] != "COM" {
				object, hasObject := objetcts[orbitComp[0]]
				if !hasObject {
					object = newObject(orbitComp[0])
					objetcts[orbitComp[0]] = object
				}
				orbObject.setOrbit(object)
				object.addOrbitingBody(orbObject)
			}
		}
	}

	totalNumbOrbits := 0
	shortestPath := 0
	for _, obj := range objetcts {
		nOrbits := obj.getNumbOrbits()
		totalNumbOrbits += nOrbits
		if obj.name == "YOU" {
			shortestPath = obj.findShortestPath("SAN")
		}
	}

	fmt.Println("The total number of orbits is ", totalNumbOrbits)
	fmt.Println("The shortest path to SAN is ", shortestPath)
}
