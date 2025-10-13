package sofascore_repositories

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/imadbelkat1/shared/sofascore_models"
)

type TeamRepo struct {
	db             *sql.DB
	LeagueStanding *sofascore_models.StandingMessage
}

func NewTeamRepo(db *sql.DB, leagueStanding *sofascore_models.StandingMessage) *TeamRepo {
	return &TeamRepo{
		db:             db,
		LeagueStanding: leagueStanding,
	}
}

func (r *TeamRepo) InsertTeamInfo(standing sofascore_models.StandingMessage) error {
	query := sq.Insert("teams").Columns(
		"team_id", "name", "primary_color", "secondary_color",
		"league_id", "season_id",
	).Suffix(
		"ON CONFLICT (team_id, league_id, season_id) DO UPDATE SET " +
			"name = EXCLUDED.name, " +
			"primary_color = EXCLUDED.primary_color, " +
			"secondary_color = EXCLUDED.secondary_color," +
			"updated_at = CURRENT_TIMESTAMP",
	).PlaceholderFormat(sq.Dollar)

	query = query.Values(
		standing.Row.Team.ID,
		standing.Row.Team.Name,
		standing.Row.Team.Colors.PrimaryColor,
		standing.Row.Team.Colors.SecondaryColor,
		standing.LeagueID,
		standing.SeasonID,
	)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}
