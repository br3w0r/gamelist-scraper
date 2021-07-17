package entity

import (
	"math"

	pb "bitbucket.org/br3w0r/gamelist-scraper/proto"
)

type GameProperties struct {
	Name            string          `json:"name"`
	Platforms       []Named         `json:"platforms"`
	YearReleased    uint32          `json:"year_released"`
	ImageUrl        string          `json:"image_url"`
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

func (g *GameProperties) ConvertToProto() pb.GameProperties {
	nPlatforms := len(g.Platforms)
	nGenres := len(g.Genres)

	platforms := make([]*pb.Named, nPlatforms)
	genres := make([]*pb.Named, nGenres)

	maxIter := int(math.Max(float64(nPlatforms), float64(nGenres)))
	for i := 0; i < maxIter; i++ {
		if i < nPlatforms {
			platforms[i] = &pb.Named{Name: g.Platforms[i].Name}
		}
		if i < nGenres {
			genres[i] = &pb.Named{Name: g.Genres[i].Name}
		}
	}

	return pb.GameProperties{
		Name:         g.Name,
		Platforms:    platforms,
		YearReleased: g.YearReleased,
		ImageUrl:     g.ImageUrl,
		Genres:       genres,
	}
}
