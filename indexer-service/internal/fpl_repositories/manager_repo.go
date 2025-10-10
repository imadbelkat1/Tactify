package fpl_repositories

import (
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/imadbelkat1/shared/models"
)

type ManagerRepo struct {
	db             *sql.DB
	Entry          *models.EntryMessage
	EntryPicks     *models.EntryEventPicksMessage
	EntryTransfers *models.EntryTransfersMessage
	EntryHistory   *models.EntryHistoryMessage
}

func NewManagerRepo(db *sql.DB, entry *models.EntryMessage, entryPicks *models.EntryEventPicksMessage, entryTransfers *models.EntryTransfersMessage, entryHistory *models.EntryHistoryMessage) *ManagerRepo {
	return &ManagerRepo{
		db:             db,
		Entry:          entry,
		EntryPicks:     entryPicks,
		EntryTransfers: entryTransfers,
		EntryHistory:   entryHistory,
	}
}

func (r *ManagerRepo) InsertManagerInfo(entry *models.EntryMessage) error {
	if entry == nil {
		return nil
	}

	query := sq.Insert("managers").Columns(
		"manager_id", "season_id", "manager_name", "player_first_name", "player_last_name", "player_region_id",
		"player_region_name", "player_region_iso_code_short", "player_region_iso_code_long", "favourite_team_id", "joined_time",
		"started_event", "years_active", "summary_overall_points", "summary_overall_rank", "summary_event_points", "summary_event_rank",
		"current_event", "name_change_blocked", "last_deadline_bank", "last_deadline_value", "last_deadline_total_transfers",
		"club_badge_src",
	).Values(
		entry.Entry.ID, entry.SeasonId, entry.Entry.Name, entry.Entry.PlayerFirstName, entry.Entry.PlayerLastName, entry.Entry.PlayerRegionID,
		entry.Entry.PlayerRegionName, entry.Entry.PlayerRegionShort, entry.Entry.PlayerRegionLong, entry.Entry.FavouriteTeam, entry.Entry.JoinedTime,
		entry.Entry.StartedEvent, entry.Entry.YearsActive, entry.Entry.SummaryOverallPoints, entry.Entry.SummaryOverallRank, entry.Entry.SummaryEventPoints, entry.Entry.SummaryEventRank,
		entry.Entry.CurrentEvent, entry.Entry.NameChangeBlocked, entry.Entry.LastDeadlineBank, entry.Entry.LastDeadlineValue, entry.Entry.LastDeadlineTransfers,
		entry.Entry.ClubBadgeSrc,
	).Suffix(`ON CONFLICT (manager_id, season_id) DO UPDATE SET 
		manager_name = EXCLUDED.manager_name,
		player_first_name = EXCLUDED.player_first_name,
		player_last_name = EXCLUDED.player_last_name,
		player_region_id = EXCLUDED.player_region_id,
		player_region_name = EXCLUDED.player_region_name,
		player_region_iso_code_short = EXCLUDED.player_region_iso_code_short,
		player_region_iso_code_long = EXCLUDED.player_region_iso_code_long,
		favourite_team_id = EXCLUDED.favourite_team_id,
		joined_time = EXCLUDED.joined_time,
		started_event = EXCLUDED.started_event,
		years_active = EXCLUDED.years_active,
		summary_overall_points = EXCLUDED.summary_overall_points,
		summary_overall_rank = EXCLUDED.summary_overall_rank,
		summary_event_points = EXCLUDED.summary_event_points,
		summary_event_rank = EXCLUDED.summary_event_rank,
		current_event = EXCLUDED.current_event,
		name_change_blocked = EXCLUDED.name_change_blocked,
		last_deadline_bank = EXCLUDED.last_deadline_bank,
		last_deadline_value = EXCLUDED.last_deadline_value,
		last_deadline_total_transfers = EXCLUDED.last_deadline_total_transfers,
		club_badge_src = EXCLUDED.club_badge_src,
		updated_at = CURRENT_TIMESTAMP`,
	).PlaceholderFormat(sq.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("building managers insert query: %w", err)
	}

	result, err := r.db.Exec(sqlQuery, args...)
	if err != nil {
		log.Printf("❌ SQL Error inserting players: %v", err)
		log.Printf("   Query length: %d characters", len(sqlQuery))
		log.Printf("   Args count: %d", len(args))
		return fmt.Errorf("executing players insert: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	log.Printf("✅ Players insert completed: %d rows affected", rowsAffected)
	return err
}
