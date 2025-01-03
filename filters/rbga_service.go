package filters

import (
	"image"
	"image/color"
)

func CopyPixels(img *image.RGBA) PixelGrid {
	rectangle := img.Bounds()

	pg := make([][]*Pixel, rectangle.Max.X)
	for x := 0; x < rectangle.Max.X; x++ {
		pg[x] = make([]*Pixel, rectangle.Max.Y)
		for y := 0; y < rectangle.Max.Y; y++ {

			i := img.PixOffset(x, y)
			s := img.Pix[i : i+4 : i+4]

			pg[x][y] = &Pixel{
				Red:       s[0],
				Green:     s[1],
				Blue:      s[2],
				Alpha:     s[3],
				Lightness: getLightness(s[0], s[1], s[2]),
			}

		}
	}

	return PixelGrid{
		MaxX: rectangle.Max.X,
		MaxY: rectangle.Max.Y,
		Grid: pg,
	}
}

func NewImageRGBA(pg PixelGrid) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, pg.MaxX, pg.MaxY))

	for x := 0; x < pg.MaxX; x++ {
		for y := 0; y < pg.MaxY; y++ {
			img.SetRGBA(
				x,
				y,
				color.RGBA{
					R: pg.Grid[x][y].Red,
					G: pg.Grid[x][y].Green,
					B: pg.Grid[x][y].Blue,
					A: pg.Grid[x][y].Alpha,
				})
		}
	}

	return img
}

// TODO Check if this is correct
func getLightness(r, g, b uint8) float64 {
	rf := float64(r) / 255.0
	gf := float64(g) / 255.0
	bf := float64(b) / 255.0

	max := rf
	if gf > max {
		max = gf
	}
	if bf > max {
		max = bf
	}

	min := rf
	if gf < min {
		min = gf
	}
	if bf < min {
		min = bf
	}

	return (max + min) / 2.0
}
