package filters

func retroGame(pg PixelGrid) PixelGrid {
	for x := 0; x < pg.MaxX; x++ {
		// chessPatternFlag is used to alternate the color of the chess pattern
		chessPatternFlag := x%2 == 0
		for y := 0; y < pg.MaxY; y++ {
			if chessPatternFlag || pg.Grid[x][y].Lightness <= 0.25 {
				// #0f380f
				pg.Grid[x][y].Red = 15
				pg.Grid[x][y].Green = 56
				pg.Grid[x][y].Blue = 15
				pg.Grid[x][y].Alpha = 255
			} else if pg.Grid[x][y].Lightness <= 0.5 {
				// #306230
				pg.Grid[x][y].Red = 48
				pg.Grid[x][y].Green = 98
				pg.Grid[x][y].Blue = 48
				pg.Grid[x][y].Alpha = 255
			} else if pg.Grid[x][y].Lightness <= 0.75 {
				// #8bac0f
				pg.Grid[x][y].Red = 139
				pg.Grid[x][y].Green = 172
				pg.Grid[x][y].Blue = 15
				pg.Grid[x][y].Alpha = 255
			} else {
				// #9bbc0f
				pg.Grid[x][y].Red = 155
				pg.Grid[x][y].Green = 188
				pg.Grid[x][y].Blue = 15
				pg.Grid[x][y].Alpha = 255
			}

			chessPatternFlag = !chessPatternFlag
		}
	}
	return pg
}
