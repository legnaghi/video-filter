package filters

func monochrome(pg PixelGrid) PixelGrid {
	for x := 0; x < pg.MaxX; x++ {
		for y := 0; y < pg.MaxY; y++ {
			grey := uint8(0.299*float32(pg.Grid[x][y].Red) + 0.587*float32(pg.Grid[x][y].Green) + 0.114*float32(pg.Grid[x][y].Blue))
			pg.Grid[x][y].Red = grey
			pg.Grid[x][y].Green = grey
			pg.Grid[x][y].Blue = grey
			pg.Grid[x][y].Alpha = 255
		}
	}
	return pg
}
