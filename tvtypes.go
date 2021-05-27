package tmdb

// TVDetails holds all details on a TV show
type TVDetails struct {
	BackdropPath        string                `json:"backdrop_path"`
	CreatedBy           []CreatedBy           `json:"created_by"`
	EpisodeRunTime      []int                 `json:"episode_run_time"`
	FirstAirDate        string                `json:"first_air_date"`
	Genres              []Genres              `json:"genres"`
	Homepage            string                `json:"homepage"`
	ID                  int                   `json:"id"`
	InProduction        bool                  `json:"in_production"`
	Languages           []string              `json:"languages"`
	LastAirDate         string                `json:"last_air_date"`
	LastEpisodeToAir    EpisodeToAir          `json:"last_episode_to_air"`
	Name                string                `json:"name"`
	NextEpisodeToAir    EpisodeToAir          `json:"next_episode_to_air"`
	Networks            []Networks            `json:"networks"`
	NumberOfEpisodes    int                   `json:"number_of_episodes"`
	NumberOfSeasons     int                   `json:"number_of_seasons"`
	OriginCountry       []string              `json:"origin_country"`
	OriginalLanguage    string                `json:"original_language"`
	OriginalName        string                `json:"original_name"`
	Overview            string                `json:"overview"`
	Popularity          float64               `json:"popularity"`
	PosterPath          string                `json:"poster_path"`
	ProductionCompanies []ProductionCompanies `json:"production_companies"`
	ProductionCountries []ProductionCountries `json:"production_countries"`
	Seasons             []Seasons             `json:"seasons"`
	SpokenLanguages     []SpokenLanguages     `json:"spoken_languages"`
	Status              string                `json:"status"`
	Tagline             string                `json:"tagline"`
	Type                string                `json:"type"`
	VoteAverage         float64               `json:"vote_average"`
	VoteCount           int                   `json:"vote_count"`
}

// CreatedBy credits show's creator
type CreatedBy struct {
	ID          int    `json:"id"`
	CreditID    string `json:"credit_id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profile_path"`
}

// EpisodeToAir holds episode information and air date
type EpisodeToAir struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int     `json:"season_number"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}

// Networks holds network information where show may be aired
type Networks struct {
	Name          string `json:"name"`
	ID            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

// Seasons holds season information about a TV Show
type Seasons struct {
	AirDate      string `json:"air_date"`
	EpisodeCount int    `json:"episode_count"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
}
