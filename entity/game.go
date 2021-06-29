package entity

type GameProperties struct {
	Name            string          `json:"name"`
	Platforms       []Named         `json:"platforms"`
	YearReleased    int             `json:"year_released"`
	ImageURL        string          `json:"image_url"`
	Genres          []Named         `json:"genres"`
	genreChecker    map[string]bool `json:"-"`
	platformChecker map[string]bool `json:"-"`
}

type Named struct {
	Name string `json:"name"`
}

var (
	platformReplacer map[string]string = map[string]string{
		"iPhone/iPad":   "iOS",
		"PlayStation":   "PS",
		"PlayStation 2": "PS 2",
		"PlayStation 3": "PS 3",
		"PlayStation 4": "PS 4",
		"PlayStation 5": "PS 5",
		"Switch":        "Nintendo Switch",
		"3DS":           "Nintento 3DS",
	}

	uselessGenres map[string]bool = map[string]bool{"3D": true, "General": true, "2D": true, "Console-style RPG": true,
		"Miscellaneous": true, "Compilation": true}
	genreReplacer map[string]string = map[string]string{
		"Role-Playing": "RPG",
	}
)

func (g *GameProperties) AddPlatform(platform string) {
	if g.platformChecker == nil {
		g.platformChecker = make(map[string]bool)
	}

	// Replace platform names
	if replace, ok := platformReplacer[platform]; ok {
		platform = replace
	}

	// Check for duplicates
	if !g.platformChecker[platform] {
		g.Platforms = append(g.Platforms, Named{platform})

		g.platformChecker[platform] = true
	}
}

func (g *GameProperties) AddGenre(genre string) {
	if g.genreChecker == nil {
		g.genreChecker = make(map[string]bool)
	}

	// Replace genre names
	if replace, ok := genreReplacer[genre]; ok {
		genre = replace
	}

	// Check for duplicates
	if !g.genreChecker[genre] && !uselessGenres[genre] {
		g.Genres = append(g.Genres, Named{genre})
		g.genreChecker[genre] = true
	}
}
