package main

import "math"

var com *object = &object{"COM", nil, nil, 0, true}

type object struct {
	name          string
	orbiting      *object
	orbitedBy     []*object
	nOrbits       int
	hasCalculated bool
}

func (obj *object) setOrbit(otherObj *object) {
	obj.orbiting = otherObj
}

func (obj *object) addOrbitingBody(otherObj *object) {
	obj.orbitedBy = append(obj.orbitedBy, otherObj)
}

func (obj *object) getNumbOrbits() int {
	if obj == com {
		return 0
	} else if obj.hasCalculated {
		return obj.nOrbits
	} else {
		obj.nOrbits = 1 + obj.orbiting.getNumbOrbits()
		obj.hasCalculated = true
		return obj.nOrbits
	}
}

func (obj *object) findShortestPath(name string) int {
	return obj.orbiting.findShortestPathTo(name, obj)
}

func (obj *object) findShortestPathTo(name string, calling *object) int {
	if obj.orbiting.name == name {
		return 2
	}

	shortestPath := math.MaxInt64
	if obj.orbiting != com && obj.orbiting != calling {
		newPathLength := obj.orbiting.findShortestPathTo(name, obj)
		if newPathLength < shortestPath {
			shortestPath = newPathLength
		}
	}

	for _, object := range obj.orbitedBy {
		if object.name == name {
			return 0
		} else if object != calling {
			newPathLength := object.findShortestPathTo(name, obj)
			if newPathLength < shortestPath {
				shortestPath = newPathLength
			}
		}
	}

	if shortestPath == math.MaxInt64 {
		return math.MaxInt64
	}
	return 1 + shortestPath
}

func newObject(name string) *object {
	return &object{name, com, nil, 0, false}
}
