package entity

type GameProperties struct {
	Name         string  `json:"name"`
	Platforms    []Named `json:"platforms"`
	YearReleased string  `json:"year_released"`
	ImageURL     string  `json:"image_url"`
}

type Named struct {
	Name string `json:"name"`
}

func (g *GameProperties) AddPlatform(platform string) {
	g.Platforms = append(g.Platforms, Named{platform})
}
