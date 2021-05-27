package tmdb

// Response is the JSON respnse by the server
// https://www.themoviedb.org/documentation/api/status-codes
type Response struct {
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message,omitempty"`
}

// Structs here are reused in different media types, ie TvClient and Movies.

type Genres struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductionCompanies struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type ProductionCountries struct {
	Iso31661 string `json:"iso_3166_1"`
	Name     string `json:"name"`
}

type SpokenLanguages struct {
	EnglishName string `json:"english_name"`
	Iso6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
}

type Videos struct {
	Id      int `json:"id"`
	Results []struct {
		Id       string `json:"id"`
		Iso6391  string `json:"iso_639_1"`
		Iso31661 string `json:"iso_3166_1"`
		Key      string `json:"key"`
		Name     string `json:"name"`
		Site     string `json:"site"`
		Size     int    `json:"size"`
		Type     string `json:"type"`
	} `json:"results"`
}

type AlternativeTitles struct {
	Id     int `json:"id"`
	Titles []struct {
		Iso31661 string `json:"iso_3166_1"`
		Title    string `json:"title"`
		Type     string `json:"type,omitempty"`
	} `json:"titles"`
}
