package sofascore_models

type LeagueCategories struct {
	Categories []LeagueCategory `json:"categories"`
}
type LeagueCategory struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Flag string `json:"flag"`
}

type LeagueUniqueTournaments struct {
	Groups []UniqueTournamentGroups `json:"groups"`
}

type UniqueTournamentGroups struct {
	UniqueTournament []UniqueTournament `json:"uniqueTournaments"`
}
