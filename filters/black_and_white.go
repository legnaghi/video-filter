package filters

func blackAndWhite(pg PixelGrid) PixelGrid {
	for x := 0; x < pg.MaxX; x++ {
		for y := 0; y < pg.MaxY; y++ {
			if pg.Grid[x][y].Lightness <= 0.5 {
				// #000000
				pg.Grid[x][y].Red = 0
				pg.Grid[x][y].Green = 0
				pg.Grid[x][y].Blue = 0
				pg.Grid[x][y].Alpha = 255
			} else {
				// #ffffff
				pg.Grid[x][y].Red = 255
				pg.Grid[x][y].Green = 255
				pg.Grid[x][y].Blue = 255
				pg.Grid[x][y].Alpha = 255
			}
		}
	}
	return pg
}
