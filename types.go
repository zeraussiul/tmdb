package tmdb

// Response is the JSON respnse by the server
// https://www.themoviedb.org/documentation/api/status-codes
type Response struct {
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message,omitempty"`
}

// Genres identifies the genres of media being requested
type Genres struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ProductionCompanies is a struct containing a production company's detail
type ProductionCompanies struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

// ProductionCountries is a struct containing countries where production took place
type ProductionCountries struct {
	Iso31661 string `json:"iso_3166_1"`
	Name     string `json:"name"`
}

// SpokenLanguages identifies spoken languages in media requested
type SpokenLanguages struct {
	EnglishName string `json:"english_name"`
	Iso6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
}

// Videos struct holds video information and url where to find video, may be a trailer, review, etc
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

// AlternativeTitles is any alternative titles media requested may have
type AlternativeTitles struct {
	Id     int `json:"id"`
	Titles []struct {
		Iso31661 string `json:"iso_3166_1"`
		Title    string `json:"title"`
		Type     string `json:"type,omitempty"`
	} `json:"titles"`
}
