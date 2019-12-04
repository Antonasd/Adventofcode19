package wire

import (
	"fmt"
	"strconv"
)

type Point struct {
	X, Y int
}

type Intersection struct {
	Intersection Point
	Distance     int
}

type segment struct {
	start Point
	end   Point
}

//Wire : Representation of a wire in the grid.
type Wire struct {
	wireSegments []segment
}

//AddWireSegment : Adds a segment to the end of the wire
func (wire *Wire) AddWireSegment(dir byte, magnitude int) error {
	var newSegment segment
	if wire.wireSegments == nil {
		newSegment.start = Point{0, 0}
		newSegment.end = Point{0, 0}
	} else {
		newSegment = segment{wire.wireSegments[len(wire.wireSegments)-1].end, wire.wireSegments[len(wire.wireSegments)-1].end}
	}

	switch dir {
	case 'U':
		newSegment.end.Y += magnitude
	case 'D':
		newSegment.end.Y -= magnitude
	case 'R':
		newSegment.end.X += magnitude
	case 'L':
		newSegment.end.X -= magnitude
	default:
		return fmt.Errorf("Invalid direction given: %v", dir)
	}

	wire.wireSegments = append(wire.wireSegments, newSegment)
	return nil
}

func (wire *Wire) GetIntersectionsPoints(otherWire Wire) (intersectionPoints []Point) {
	for _, wireSegment := range wire.wireSegments {
		for _, otherSegment := range otherWire.wireSegments {
			intersectPoints, intersects := wireSegment.intersection(otherSegment)
			if intersects {
				intersectionPoints = append(intersectionPoints, intersectPoints...)
			}
		}
	}
	return
}

func (wire *Wire) GetIntersections(otherWire Wire) (uniqueIntersections []Intersection) {
	distance := 0
	intersections := make(map[int]map[int]Intersection)
	for _, wireSegment := range wire.wireSegments {
		for _, otherSegment := range otherWire.wireSegments {
			intersectPoints, intersects := wireSegment.intersection(otherSegment)
			if intersects {
				for _, intersectPoint := range intersectPoints {
					_, hasMap := intersections[intersectPoint.X]
					hasCrossed := false

					if !hasMap {
						intersections[intersectPoint.X] = make(map[int]Intersection)
					} else {
						_, hasCrossed = intersections[intersectPoint.X][intersectPoint.Y]
					}
					if !hasCrossed {
						intersections[intersectPoint.X][intersectPoint.Y] = Intersection{intersectPoint, distance + wireSegment.start.ManhattanDistance(intersectPoint)}
					}
				}
			}
		}
		distance += wireSegment.start.ManhattanDistance(wireSegment.end)
	}

	for _, subMap := range intersections {
		for _, value := range subMap {
			uniqueIntersections = append(uniqueIntersections, value)
		}
	}
	return
}

func (seg *segment) intersection(otherSeg segment) (pointIntersects []Point, intersects bool) {
	intersects = seg.intersectsWith(otherSeg)

	if intersects {
		if (seg.end.X == otherSeg.end.X && seg.start.X == otherSeg.start.X) || (seg.end.Y == otherSeg.end.Y && seg.start.Y == otherSeg.start.Y) {
			var innerSegment segment
			if seg.containsPoint(otherSeg.end) {
				if seg.containsPoint(otherSeg.start) {
					innerSegment = otherSeg
				} else if otherSeg.containsPoint(seg.end) {
					innerSegment.end = seg.end
					innerSegment.start = otherSeg.end
				} else {
					innerSegment.end = seg.start
					innerSegment.start = otherSeg.end
				}
			} else if seg.containsPoint(otherSeg.start) {
				if otherSeg.containsPoint(seg.end) {
					innerSegment.end = seg.end
					innerSegment.start = otherSeg.start
				} else {
					innerSegment.end = seg.start
					innerSegment.start = otherSeg.start
				}
			} else {
				innerSegment = *seg
			}
			pointIntersects = append(pointIntersects, innerSegment.getPointsInSegment()...)
		} else {
			if seg.end.X == seg.start.X {
				pointIntersects = append(pointIntersects, Point{seg.end.X, otherSeg.end.Y})
			} else {
				pointIntersects = append(pointIntersects, Point{otherSeg.end.X, seg.end.Y})
			}
		}
	}
	return
}

func (seg *segment) intersectsWith(otherSeg segment) bool {
	if seg.end.X == seg.start.X {
		if otherSeg.end.X == otherSeg.start.X {
			return seg.containsPoint(otherSeg.start) || seg.containsPoint(otherSeg.end) || otherSeg.containsPoint(seg.start) || otherSeg.containsPoint(seg.end)
		}
		return seg.containsPoint(Point{seg.end.X, otherSeg.end.Y}) && otherSeg.containsPoint(Point{seg.end.X, otherSeg.end.Y})
	}

	if otherSeg.end.Y == otherSeg.start.Y {
		return seg.containsPoint(otherSeg.start) || seg.containsPoint(otherSeg.end) || otherSeg.containsPoint(seg.start) || otherSeg.containsPoint(seg.end)
	}
	return seg.containsPoint(Point{otherSeg.end.X, seg.end.Y}) && otherSeg.containsPoint(Point{otherSeg.end.X, seg.end.Y})
}

func (seg *segment) containsPoint(point Point) bool {
	if point.X == seg.end.X && point.X == seg.start.X {
		return !((point.Y < seg.end.Y && point.Y < seg.start.Y) || (point.Y > seg.end.Y && point.Y > seg.start.Y))
	} else if point.Y == seg.end.Y && point.Y == seg.start.Y {
		return !((point.X < seg.end.X && point.X < seg.start.X) || (point.X > seg.end.X && point.X > seg.start.X))
	}
	return false
}

func (seg *segment) getPointsInSegment() (points []Point) {
	points = append(points, seg.start)

	if seg.end.X == seg.start.X {
		length := abs(seg.end.Y - seg.start.Y)
		if length == 0 {
			return
		}
		dir := (seg.end.Y - seg.start.Y) / length

		for i := 1; i <= length; i++ {
			points = append(points, Point{seg.start.X, seg.start.Y + i*dir})
		}
	} else {
		length := abs(seg.end.X - seg.start.X)
		dir := (seg.end.X - seg.start.X) / length

		for i := 1; i <= length; i++ {
			points = append(points, Point{seg.start.X + i*dir, seg.start.Y})
		}
	}
	return
}

func (point *Point) ManhattanDistance(otherPoint Point) int {
	return abs(point.X-otherPoint.X) + abs(point.Y-otherPoint.Y)
}

func (point *Point) ManhattanDistanceOrigo() int {
	return abs(point.X) + abs(point.Y)
}

//NewWire : Creates a wire from the wire description given in the input.
func NewWire(wireDesc []string) (*Wire, error) {
	var newWire *Wire = new(Wire)

	for _, desc := range wireDesc {
		if desc != "" {
			magnitude, convErr := strconv.Atoi(desc[1:])
			if convErr != nil {
				return nil, convErr
			}

			wireErr := newWire.AddWireSegment(desc[0], magnitude)
			if wireErr != nil {
				return nil, wireErr
			}
		}
	}
	return newWire, nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
