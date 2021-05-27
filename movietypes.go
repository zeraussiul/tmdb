package tmdb

import "time"

// MovieDetails holds all details of a movie
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

// MoviesExternalIDs holds external ids where movie data may be found, outside of TMDB
type MoviesExternalIDs struct {
	ImdbId      string `json:"imdb_id"`
	FacebookId  string `json:"facebook_id"`
	InstagramId string `json:"instagram_id"`
	TwitterId   string `json:"twitter_id"`
	Id          int    `json:"id"`
}

// Results is a slice of results for results when returning more than one movie
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

// NowPlaying holds currently playing movies
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

// Popular holds current popular movies
type Popular PopularityRating

// TopRated holds current top rated movies
type TopRated PopularityRating

// SimilarMovies holds movies similar to requested movie
type SimilarMovies PopularityRating

// RecommendedMovies holds recommended movies for a movie
type RecommendedMovies PopularityRating

// PopularityRating takes care of TopRated and Popular results
type PopularityRating struct {
	Page         int `json:"page"`
	Results      `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// ReleaseDates holds the release date along with the certification for a movie
//
// Types supported:
//
// 1. Premiere
// 2. Theatrical (limited)
// 3. Theatrical
// 4. Digital
// 5. Physical
// 6. TV
type ReleaseDates struct {
	Id      int `json:"id"`
	Results []struct {
		Iso31661     string `json:"iso_3166_1"`
		ReleaseDates []struct {
			Certification string    `json:"certification"`
			Iso6391       string    `json:"iso_639_1"`
			ReleaseDate   time.Time `json:"release_date"`
			Type          int       `json:"type"`
			Note          string    `json:"note,omitempty"`
		} `json:"release_dates"`
	} `json:"results"`
}
