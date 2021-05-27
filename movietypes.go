package tmdb

type MovieDetails struct {
	Adult               bool                  `json:"adult"`
	BackdropPath        string                `json:"backdrop_path"`
	BelongsToCollection interface{}           `json:"belongs_to_collection"`
	Budget              int                   `json:"budget"`
	Genres              []Genres              `json:"genres"`
	Homepage            string                `json:"homepage"`
	ID                  int                   `json:"id"`
	ImdbID              string                `json:"imdb_id"`
	OriginalLanguage    string                `json:"original_language"`
	OriginalTitle       string                `json:"original_title"`
	Overview            string                `json:"overview"`
	Popularity          float64               `json:"popularity"`
	PosterPath          interface{}           `json:"poster_path"`
	ProductionCompanies []ProductionCompanies `json:"production_companies"`
	ProductionCountries []ProductionCountries `json:"production_countries"`
	ReleaseDate         string                `json:"release_date"`
	Revenue             int                   `json:"revenue"`
	Runtime             int                   `json:"runtime"`
	SpokenLanguages     []SpokenLanguages     `json:"spoken_languages"`
	Status              string                `json:"status"`
	Tagline             string                `json:"tagline"`
	Title               string                `json:"title"`
	Video               bool                  `json:"video"`
	VoteAverage         float64               `json:"vote_average"`
	VoteCount           int                   `json:"vote_count"`
}

// TODO: refactor structs below, get rid of duplication

type MoviesExternalIDs struct {
	ImdbId      string `json:"imdb_id"`
	FacebookId  string `json:"facebook_id"`
	InstagramId string `json:"instagram_id"`
	TwitterId   string `json:"twitter_id"`
	Id          int    `json:"id"`
}

type Results []struct {
	PosterPath       string  `json:"poster_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Title            string  `json:"title"`
	BackdropPath     string  `json:"backdrop_path"`
	Popularity       float64 `json:"popularity"`
	VoteCount        int     `json:"vote_count"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
}

type NowPlaying Upcoming

// Upcoming takes care of upcoming and now playing results
type Upcoming struct {
	Page    int `json:"page"`
	Results `json:"results"`
	Dates   struct {
		Maximum string `json:"maximum"`
		Minimum string `json:"minimum"`
	} `json:"dates"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type Popular PopularityRating
type TopRated PopularityRating

// PopularityRating takes care of TopRated and Popular results
type PopularityRating struct {
	Page         int `json:"page"`
	Results      `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}
