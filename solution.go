package square

import (
	"math"
)

const (
	SidesTriangle sides_number = 3
	SidesSquare   sides_number = 4
	SidesCircle   sides_number = 0
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type sides_number int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum sides_number) float64 {
	if sidesNum == SidesTriangle {
		return math.Pow(sideLen, 2) / 2
	}
	if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	}
	if sidesNum == SidesCircle {
		return sideLen / 2 * math.Pi
	} else {
		return 0
	}

}
