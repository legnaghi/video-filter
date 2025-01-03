package filters

func setZoneBasedOnLightness(zoneAmount int, pg PixelGrid) PixelGrid {
	zones := make(map[int]float64, zoneAmount)

	zoneStep := float64(1 / zoneAmount)

	for i := range zoneAmount {
		zones[i] = 0 + zoneStep*float64(i+1)
	}

	for x := 0; x < pg.MaxX; x++ {
		for y := 0; y < pg.MaxY; y++ {
			for i := range zoneAmount {
				if pg.Grid[x][y].Lightness <= zones[i] {
					pg.Grid[x][y].Zone = i
					break
				}
			}
		}
	}

	return pg
}
