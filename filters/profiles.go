package filters

import (
	"fmt"
	"image"
	"sort"
)

const (
	BlackAndWhite = iota + 1
	RetroGame
	Monochrome
	SortedByX
	SortedByY
	Edges
)

var ProfileNames = map[int]string{
	BlackAndWhite: "Black and White",
	RetroGame:     "RetroGame",
	Monochrome:    "Monochrome",
	SortedByX:     "Sorted by X",
	SortedByY:     "Sorted by Y",
	Edges:         "Edges",
}

type Profile struct {
	Actions []func(PixelGrid) PixelGrid
}

var Profiles = map[int]Profile{
	BlackAndWhite: {
		Actions: []func(PixelGrid) PixelGrid{
			blackAndWhite,
		},
	},
	RetroGame: {
		Actions: []func(PixelGrid) PixelGrid{
			retroGame,
		},
	},
	Monochrome: {
		Actions: []func(PixelGrid) PixelGrid{
			monochrome,
		},
	},
	SortedByX: {
		Actions: []func(PixelGrid) PixelGrid{
			func(pg PixelGrid) PixelGrid { return setZoneBasedOnLightness(4, pg) },
			func(pg PixelGrid) PixelGrid { return sortXByZone(4, pg) },
		},
	},
	SortedByY: {
		Actions: []func(PixelGrid) PixelGrid{
			func(pg PixelGrid) PixelGrid { return setZoneBasedOnLightness(4, pg) },
			func(pg PixelGrid) PixelGrid { return sortYByZone(4, pg) },
		},
	},
	Edges: {
		Actions: []func(PixelGrid) PixelGrid{
			func(pg PixelGrid) PixelGrid { return setZoneBasedOnLightness(4, pg) },
			naiveEdgeDetectionByZone,
		},
	},
}

func Run(img *image.RGBA, profileId int) *image.RGBA {
	pg := CopyPixels(img)

	for i := range Profiles[profileId].Actions {
		pg = Profiles[profileId].Actions[i](pg)
	}

	return NewImageRGBA(pg)
}

func GetProfileNames() string {
	keys := make([]int, 0, len(ProfileNames))
	for k := range ProfileNames {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	var names string
	for _, k := range keys {
		names += fmt.Sprintf("\n    %d: %s", k, ProfileNames[k])
	}

	return names
}
