# tmdb client

[![Go Reference](https://pkg.go.dev/badge/github.com/zeraussiul/tmdb.svg)](https://pkg.go.dev/github.com/zeraussiul/tmdb)

- api wrapper for [tmdb](https://developers.themoviedb.org/3) api
- used in a personal project to retrieve movies and shows info from TMDB

### Usage:
```go
// create a client
client, err := tmdb.NewClient(API_KEY)


// get a tv show details
tv := tmdb.TV(client)

tvshow, err := tv.GetDetails(71712)


// get a movie detail is similar:
mov := tmdb.Movie(client)

m, err := mov.GetDetails(550)


// search for movies or show
s := tmdb.Search(client)

movs, err := s.Movies("fight club")

shws, err := s.TV("game of thrones")

```

### Installation
``` TODO ```

## Endpoints:
* [ ] Movies (Required Params):
  * [x] Get Details
  * [ ] ~~Get Account States~~
  * [x] Get Alternative Titles
  * [ ] Get Credits
  * [ ] Get Changes
  * [x] Get External IDs
  * [ ] Get Images
  * [ ] Get Keywords
  * [ ] Get Lists
  * [x] Get Recommendations
  * [x] Get Release Dates
  * [ ] Get Reviews
  * [x] Get Similar Movies
  * [ ] Get Translations
  * [x] Get Videos
  * [ ] Get Watch Providers
  * [x] Get Now Playing
  * [x] Get Popular
  * [x] Get Top Rated
  * [x] Get Upcoming
  
* [ ] TV Shows:
  * [x] Get Details
  
* [ ] Search
  * [x] Search Movies
  * [x] Search TV Shows

## TODO:

* [ ] Proper tests
* [ ] Optional parameters on all endpoints as needed

## License
Distributed under the MIT License. See ```LICENSE``` for more information.