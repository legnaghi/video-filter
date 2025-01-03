package filters

func naiveEdgeDetectionByZone(pg PixelGrid) PixelGrid {
	for x := 1; x < pg.MaxX-1; x++ {
		for y := 1; y < pg.MaxY-1; y++ {
			center := pg.Grid[x][y]

			if center.Zone != pg.Grid[x-1][y-1].Zone || center.Zone != pg.Grid[x-1][y].Zone || center.Zone != pg.Grid[x-1][y+1].Zone ||
				center.Zone != pg.Grid[x][y-1].Zone || /*center == center || */ center.Zone != pg.Grid[x][y+1].Zone ||
				center.Zone != pg.Grid[x+1][y-1].Zone || center.Zone != pg.Grid[x+1][y].Zone || center.Zone != pg.Grid[x+1][y+1].Zone {

				pg.Grid[x][y].Red = 0
				pg.Grid[x][y].Green = 0
				pg.Grid[x][y].Blue = 0
			} else {
				pg.Grid[x][y].Red = 255
				pg.Grid[x][y].Green = 255
				pg.Grid[x][y].Blue = 255
			}
		}
	}

	return pg
}
