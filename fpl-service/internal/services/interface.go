package services

import (
	model "github.com/imadbelkat1/fpl-service/internal/models"
)

// Services aggregates all service interfaces for FPL data
type Services struct {
	ElementType       ElementTypeService
	Entry             EntryService
	EntryEvent        EntryEventService
	EntryHistory      EntryHistoryService
	Fixture           FixtureService
	LiveEvent         LiveEventService
	Player            PlayerService
	PlayerHistory     PlayerHistoryService
	PlayerPastHistory PlayerPastHistoryService
	Scoring           ScoringService
	Team              TeamService
	Gameweek          GameweekService
}

// TeamService handles Premier League team data
type TeamService interface {
	UpdateTeams(teams []model.Team) error
	GetTeam(id int) (*model.Team, error)
	GetAllTeams() ([]model.Team, error)
	GetTeamsByStrength(minStrength int) ([]model.Team, error)
	GetLeagueTable() ([]model.Team, error)
	GetTeamFixtures(teamID int, upcoming bool) ([]model.Fixture, error)
	GetTeamForm(teamID int, gameweeks int) ([]model.Team, error)
}

// PlayerService handles player data and statistics
type PlayerService interface {
	GetPlayer(id int) (*model.Player, error)
	GetAllPlayers() ([]model.Player, error)
	GetPlayersByTeam(teamID int) ([]model.Player, error)
	GetPlayersByPosition(elementType int) ([]model.Player, error)
	GetPlayersByPriceRange(minPrice, maxPrice int) ([]model.Player, error)
	GetAvailablePlayers() ([]model.Player, error)
	GetPlayersByOwnership(minPercent float64) ([]model.Player, error)
	GetTopScoringPlayers(position int, gameweeks int) ([]model.Player, error)
	SearchPlayersByName(name string) ([]model.Player, error)
	GetPlayersByForm(minForm float64) ([]model.Player, error)
}

// FixtureService handles match fixtures and results
type FixtureService interface {
	UpdateFixtures(fixtures []model.Fixture) error
	GetFixture(id int) (*model.Fixture, error)
	GetFixturesByGameweek(gameweek int) ([]model.Fixture, error)
	GetFixturesByTeam(teamID int) ([]model.Fixture, error)
	GetUpcomingFixtures(limit int) ([]model.Fixture, error)
	GetCompletedFixtures(gameweek int) ([]model.Fixture, error)
	GetCurrentGameweekFixtures() ([]model.Fixture, error)
}

// LiveEventService handles live gameweek data
type LiveEventService interface {
	GetLiveEvent(gameweek int) (*model.LiveEvent, error)
	GetLivePlayerStats(playerID, gameweek int) (*model.LiveElement, error)
}

// PlayerHistoryService handles player gameweek performance
type PlayerHistoryService interface {
	GetPlayerHistory(playerID int) ([]model.PlayerHistory, error)
	GetPlayerFixtureHistory(playerID, fixtureID int) (*model.PlayerHistory, error)
	GetPlayerGameweekHistory(playerID, gameweek int) ([]model.PlayerHistory, error)
	GetPlayerHomeAwayStats(playerID int) ([]model.PlayerHistory, []model.PlayerHistory, error)
	GetPlayerRecentForm(playerID int, gameweeks int) ([]model.PlayerHistory, error)
}

// PlayerPastHistoryService handles player season history
type PlayerPastHistoryService interface {
	GetPlayerPastHistory(playerID int) ([]model.PlayerPastHistory, error)
	GetPlayerSeasonHistory(playerID int, season string) (*model.PlayerPastHistory, error)
	GetPlayerCareerStats(playerID int) ([]model.PlayerPastHistory, error)
}

// ScoringService handles FPL scoring rules and calculations
type ScoringService interface {
	GetScoringRules() (*model.Scoring, error)
	CalculatePlayerGameweekPoints(playerHistory *model.PlayerHistory) (int, error)
	CalculateBonusPoints(fixtureID int) (map[int]int, error)
}

// GameweekService handles current gameweek information
type GameweekService interface {
	GetCurrentGameweek() (int, error)
	GetGameweekDeadline(gameweek int) (string, error)
	GetGameweekStatus(gameweek int) (string, error)
	IsGameweekLive(gameweek int) (bool, error)
}

// ElementTypeService handles operations for ElementType (positions: GKP, DEF, MID, FWD)
type ElementTypeService interface {
	GetElementType(id int) (*model.ElementType, error)
	GetAllElementTypes() ([]model.ElementType, error)
}

// EntryService handles operations for manager entries/teams
type EntryService interface {
	GetEntry(id int) (*model.Entry, error)
	GetEntriesByLeague(leagueID int) ([]model.Entry, error)
	UpdateEntry(entry *model.Entry) error
	GetEntryRankings(limit int) ([]model.Entry, error)
}

// EntryEventService handles gameweek-specific team data
type EntryEventService interface {
	GetEntryEvent(entryID, eventID int) (*model.EntryEvent, error)
	GetEntryPicks(entryID, eventID int) ([]model.Pick, error)
	GetAutomaticSubs(entryID, eventID int) ([]model.AutomaticSub, error)
}

// EntryHistoryService handles manager historical performance
type EntryHistoryService interface {
	GetEntryHistory(entryID int) (*model.EntryHistory, error)
	GetEntryCurrentHistory(entryID int) ([]model.EntryHistoryCurrent, error)
	GetEntryPastHistory(entryID int) ([]model.EntryHistoryPast, error)
	GetEntryChips(entryID int) ([]model.EntryHistoryChip, error)
}
