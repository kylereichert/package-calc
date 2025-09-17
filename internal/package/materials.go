package materials

import (
	"math"
)

type Materials struct {
	Spruce2x4  int
	Treated2x4 int
	Spruce2x6  int
	Treated2x6 int
	Spruce2x8  int
	Treated2x8 int
	Rebar10M   int
	Rebart15M  int
	Tubes      int
}

func evenCeiling(x float64) int {
	y := int(math.Ceil(x))
	return y + (y % 2)
}

type builder interface {
	Calc(footage int, grid int, openings int) Materials
}

type GreenCedar struct{}

func (GreenCedar) Calc(f float64, g float64, o float64) Materials {
	/*
		This will need to be more flexible, this is moreso a framework to work from
		Could consider a job struct or something to allow for overrides, options in
		case of walkouts, etc.
	*/
	const ladderFactor float64 = 6
	const rebarFactor float64 = 3
	const openFactor float64 = 1.5

	m := Materials{
		Spruce2x4:  20,
		Treated2x4: evenCeiling(f / ladderFactor),
		Treated2x8: int(math.Ceil(o*openFactor)) + 1,
		Rebar10M:   evenCeiling((f / rebarFactor) + g),
	}

	return m

}
