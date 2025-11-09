package sofascore_repositories

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/imadbelkat1/shared/sofascore_models"
)

type LeagueRepo struct {
	db     *sql.DB
	League *sofascore_models.LeagueUniqueTournaments
}

func NewLeagueRepo(
	db *sql.DB,
	League *sofascore_models.LeagueUniqueTournaments,
) *LeagueRepo {
	return &LeagueRepo{
		db:     db,
		League: League,
	}
}

func (l *LeagueRepo) InsertLeagueInfo(uniqueTournament sofascore_models.LeagueUniqueTournaments) error {
	query := sq.Insert("leagues").Columns(
		"league_id", "name", "country").Suffix(
		"ON CONFLICT (league_id) DO UPDATE SET " +
			"name=EXCLUDED.name, " +
			"country=EXCLUDED.country, " +
			"updated_at= CURRENT_TIMESTAMP",
	).PlaceholderFormat(sq.Dollar)

	for _, group := range uniqueTournament.Groups {
		for _, league := range group.UniqueTournament {
			query = query.Values(league.ID, league.Name, league.Category.Name)
		}
	}

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = l.db.Exec(sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}
