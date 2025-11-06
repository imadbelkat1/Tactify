package sofascore_models

type PlayerMatchStatsMessage struct {
	PlayerName  string      `json:"player_name"`
	SeasonID    int         `json:"season"`
	LeagueID    int         `json:"league"`
	MatchID     int         `json:"match"`
	Round       int         `json:"round"`
	MatchPlayer MatchPlayer `json:"player"`
}
type MatchLineupMessage struct {
	SeasonID    int         `json:"season"`
	LeagueID    int         `json:"league"`
	MatchID     int         `json:"match"`
	Round       int         `json:"round"`
	MatchLineup MatchLineup `json:"lineup"`
}
type MatchLineup struct {
	Confirmed bool       `json:"confirmed"`
	Home      TeamLineup `json:"home"`
	Away      TeamLineup `json:"away"`
}
type PlayerStatistics struct {
	MinutesPlayed                int     `json:"minutesPlayed"`
	Rating                       float64 `json:"rating"`
	Touches                      int     `json:"touches"`
	TotalPass                    int     `json:"totalPass"`
	AccuratePass                 int     `json:"accuratePass"`
	TotalLongBalls               int     `json:"totalLongBalls,omitempty"`
	AccurateLongBalls            int     `json:"accurateLongBalls,omitempty"`
	AccurateOwnHalfPasses        int     `json:"accurateOwnHalfPasses,omitempty"`
	TotalOwnHalfPasses           int     `json:"totalOwnHalfPasses,omitempty"`
	AccurateOppositionHalfPasses int     `json:"accurateOppositionHalfPasses,omitempty"`
	TotalOppositionHalfPasses    int     `json:"totalOppositionHalfPasses,omitempty"`
	TotalCross                   int     `json:"totalCross,omitempty"`
	AccurateCross                int     `json:"accurateCross,omitempty"`
	KeyPass                      int     `json:"keyPass,omitempty"`
	GoalAssist                   int     `json:"goalAssist,omitempty"`
	BigChanceCreated             int     `json:"bigChanceCreated,omitempty"`
	TotalShots                   int     `json:"totalShots"`
	ShotOffTarget                int     `json:"shotOffTarget,omitempty"`
	OnTargetScoringAttempt       int     `json:"onTargetScoringAttempt,omitempty"`
	BlockedScoringAttempt        int     `json:"blockedScoringAttempt,omitempty"`
	HitWoodwork                  int     `json:"hitWoodwork,omitempty"`
	Goals                        int     `json:"goals,omitempty"`
	BigChanceMissed              int     `json:"bigChanceMissed,omitempty"`
	ExpectedGoals                float64 `json:"expectedGoals,omitempty"`
	ExpectedGoalsOnTarget        float64 `json:"expectedGoalsOnTarget,omitempty"`
	ExpectedAssists              float64 `json:"expectedAssists,omitempty"`
	TotalTackle                  int     `json:"totalTackle,omitempty"`
	WonTackle                    int     `json:"wonTackle,omitempty"`
	TotalClearance               int     `json:"totalClearance,omitempty"`
	InterceptionWon              int     `json:"interceptionWon,omitempty"`
	OutfielderBlock              int     `json:"outfielderBlock,omitempty"`
	BallRecovery                 int     `json:"ballRecovery,omitempty"`
	DuelLost                     int     `json:"duelLost,omitempty"`
	DuelWon                      int     `json:"duelWon,omitempty"`
	AerialLost                   int     `json:"aerialLost,omitempty"`
	AerialWon                    int     `json:"aerialWon,omitempty"`
	TotalContest                 int     `json:"totalContest,omitempty"`
	WonContest                   int     `json:"wonContest,omitempty"`
	ChallengeLost                int     `json:"challengeLost,omitempty"`
	PossessionLostCtrl           int     `json:"possessionLostCtrl,omitempty"`
	UnsuccessfulTouch            int     `json:"unsuccessfulTouch,omitempty"`
	Dispossessed                 int     `json:"dispossessed,omitempty"`
	Fouls                        int     `json:"fouls,omitempty"`
	WasFouled                    int     `json:"wasFouled,omitempty"`
	TotalOffside                 int     `json:"totalOffside,omitempty"`
	Saves                        int     `json:"saves,omitempty"`
	TotalKeeperSweeper           int     `json:"totalKeeperSweeper,omitempty"`
	AccurateKeeperSweeper        int     `json:"accurateKeeperSweeper,omitempty"`
	KeeperSaveValue              float64 `json:"keeperSaveValue,omitempty"`
	GoalsPrevented               float64 `json:"goalsPrevented,omitempty"`
	ErrorLeadToAShot             int     `json:"errorLeadToAShot,omitempty"`
	ErrorLeadToAGoal             int     `json:"errorLeadToAGoal,omitempty"`
	ShotValueNormalized          float64 `json:"shotValueNormalized,omitempty"`
	PassValueNormalized          float64 `json:"passValueNormalized,omitempty"`
	DribbleValueNormalized       float64 `json:"dribbleValueNormalized,omitempty"`
	DefensiveValueNormalized     float64 `json:"defensiveValueNormalized,omitempty"`
	GoalkeeperValueNormalized    float64 `json:"goalkeeperValueNormalized,omitempty"`
	RatingVersions               *struct {
		Original    float64 `json:"original"`
		Alternative float64 `json:"alternative"`
	} `json:"ratingVersions,omitempty"`
}

type LineupPlayer struct {
	ID                     int     `json:"id"`
	Name                   string  `json:"name"`
	FirstName              string  `json:"firstName,omitempty"`
	LastName               string  `json:"lastName,omitempty"`
	Slug                   string  `json:"slug"`
	ShortName              string  `json:"shortName"`
	Position               string  `json:"position"`
	JerseyNumber           string  `json:"jerseyNumber"`
	Height                 int     `json:"height"`
	UserCount              int     `json:"userCount"`
	Gender                 string  `json:"gender"`
	SofascoreID            string  `json:"sofascoreId,omitempty"`
	Country                Country `json:"country"`
	MarketValueCurrency    string  `json:"marketValueCurrency,omitempty"`
	DateOfBirthTimestamp   int64   `json:"dateOfBirthTimestamp"`
	ProposedMarketValueRaw *struct {
		Value    int    `json:"value"`
		Currency string `json:"currency"`
	} `json:"proposedMarketValueRaw,omitempty"`
	FieldTranslations *struct {
		NameTranslation      map[string]string `json:"nameTranslation,omitempty"`
		ShortNameTranslation map[string]string `json:"shortNameTranslation,omitempty"`
	} `json:"fieldTranslations,omitempty"`
}

type MatchPlayer struct {
	Player       Player           `json:"player"`
	TeamID       int              `json:"teamId"`
	ShirtNumber  int              `json:"shirtNumber"`
	JerseyNumber string           `json:"jerseyNumber"`
	Position     string           `json:"position"`
	Substitute   bool             `json:"substitute"`
	Captain      bool             `json:"captain,omitempty"`
	Statistics   PlayerStatistics `json:"statistics"`
}

type MissingPlayer struct {
	Player          Player `json:"player"`
	Type            string `json:"type"`
	Reason          int    `json:"reason"`
	Description     string `json:"description,omitempty"`
	ExternalType    int    `json:"externalType"`
	ExpectedEndDate string `json:"expectedEndDate,omitempty"`
}

type PlayerColor struct {
	Primary     string `json:"primary"`
	Number      string `json:"number"`
	Outline     string `json:"outline"`
	FancyNumber string `json:"fancyNumber"`
}

type TeamLineup struct {
	Players         []MatchPlayer   `json:"players"`
	SupportStaff    []interface{}   `json:"supportStaff"`
	Formation       string          `json:"formation"`
	PlayerColor     PlayerColor     `json:"playerColor"`
	GoalkeeperColor PlayerColor     `json:"goalkeeperColor"`
	MissingPlayers  []MissingPlayer `json:"missingPlayers,omitempty"`
}
