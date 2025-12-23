-- Sofascore Service Database Schema
\connect sofascore;
-- Leagues
CREATE TABLE IF NOT EXISTS leagues (
                                       league_id INTEGER PRIMARY KEY,
                                       name VARCHAR(255),
                                       country VARCHAR(255),
                                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Seasons
CREATE TABLE IF NOT EXISTS seasons (
                                       season_id INTEGER PRIMARY KEY,
                                       league_id INTEGER,
                                       name VARCHAR(255),
                                       year VARCHAR(20),
                                       is_current BOOLEAN,
                                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Matches
CREATE TABLE IF NOT EXISTS matches (
                                       match_id INTEGER NOT NULL,
                                       season_id INTEGER NOT NULL,
                                       league_id INTEGER NOT NULL,
                                       home_team_id INTEGER,
                                       away_team_id INTEGER,
                                       home_team_name VARCHAR(255),
                                       away_team_name VARCHAR(255),
                                       start_time TIMESTAMP,
                                       round VARCHAR(50), -- or INTEGER depending on data
                                       status VARCHAR(50),
                                       status_description VARCHAR(255),
                                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       PRIMARY KEY (match_id, season_id, league_id)
);

-- Teams (Sofascore)
CREATE TABLE IF NOT EXISTS teams (
                                     team_id INTEGER NOT NULL,
                                     league_id INTEGER NOT NULL,
                                     name VARCHAR(255),
                                     primary_color VARCHAR(50),
                                     secondary_color VARCHAR(50),
                                     PRIMARY KEY (team_id, league_id)
);

-- Team Overall Stats
CREATE TABLE IF NOT EXISTS team_overall_stats (
    -- Primary Keys
                                                  team_id INTEGER NOT NULL,
                                                  league_id INTEGER NOT NULL,
                                                  season_id INTEGER NOT NULL,

    -- Offensive Stats
                                                  goals_scored INTEGER,
                                                  goals_conceded INTEGER,
                                                  own_goals INTEGER,
                                                  assists INTEGER,
                                                  shots INTEGER,
                                                  shots_on_target INTEGER,
                                                  shots_off_target INTEGER,
                                                  penalty_goals INTEGER,
                                                  penalties_taken INTEGER,
                                                  free_kick_goals INTEGER,
                                                  free_kick_shots INTEGER,

    -- Positional Goals
                                                  goals_from_inside_box INTEGER,
                                                  goals_from_outside_box INTEGER,
                                                  headed_goals INTEGER,
                                                  left_foot_goals INTEGER,
                                                  right_foot_goals INTEGER,

    -- Chances
                                                  big_chances INTEGER,
                                                  big_chances_created INTEGER,
                                                  big_chances_missed INTEGER,

    -- Shooting Details
                                                  shots_from_inside_box INTEGER,
                                                  shots_from_outside_box INTEGER,
                                                  blocked_scoring_attempt INTEGER,
                                                  hit_woodwork INTEGER,

    -- Dribbling
                                                  successful_dribbles INTEGER,
                                                  dribble_attempts INTEGER,

    -- Set Pieces
                                                  corners INTEGER,
                                                  free_kicks INTEGER,
                                                  throw_ins INTEGER,
                                                  goal_kicks INTEGER,

    -- Fast Breaks
                                                  fast_breaks INTEGER,
                                                  fast_break_goals INTEGER,
                                                  fast_break_shots INTEGER,

    -- Possession & Passing
                                                  average_ball_possession DECIMAL,
                                                  total_passes INTEGER,
                                                  accurate_passes INTEGER,
                                                  accurate_passes_percentage DECIMAL,

    -- Passing by Zone
                                                  total_own_half_passes INTEGER,
                                                  accurate_own_half_passes INTEGER,
                                                  accurate_own_half_passes_percentage DECIMAL,
                                                  total_opposition_half_passes INTEGER,
                                                  accurate_opposition_half_passes INTEGER,
                                                  accurate_opposition_half_passes_percentage DECIMAL,

    -- Long Balls & Crosses
                                                  total_long_balls INTEGER,
                                                  accurate_long_balls INTEGER,
                                                  accurate_long_balls_percentage DECIMAL,
                                                  total_crosses INTEGER,
                                                  accurate_crosses INTEGER,
                                                  accurate_crosses_percentage DECIMAL,

    -- Defensive Stats
                                                  clean_sheets INTEGER,
                                                  tackles INTEGER,
                                                  interceptions INTEGER,
                                                  saves INTEGER,
                                                  clearances INTEGER,
                                                  clearances_off_line INTEGER,
                                                  last_man_tackles INTEGER,
                                                  ball_recovery INTEGER,

    -- Errors
                                                  errors_leading_to_goal INTEGER,
                                                  errors_leading_to_shot INTEGER,

    -- Penalties
                                                  penalties_commited INTEGER,
                                                  penalty_goals_conceded INTEGER,

    -- Duels
                                                  total_duels INTEGER,
                                                  duels_won INTEGER,
                                                  duels_won_percentage DECIMAL,
                                                  total_ground_duels INTEGER,
                                                  ground_duels_won INTEGER,
                                                  ground_duels_won_percentage DECIMAL,
                                                  total_aerial_duels INTEGER,
                                                  aerial_duels_won INTEGER,
                                                  aerial_duels_won_percentage DECIMAL,

    -- Discipline
                                                  possession_lost INTEGER,
                                                  offsides INTEGER,
                                                  fouls INTEGER,
                                                  yellow_cards INTEGER,
                                                  yellow_red_cards INTEGER,
                                                  red_cards INTEGER,

    -- Performance
                                                  avg_rating DECIMAL,
                                                  matches INTEGER,
                                                  awarded_matches INTEGER,

    -- Stats Against (Defensive Perspective)
                                                  accurate_final_third_passes_against INTEGER,
                                                  accurate_opposition_half_passes_against INTEGER,
                                                  accurate_own_half_passes_against INTEGER,
                                                  accurate_passes_against INTEGER,
                                                  big_chances_against INTEGER,
                                                  big_chances_created_against INTEGER,
                                                  big_chances_missed_against INTEGER,
                                                  clearances_against INTEGER,
                                                  corners_against INTEGER,
                                                  crosses_successful_against INTEGER,
                                                  crosses_total_against INTEGER,
                                                  dribble_attempts_total_against INTEGER,
                                                  dribble_attempts_won_against INTEGER,
                                                  errors_leading_to_goal_against INTEGER,
                                                  errors_leading_to_shot_against INTEGER,
                                                  hit_woodwork_against INTEGER,
                                                  interceptions_against INTEGER,
                                                  key_passes_against INTEGER,
                                                  long_balls_successful_against INTEGER,
                                                  long_balls_total_against INTEGER,
                                                  offsides_against INTEGER,
                                                  red_cards_against INTEGER,
                                                  shots_against INTEGER,
                                                  shots_blocked_against INTEGER,
                                                  shots_from_inside_box_against INTEGER,
                                                  shots_from_outside_box_against INTEGER,
                                                  shots_off_target_against INTEGER,
                                                  shots_on_target_against INTEGER,
                                                  blocked_scoring_attempt_against INTEGER,
                                                  tackles_against INTEGER,
                                                  total_final_third_passes_against INTEGER,
                                                  opposition_half_passes_total_against INTEGER,
                                                  own_half_passes_total_against INTEGER,
                                                  total_passes_against INTEGER,
                                                  yellow_cards_against INTEGER,

                                                  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                  PRIMARY KEY (team_id, league_id, season_id)
);

-- Match Stats Tables (Columns inferred from base; specific stats are dynamic in code)

CREATE TABLE IF NOT EXISTS match_overview (
                                              match_id INTEGER NOT NULL,
                                              period VARCHAR(20) NOT NULL,
                                              home_team_id INTEGER,
                                              away_team_id INTEGER,
                                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Add dynamic stats columns here (e.g. possession, corners)
                                              PRIMARY KEY (match_id, period)
);

CREATE TABLE IF NOT EXISTS match_shots (
                                           match_id INTEGER NOT NULL,
                                           period VARCHAR(20) NOT NULL,
                                           home_team_id INTEGER,
                                           away_team_id INTEGER,
                                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Add dynamic stats columns here (e.g. total_shots, on_target)
                                           PRIMARY KEY (match_id, period)
);

CREATE TABLE IF NOT EXISTS match_attack (
                                            match_id INTEGER NOT NULL,
                                            period VARCHAR(20) NOT NULL,
                                            home_team_id INTEGER,
                                            away_team_id INTEGER,
                                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Add dynamic stats columns here
                                            PRIMARY KEY (match_id, period)
);

CREATE TABLE IF NOT EXISTS match_passes (
                                            match_id INTEGER NOT NULL,
                                            period VARCHAR(20) NOT NULL,
                                            home_team_id INTEGER,
                                            away_team_id INTEGER,
                                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Add dynamic stats columns here
                                            PRIMARY KEY (match_id, period)
);

CREATE TABLE IF NOT EXISTS match_duels (
                                           match_id INTEGER NOT NULL,
                                           period VARCHAR(20) NOT NULL,
                                           home_team_id INTEGER,
                                           away_team_id INTEGER,
                                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Add dynamic stats columns here
                                           PRIMARY KEY (match_id, period)
);

CREATE TABLE IF NOT EXISTS match_defending (
                                               match_id INTEGER NOT NULL,
                                               period VARCHAR(20) NOT NULL,
                                               home_team_id INTEGER,
                                               away_team_id INTEGER,
                                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Add dynamic stats columns here
                                               PRIMARY KEY (match_id, period)
);

CREATE TABLE IF NOT EXISTS match_goalkeeping (
                                                 match_id INTEGER NOT NULL,
                                                 period VARCHAR(20) NOT NULL,
                                                 home_team_id INTEGER,
                                                 away_team_id INTEGER,
                                                 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Add dynamic stats columns here
                                                 PRIMARY KEY (match_id, period)
);