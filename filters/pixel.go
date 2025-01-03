package filters

type PixelGrid struct {
	MaxX int
	MaxY int
	Grid [][]*Pixel
}

type Pixel struct {
	Red       uint8
	Green     uint8
	Blue      uint8
	Alpha     uint8
	Lightness float64
	Zone      int
}
